package script

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"scripter/internal/utils"
	"strings"
)

func askName(isAsked bool) string {
	var name string

	if isAsked {
		for {
			fmt.Print("Enter project name: ")
			fmt.Scan(&name)
			if e, err := utils.IsDirExist(name); err != nil {
				log.Fatal("Failed to check dir: ", err)
				break
			} else if e {
				fmt.Printf("Directory %s already exist\n", name)
				continue
			}
			break
		}
	}

	return name
}

func askPackages(isAsked bool) []string {
	var packages []string

	if isAsked {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(
			"Enter packages names you want to preinstall (enter when done):")
		if scanner.Scan() {
			packs := scanner.Text()
			packages = strings.Fields(packs)
		}
	}
	return packages
}

