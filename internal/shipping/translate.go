package shipping

import (
	"fmt"
	"strconv"

	"github.com/beephsupreme/gomaterials/helpers"
)

// Translate converts Toki units into TLI units
func Translate(S, T [][]string) [][]string {
	fmt.Println("Translating...")
	m := MakeTranslationMap(T)

	for _, row := range S {
		// Look for a conversion factor
		if factor, ok := m[row[app.PN]]; !ok {
			continue
		} else {
			// Convert values if conversion factor found for this row
			for i := 1; i < app.NumDates; i++ {
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
		pn := S[i][app.PN]
		factor, err := strconv.ParseFloat(S[i][app.Factor], app.Bits)
		helpers.CheckError("[shipping.MakeTranslationMap.ParseFloat(factor)] ", err)
		m[pn] = factor
	}
	return m
}
