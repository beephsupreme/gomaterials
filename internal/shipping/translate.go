package shipping

import (
	"fmt"
	"strconv"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/beephsupreme/gomaterials/internal/config"
)

// Translate converts Toki units into TLI units
func Translate(S, T [][]string, width int) [][]string {
	fmt.Println("Translating...")
	m := MakeTranslationMap(T)

	for _, row := range S {
		// Look for a conversion factor
		if factor, ok := m[row[config.PN]]; !ok {
			continue
		} else {
			// Convert values if conversion factor found for this row
			for i := 1; i < width; i++ {
				f, err := strconv.ParseFloat(row[i], 64)
				helpers.CheckError("[shipping.Translation.ParseFloat(factor)] ", err)
				row[i] = fmt.Sprintf("%f", factor*f)
			}
		}
	}

	return S
}

func MakeTranslationMap(S [][]string) map[string]float64 {
	m := make(map[string]float64)

	for i := 1; i < len(S); i++ {
		pn := S[i][config.PN]
		factor, err := strconv.ParseFloat(S[i][config.FACTOR], config.BITS)
		helpers.CheckError("[shipping.MakeTranslationMap.ParseFloat(factor)] ", err)
		m[pn] = factor
	}
	return m
}
