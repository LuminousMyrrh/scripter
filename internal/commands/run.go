package commands

import (
	"errors"
	"fmt"
	"scripter/internal/config"
	"scripter/internal/mainconfig"
	"scripter/internal/script"
)

func CommandRun(args []string, config *config.Config,
	mainConfig *mainconfig.MainConfig) error {
	if len(args) != 2 {
		return errors.New("Expected script name after run")
	}

	cmdName := args[1]
	var src *script.Script = nil

	for _, script := range config.Scripts {
		if cmdName == script.Name {
			src = &script
		}
	}

	if src == nil {
		return fmt.Errorf("Script \"%s\" doesn't exist", cmdName)
	}

	secArg := "."
	if len(args) == 3 {
		secArg = args[2]
	}

	err := src.ExecuteSrcipt(mainConfig, secArg)
	if err != nil {
		return err
	}

	return nil
}


