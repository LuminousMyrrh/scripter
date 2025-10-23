package mainconfig

import (
	"encoding/json"
	"log"
	"os"
	"scripter/internal/utils"
)

type MainConfig struct {
	Templates []string
}

func NewMainConfig() *MainConfig {
	return &MainConfig{}
}

func (mainCfg *MainConfig) readMainConfig(configFile string, xdgConfigDir string) error {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(configData, &mainCfg); err != nil {
		return err
	}

	templateDir := xdgConfigDir + "/scripter" + "/templates/"

	for i, temp := range mainCfg.Templates {
		mainCfg.Templates[i] = templateDir + temp
	}

	return nil
}

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

