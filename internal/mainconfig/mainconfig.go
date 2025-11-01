package mainconfig

import (
	"encoding/json"
	"os"
)

type MainConfig struct {
	Templates []string `json:"templates"`
	ConfigPath string `json:"-"`
}

func NewMainConfig() *MainConfig {
	return &MainConfig{}
}

type templatesJSON struct {
	Templates []string `json:"templates"`
}

func (mainCfg *MainConfig) readMainConfig(configFile string, xdgConfigDir string) error {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	var templates templatesJSON

	if err := json.Unmarshal(configData, &templates); err != nil {
		return err
	}

	templateDir := xdgConfigDir + "/scripter" + "/templates/"
	
	//validate templates
	if err := mainCfg.ValidateNewTemplates(templates.Templates, templateDir); err != nil {
		return err
	}

	return nil
}

func (m *MainConfig) AddTemplate(tempName string) {
	m.Templates = append(m.Templates, tempName)
}

func (mc *MainConfig) UpdateConfigFile() error {
	updatedTemps, err := json.MarshalIndent(mc, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(mc.ConfigPath + "config.json", updatedTemps, 0644)
	if err != nil {
		return err
	}

	return nil
}
