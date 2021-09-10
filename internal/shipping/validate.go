package shipping

import (
	"fmt"
	"log"

	"github.com/beephsupreme/gomaterials/internal/config"
)

// Validate checks for invalid part numbers corrects them
func Validate(S, V [][]string) [][]string {
	fmt.Println("Validating...")
	m, s := MakeValidationStructures(V)
	for _, pn := range S[1:] {
		found := false
		for i := 1; i < len(s); i++ {
			if pn[0] == s[i] {
				found = true
				break
			}
		}
		if !found {
			if _, ok := m[pn[0]]; !ok {
				log.Fatal("[shipping.Validate() item not found]: ", pn[0])
			} else {
				fmt.Println("found conversion")
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
