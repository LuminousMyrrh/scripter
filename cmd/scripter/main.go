package main

import (
	"fmt"
	"os"
	"scripter/internal/commands"
	"scripter/internal/config"
	"scripter/internal/mainconfig"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: scripter run <your command>")
		os.Exit(1)
	}

	args := os.Args[1:]

	configData, err := os.ReadFile("scripts.json")
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		os.Exit(1)
	}

	mainConfig := mainconfig.NewMainConfig()
	if err := mainConfig.CheckMainConfig(); err != nil {
		fmt.Println("Failed to check main config:", err)
		os.Exit(1)
	}

	localConfig, err := config.NewLocalConfig(configData)
	if err != nil {
		fmt.Println("Failed to read local config:", err)
		os.Exit(1)
	}

	switch args[0] {
	case "run":
		if err := commands.CommandRun(args, localConfig, mainConfig); err != nil {
			fmt.Println("Failed to run 'run':", err)
			os.Exit(1)
		}
	case "make":
		if err := commands.CommandMake(args, mainConfig); err != nil {
			fmt.Println("Failed to run 'make':", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		fmt.Println("usage: scripter run <your command>")
		os.Exit(1)
	}
}
