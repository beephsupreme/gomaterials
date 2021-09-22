package inventory

import (
	"fmt"
	"strconv"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/internal/models"
)

var app *config.AppConfig

// LoadData takes [][]string created from accounting
// system export (data.csv) and loads into a 'Data' struct
func LoadData(D [][]string) []models.Data {
	var d models.Data
	var data []models.Data
	var err error

	fmt.Println("Processing...")
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

func LoadConfig(a *config.AppConfig) {
	app = a
}
