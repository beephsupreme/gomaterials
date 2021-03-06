package sales

import (
	"strconv"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/beephsupreme/gomaterials/internal/config"
)

var app *config.AppConfig

// LoadData takes [][]string and returns a map[string]float64
func LoadData(filename string) map[string]float64 {

	S := helpers.LoadFile(filename)

	m := make(map[string]float64)

	for i := 1; i < len(S); i++ {
		pn := S[i][app.PN]
		qty, err := strconv.ParseFloat(S[i][app.QTY], app.Bits)
		helpers.CheckError("[sales.LoadData.ParseFloat(qty)] ", err)
		uom, err := strconv.ParseFloat(S[i][app.UOM], app.Bits)
		helpers.CheckError("[sales.LoadData.ParseFloat(uom)] ", err)
		if val, ok := m[pn]; !ok {
			m[pn] = qty * uom
		} else {
			m[pn] = val + qty*uom
		}
	}
	return m
}

func LoadConfig(a *config.AppConfig) {
	app = a
}
