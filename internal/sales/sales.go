package sales

import (
	"fmt"
	"strconv"

	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/pkg/utility"
)

var app *config.AppConfig

// LoadData takes [][]string and returns a map[string]float64
func LoadData(S [][]string) map[string]float64 {
	m := make(map[string]float64)

	fmt.Println("Processing...")
	for i := 1; i < len(S); i++ {
		pn := S[i][config.PN]
		qty, err := strconv.ParseFloat(S[i][config.QTY], config.BITS)
		utility.CheckError("[sales.LoadData.ParseFloat(qty)] ", err)
		uom, err := strconv.ParseFloat(S[i][config.UOM], config.BITS)
		utility.CheckError("[sales.LoadData.ParseFloat(uom)] ", err)
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
