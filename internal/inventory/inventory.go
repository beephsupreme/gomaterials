package inventory

import (
	"fmt"
	"strconv"

	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/internal/models"
	"github.com/beephsupreme/gomaterials/pkg/utility"
)

// LoadInventoryData takes [][]string and returns an array of structs
func LoadData(D [][]string) []models.Data {
	var d models.Data
	var data []models.Data
	var err error

	fmt.Println("Loading Inventory Data...")
	for i := 1; i < len(D); i++ {
		d.PartNum = D[i][config.PN]
		d.OnHand, err = strconv.ParseFloat(D[i][config.OH], config.BITS)
		utility.CheckError("inventory.LoadData.ParseFloat(OnHand)", err)
		d.OnOrder, err = strconv.ParseFloat(D[i][config.OO], config.BITS)
		utility.CheckError("inventory.LoadData.ParseFloat(OnOrder)", err)
		d.ReOrder, err = strconv.ParseFloat(D[i][config.RO], config.BITS)
		utility.CheckError("inventory.LoadData.ParseFloat(ReOrder)", err)
		data = append(data, d)
	}
	return data
}
