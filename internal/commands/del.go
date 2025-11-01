package commands

import (
	"errors"
	"fmt"
	"os"
	"scripter/internal/mainconfig"
	"scripter/internal/utils"
	"strings"
)

func CommandDel(args []string, mainConfig *mainconfig.MainConfig) error {
	if len(args) < 2 {
		return errors.New("Expected template name")
	}
	isDeleted := false

	tempName := args[1]
	for i, temp := range mainConfig.Templates {
		parts := strings.Split(temp, "/")
		name := parts[len(parts)-1]
		if name == tempName {
			err := os.RemoveAll(mainConfig.Templates[i])
			if err != nil {
				return err
			}
			isDeleted = true
			break
		}
	}

	if !isDeleted {
		return fmt.Errorf("Template %s not deleted", tempName)
	}

	err := mainConfig.ValidateExistingTemplates()
	if err != nil {
		return err 
	}

	utils.PerpareTemps(mainConfig.Templates)
	
	err = mainConfig.UpdateConfigFile()
	if err != nil {
		return err
	}

	fmt.Printf("Template '%s' successfully deleted\n", tempName)
	return nil
}
