package script

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"scripter/internal/mainconfig"
	"scripter/internal/utils"
	"strings"
)

func (script Script) ExecuteSrcipt(mainCfg *mainconfig.MainConfig, destination string) error {
	name := askName(script.Ask.PName)
	packages := askPackages(script.Ask.PPackages)
	packages = append(packages, script.InstallPackages...)

	namePath := destination + "/" + name

	os.Mkdir(namePath, 0755)
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = namePath
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to init go project: ", err)
		return err
	}

    re := regexp.MustCompile(`module declares its path as:\s*([\w./-]+)`)
	for _, pack := range packages {
		err := getPackage(pack, name, re)
		if err != nil {
			fmt.Printf("Failed to get package: %v\n", err)
			fmt.Print("Want to resume downloading other packages, or retry? [1/2]: ")
			var ans int
			fmt.Scanln(&ans)
			if ans == 2 {
				return err
			} else {
				continue
			}
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

func getPackage(packName, dirName string, re *regexp.Regexp) error {
	fullName, err := CheckPackage(packName)
	if err != nil {
		return err
	}

	fmt.Println("Installing: ", fullName)
	err = installPackage(fullName, dirName, re)
	return err
}

func installPackage(packName string, dirName string, re *regexp.Regexp) error {
	result, err := commandInstallPackage(packName, dirName)
	if err != nil {
		newPackName := GetPackageNameFromErrOutput(string(result), re)
		if newPackName == "" {
			fmt.Println("go get failed:", err)
			return err
		} else {
			res, err := commandInstallPackage(newPackName, dirName)
			fmt.Println(res)
			return err
		}
	}

	return nil
}

func commandInstallPackage(name, dirName string) (string, error) {
	packCmd := exec.Command("go", "get", "-u", name)
	packCmd.Dir = dirName
	output, err := packCmd.CombinedOutput()
	return string(output), err
}

