package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type Script struct {
	Name string
	Template string
	Ask struct {
		PName bool `json:"name"`
		PPackages bool `json:"packages"`
	}
}

type Config struct {
	Predef bool
	Scripts []Script
}

type MainConfig struct {
	Templates []string
}

var mainCfg *MainConfig

func IsDirExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}

func initConfigPath(xdgConfigDir string) {
	if len(xdgConfigDir) == 0 {
		xdgConfigDir = "~/.config"
	}

	configPath := xdgConfigDir + "/scripter/"
	os.Mkdir(configPath, 0755)
	os.Mkdir(configPath + "templates", 0755)
	os.Create(configPath + "config.json")
}

func ReadMainConfig(configFile string, xdgConfigDir string) error {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(configData, &mainCfg); err != nil {
		return err
	}

	templateDir := xdgConfigDir + "/scripter" + "/templates/"

	for i, temp := range mainCfg.Templates {
		mainCfg.Templates[i] = templateDir + temp
	}

	return nil
}

type SearchResponse struct {
	Items []struct {
		FullName string `json:"full_name"`
		Description string `json:"description"`
		HTMLURL string `json:"html_url"`
	} `json:"items"`
}

func checkPackage(packageName string) (string, error) {
	if strings.Contains(packageName, "github") {
		return packageName, nil
	}
	// very bad!!!
	if packageName == "gorm" {
		return "gorm.io/" + packageName, nil
	}

	baseUrl := "https://api.github.com/search/repositories"
	params := url.Values{}
	params.Add("q", fmt.Sprintf("%s in:name language:go", packageName))
	params.Add("sort", "stars")
	params.Add("order", "desc")
	params.Add("per_page", "2")

	reqUrl := baseUrl + "?" + params.Encode()

	resp, err := http.Get(reqUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	var res SearchResponse
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if len(res.Items) == 0 {
		return "", err
	}

	return "github.com/" + res.Items[0].FullName, nil
}

func executeSrcipt(script Script, destination string) {
	var name string
	var packages []string

	if script.Ask.PName {
		for {
			fmt.Print("Enter project name: ")
			fmt.Scan(&name)
			if e, err := IsDirExist(name); err != nil{
				log.Fatal("Failed to check dir: ", err)
				break
			} else if e {
				fmt.Printf("Directory %s already exist\n", name)
				continue
			}
			break
		}
	}
	namePath := destination + "/" + name

	if script.Ask.PPackages {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(
			"Enter packages names you want to preinstall (enter when done):")
		if scanner.Scan() {
			packs := scanner.Text()
			packages = strings.Fields(packs)
		}
	}

	os.Mkdir(name, 0755)
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = namePath
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to init go project: ", err)
		return
	}

	for _, pack := range packages {
		fullName, err := checkPackage(pack)
		if err != nil {
			fmt.Printf("Failed to check package: %v\n", err)
			fmt.Print("Want to resume downloading other packages, or retry? [1/2]: ")
			var ans int
			fmt.Scanln(&ans)
			if ans == 2 {
				break
			}
		}

		fmt.Println("Installing: ", fullName)
		packCmd := exec.Command("go", "get", "-u", fullName)
		packCmd.Dir = name
		output, err := packCmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(string(output))
	}

	templatePath := ""
	for i, templ := range mainCfg.Templates {
		parts := strings.Split(templ, "/")
		if parts[len(parts) -1 ] == script.Template {
			templatePath = mainCfg.Templates[i]
		}
	}
	if len(templatePath) == 0 {
		fmt.Printf("Template %s does not exist\n", script.Template)
		return
	}

	err = copyTemplate(templatePath, namePath)
	if err != nil {
		fmt.Println("Failed to copy template: ", err)
		return 
	}
	fmt.Println("Done!")
}

func copyFile(src string, dest string) error {
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcFileStat.Mode().IsRegular() {
		return fmt.Errorf("File %s is not a regular file.", srcFileStat.Name())
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return err 
	}
	if nBytes == 0 {
		return fmt.Errorf("No bytes were copied.")
	}

	return nil
}

func copyFiles(dirPath string, destPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	dirFiles, err := dir.ReadDir(0)
	if err != nil {
		return err
	}
	for _, entity := range dirFiles {
		entityPath := dirPath + "/" + entity.Name()
		fmt.Println(entityPath)
		fmt.Println(destPath)
		if entity.Type().IsDir() {
			copyFiles(entityPath, destPath)
		} else {
			err := copyFile(entityPath, destPath + "/" + entity.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyTemplate(template string, dest string) error {
	err := copyFiles(template, dest)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	xdgConfigDir := os.Getenv("XDG_CONFIG_HOME")
	if len(xdgConfigDir) == 0 {
		xdgConfigDir = os.Getenv("HOME") + "/.config"
	}
	if exist, err := IsDirExist(xdgConfigDir + "/scripter"); err != nil {
		log.Fatal("Failed to check main config dir: ", err)
		return
	} else if !exist {
		log.Println("Initializing main config dir")
		initConfigPath(xdgConfigDir)
	}
	err := ReadMainConfig(xdgConfigDir + "/scripter/config.json", xdgConfigDir)
	if err != nil {
		fmt.Println("Failed to read main config: ", err)
		return
	}

	if len(os.Args) < 2 {
		log.Fatal("usage: scripter run <your command>")
		return
	}

	args := os.Args[1:]

	configData, err := os.ReadFile("scripts.json")
	if err != nil {
		fmt.Println("Failed to read config file: ", err)
		return
	}

	var config Config
	if err = json.Unmarshal(configData, &config); err != nil {
		fmt.Println("Failed to unmarshal config: ", err)
		return
	}

	if args[0] == "run" {
		cmdName := args[1]
		for _, script := range config.Scripts {
			if cmdName == script.Name {
				if len(args) == 3 {
					executeSrcipt(script, args[2])
				} else {
					executeSrcipt(script, ".")
				}
			}
		}
		
	}
}
