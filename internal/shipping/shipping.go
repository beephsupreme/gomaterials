package shipping

import "github.com/beephsupreme/gomaterials/internal/config"

var app *config.AppConfig

func LoadConfig(a *config.AppConfig) {
	app = a
}
