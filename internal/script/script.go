package script

type Script struct {
	Name     string `json:"name"`
	Template string `json:"template"`
	Ask      struct {
		PName     bool `json:"name"`
		PPackages bool `json:"packages"`
	} `json:"ask"`
	InstallPackages []string `json:"installPackages"`
}

