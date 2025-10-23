package main

import (
	"fmt"
	"log"
	"os"
	"scripter/internal/commands"
	"scripter/internal/config"
	"scripter/internal/mainconfig"
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

	switch (args[0]) {
	case "run":
		err := commands.CommandRun(args, config, mainConfig)
		if err != nil {
			fmt.Println("Failed to run 'run':", err)
			return
		}
		fmt.Println("Done")
	case "make":
		err := commands.CommandMake(args, mainConfig)
		if err != nil {
			fmt.Println("Failed to run 'make':", err)
			return
		}
		fmt.Println("Done")
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
	}
}
