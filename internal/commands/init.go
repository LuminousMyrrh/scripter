package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"scripter/internal/config"
	"scripter/internal/mainconfig"
)

func CommandInit(mainConfig *mainconfig.MainConfig) error {
	name := "scripts.json"

	if _, err := os.Stat(name); os.IsNotExist(err) {
		defaultConfig := config.NewDefault()

		data, err := json.MarshalIndent(defaultConfig, "", " ")
		if err != nil {
			return nil
		}

		err = os.WriteFile("scripts.json", data, 0644)
		if err != nil {
			return err
		}

		fmt.Println("Done")
	} else {
		fmt.Println("Scripts file already exist!")
	}
	
	return nil
}
