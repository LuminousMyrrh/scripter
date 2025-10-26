package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"scripter/internal/mainconfig"
	"scripter/internal/utils"
	"strings"
)

func CommandMake(args []string, mainConfig *mainconfig.MainConfig) error {
	if (len(args) > 3) {
		fmt.Print("Scipping args after: ", args[1])
	}

	path := "."
	if (len(args) == 3) {
		path = args[2]
	}

	tempsPath := mainConfig.ConfigPath + "templates/"
	tempName, err := getName(args, tempsPath)
	if err != nil {
		return err
	}
	
	templatePath := tempsPath + tempName
	os.Mkdir(templatePath, 0755)

	err = utils.CopyDir(path, templatePath)
	if err != nil {
		return err
	}

	mainConfig.AddTemplate(tempName)
	perpareTemps(mainConfig)
	
	updatedTemps, err := json.MarshalIndent(mainConfig, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(mainConfig.ConfigPath + "config.json", updatedTemps, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Added new template: ", tempName)

	return nil
}

func getName(args []string, cfgDir string) (string, error) {
    var name string

    if len(args) == 2 {
        exists, err := utils.IsDirExist(cfgDir + args[1])
        if err != nil {
            return "", err
        }
        if exists {
            return "", fmt.Errorf("Template %s already exists", args[1])
        }
        name = args[1]
    } else {
        var err error
        name, err = askTemplateName(cfgDir)
        if err != nil {
            return "", err
        }
    }

    return name, nil
}

func askTemplateName(cfgDir string) (string, error) {
	var name string

	for {
		fmt.Print("Enter template name: ")
		fmt.Scan(&name)
		if e, err := utils.IsDirExist(cfgDir + name); err != nil {
			return "", err
		} else if e {
			fmt.Printf("Template %s already exist\n", name)
			continue
		}
		break
	}

	return name, nil
}

func perpareTemps(mainConfig *mainconfig.MainConfig) {
	for i, temp := range mainConfig.Templates {
		parts := strings.Split(temp, "/")
		mainConfig.Templates[i] = parts[len(parts)-1]
	}
}
