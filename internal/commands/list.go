package commands

import (
	"fmt"
	"scripter/internal/mainconfig"
	"scripter/internal/utils"
)

func CommandList(mainConfig *mainconfig.MainConfig) error {
	err := mainConfig.ValidateExistingTemplates()
	if err != nil {
		return err
	}

	if len(mainConfig.Templates) == 0 {
		fmt.Println("No templates were found")
	} else {
		utils.PerpareTemps(mainConfig.Templates)
		for _, temp := range mainConfig.Templates {
			fmt.Println(temp)
		}
	}

	return nil
}
