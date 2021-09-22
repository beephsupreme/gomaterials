package inventory

import (
	"strconv"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/internal/models"
)

var app *config.AppConfig

// LoadData reads file from AV export and loads each line
// into an array of structures
func LoadData() []models.Data {

	var D [][]string
	var d models.Data
	var data []models.Data
	var err error

	D = helpers.LoadFile(app.Data)

	for i := 1; i < len(D); i++ {
		d.PartNum = D[i][app.PN]
		d.OnHand, err = strconv.ParseFloat(D[i][app.OH], app.Bits)
		helpers.CheckError("[inventory.LoadData.ParseFloat(OnHand)] ", err)
		d.OnOrder, err = strconv.ParseFloat(D[i][app.OO], app.Bits)
		helpers.CheckError("[inventory.LoadData.ParseFloat(OnOrder)] ", err)
		d.ReOrder, err = strconv.ParseFloat(D[i][app.RO], app.Bits)
		helpers.CheckError("[inventory.LoadData.ParseFloat(ReOrder)] ", err)
		data = append(data, d)
	}
	return data
}

// LoadConfig makes AppConfig available to this package
func LoadConfig(a *config.AppConfig) {
	app = a
}
