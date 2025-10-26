package script

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"scripter/internal/mainconfig"
	"scripter/internal/utils"
	"strings"
)

type Script struct {
	Name     string
	Template string
	Ask      struct {
		PName     bool `json:"name"`
		PPackages bool `json:"packages"`
	}
}

func (script Script) ExecuteSrcipt(mainCfg *mainconfig.MainConfig, destination string) error {
	name := askName(script.Ask.PName)
	packages := askPackages(script.Ask.PPackages)

	namePath := destination + "/" + name

	os.Mkdir(name, 0755)
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = namePath
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to init go project: ", err)
		return err
	}

	for _, pack := range packages {
		err := getPackage(pack, name)
		if err != nil {
			return err
		}
	}

	templatePath := ""
	for i, templ := range mainCfg.Templates {
		parts := strings.Split(templ, "/")
		if parts[len(parts)-1] == script.Template {
			templatePath = mainCfg.Templates[i]
		}
	}
	if len(templatePath) == 0 {
		return fmt.Errorf("Template %s does not exist\n", script.Template)
	}

	err = utils.CopyTemplate(templatePath, namePath)
	if err != nil {
		return err
	}

	return nil
}

func askName(isAsked bool) string {
	var name string

	if isAsked {
		for {
			fmt.Print("Enter project name: ")
			fmt.Scan(&name)
			if e, err := utils.IsDirExist(name); err != nil {
				log.Fatal("Failed to check dir: ", err)
				break
			} else if e {
				fmt.Printf("Directory %s already exist\n", name)
				continue
			}
			break
		}
	}

	return name
}

func askPackages(isAsked bool) []string {
	var packages []string

	if isAsked {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(
			"Enter packages names you want to preinstall (enter when done):")
		if scanner.Scan() {
			packs := scanner.Text()
			packages = strings.Fields(packs)
		}
	}
	return packages
}

func getPackage(pack string, name string) error {
	fullName, err := CheckPackage(pack)
	if err != nil {
		fmt.Printf("Failed to check package: %v\n", err)
		fmt.Print("Want to resume downloading other packages, or retry? [1/2]: ")
		var ans int
		fmt.Scanln(&ans)
		if ans == 2 {
			return errors.New("Break")
		} 
	}

	fmt.Println("Installing: ", fullName)
	packCmd := exec.Command("go", "get", "-u", fullName)
	packCmd.Dir = name
	output, err := packCmd.CombinedOutput()
	if err != nil {
		return err
	}
	
	fmt.Println(string(output))
	return nil
}
