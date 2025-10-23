package main

import (
	"fmt"
	"log"
	"os"
	"scripter/internal/config"
	"scripter/internal/mainconfig"
	"scripter/internal/script"
)

func main() {
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

	mainConfig := mainconfig.NewMainConfig()
	err = mainConfig.CheckMainConfig()
	if err != nil {
		fmt.Println("Failed to check main config: ", err)
		return
	}

	config, err := config.NewLocalConfig(configData)
	if err != nil {
		fmt.Println("Failed to read local config")
		return
	}

	if args[0] == "run" {
		if len(args) != 2 {
			fmt.Println("Expected script name after run")
			return
		}

		cmdName := args[1]
		var src *script.Script = nil

		for _, script := range config.Scripts {
			if cmdName == script.Name {
				src = &script
			}
		}

		if src == nil {
			fmt.Printf("Script \"%s\" doesn't exist", cmdName)
			return
		}

		if len(args) == 3 {
			err := src.ExecuteSrcipt(mainConfig, args[2])
			if err != nil {
				fmt.Println("Failed to execute script: ", err)
				return
			}
			fmt.Println("Done")
		} else {
			err := src.ExecuteSrcipt(mainConfig, ".")
			if err != nil {
				fmt.Println("Failed to execute script: ", err)
				return
			}
			fmt.Println("Done")
		}
	} else {
		fmt.Printf("Unknown command: %s\n", args[0])
		return
	}
}
