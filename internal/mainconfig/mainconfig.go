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
	mainCfg.validateTemplates(templates.Templates, templateDir)

	return nil
}

func (m *MainConfig) AddTemplate(tempName string) {
	m.Templates = append(m.Templates, tempName)
}
