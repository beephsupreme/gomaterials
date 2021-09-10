package shipping

import (
	"fmt"
	"log"
	"strings"

	"github.com/beephsupreme/gomaterials/internal/config"
)

// Validate checks for invalid part numbers corrects them
func Validate(S, V [][]string) [][]string {
	fmt.Println("Validating...")
	m, s := MakeValidationStructures(V)
	for _, pn := range S[1:] {
		found := false
		for i := 1; i < len(s); i++ {
			if pn[config.PN] == s[i] {
				found = true
				break
			}
		}
		if strings.Contains(pn[config.PN], "RV-SEALANT") {
			found = true
		}
		if !found {
			if _, ok := m[pn[config.PN]]; !ok {
				log.Fatal("[shipping.Validate() item not found]: ", pn[config.PN])
			} else {
				pn[config.PN] = m[pn[config.PN]]
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
		toki := S[i][config.TOKI]
		tli := S[i][config.TLI]
		s[i-1] = tli
		m[toki] = tli
	}
	return m, s
}
