package main

import (
	"fmt"
	"os"
	"scripter/internal/commands"
	"scripter/internal/config"
	"scripter/internal/mainconfig"
)

func printHelp() {
	fmt.Println("Usage: scripter <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  run    - Run a script defined in scripts.json")
	fmt.Println("  make   - Make new template")
	fmt.Println("  init   - Initialize current folder")
	fmt.Println("  del    - Delete template")
	fmt.Println("  list   - List available templates")
	fmt.Println("  -h, --help - Show this help message")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: scripter <cmd>")
		os.Exit(1)
	}

	args := os.Args[1:]

	if args[0] == "-h" || args[0] == "--help" {
		printHelp()
		return
	}

	mainConfig := mainconfig.NewMainConfig()
	if err := mainConfig.CheckMainConfig(); err != nil {
		fmt.Println("Failed to check main config:", err)
		os.Exit(1)
	}

	switch args[0] {
	case "run":
		configData, err := os.ReadFile("scripts.json")
		if err != nil {
			fmt.Println("Failed to read config file:", err)
			os.Exit(1)
		}

		localConfig, err := config.NewLocalConfig(configData)
		if err != nil {
			fmt.Println("Failed to read local config:", err)
			os.Exit(1)
		}
		if err := commands.CommandRun(args, localConfig, mainConfig); err != nil {
			fmt.Println("Failed to run 'run':", err)
			os.Exit(1)
		}
	case "make":
		if err := commands.CommandMake(args, mainConfig); err != nil {
			fmt.Println("Failed to run 'make':", err)
			os.Exit(1)
		}
	case "init":
		if err := commands.CommandInit(mainConfig); err != nil {
			fmt.Println("Failed to run 'init':", err)
			os.Exit(1)
		}
	case "del":
		if err := commands.CommandDel(args, mainConfig); err != nil {
			fmt.Println("Failed to run 'del':", err)
			os.Exit(1)
		}
	case "list":
		if err := commands.CommandList(mainConfig); err != nil {
			fmt.Println("Failed to run 'list':", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		printHelp()
		os.Exit(1)
	}
}
