package commands

import (
	"os"
	"scripter/internal/mainconfig"
)

func CommandInit(mainConfig *mainconfig.MainConfig) error {
	file, err := os.Create("./scripts.json")
	if err != nil {
		return err
	}
	

	
	return nil
}
