package mainconfig

import (
	"log"
	"os"
	"scripter/internal/utils"
)

func (mainCfg *MainConfig) CheckMainConfig() error {
	xdgConfigDir := os.Getenv("XDG_CONFIG_HOME")
	if len(xdgConfigDir) == 0 {
		xdgConfigDir = os.Getenv("HOME") + "/.config"
	}
	if exist, err := utils.IsDirExist(xdgConfigDir + "/scripter"); err != nil {
		return err
	} else if !exist {
		log.Println("Initializing main config dir")
		InitMainConfigPath(xdgConfigDir)
	}

	mainCfg.ConfigPath = xdgConfigDir + "/scripter/"

	err := mainCfg.readMainConfig(xdgConfigDir + "/scripter/config.json",
		xdgConfigDir)

	return err
}

func InitMainConfigPath(xdgConfigDir string) {
	if len(xdgConfigDir) == 0 {
		xdgConfigDir = "~/.config"
	}

	configPath := xdgConfigDir + "/scripter/"
	os.Mkdir(configPath, 0755)
	os.Mkdir(configPath + "templates", 0755)
	os.Create(configPath + "config.json")
}

func (mc *MainConfig) validateTemplates(templates []string, templateDir string) error {
	for _, temp := range templates {
		fullDirTemp := templateDir + temp
		if exist, err := utils.IsDirExist(fullDirTemp); err != nil {
			return err
		} else if exist {
			mc.Templates = append(mc.Templates, fullDirTemp)
		}
	}

	return nil
}

