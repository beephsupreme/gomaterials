package shipping

import (
	"log"
	"strings"

	"github.com/beephsupreme/gomaterials/helpers"
)

// Validate checks for invalid part numbers corrects them
func Validate(S [][]string) [][]string {

	V := helpers.LoadFile(app.Validate)

	m, s := MakeValidationStructures(V)
	for _, pn := range S[1:] {
		found := false
		for i := 1; i < len(s); i++ {
			if pn[app.PN] == s[i] {
				found = true
				break
			}
		}
		if strings.Contains(pn[app.PN], "RV-SEALANT") {
			found = true
		}
		if !found {
			if _, ok := m[pn[app.PN]]; !ok {
				log.Fatal("[shipping.Validate() item not found]: ", pn[app.PN])
			} else {
				pn[app.PN] = m[pn[app.PN]]
			}
		}
	}
	return S
}

// MakeValidate prepares structures to validate and update part numbers
func MakeValidationStructures(S [][]string) (map[string]string, []string) {
	m := make(map[string]string)
	s := make([]string, len(S)-1)
	for i := 1; i < len(S); i++ {
		toki := S[i][app.TOKI]
		tli := S[i][app.TLI]
		s[i-1] = tli
		m[toki] = tli
	}
	return m, s
}
