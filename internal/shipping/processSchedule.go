package shipping

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/pkg/utility"
)

// MakeTable converts []string into [][]string
func MakeTable(S []string) ([][]string, int, strings.Builder) {
	var data [][]string
	var row []string
	var td string
	var firstLine, firstDate, numDates int
	var header strings.Builder
	fmt.Println("Processing...")
	// Find start of regular data
	for firstLine, td = range S {
		if td == config.FIRSTLINE_TEXT {
			break
		}
	}
	// Backup until you have the start of shipping dates
	for firstDate = firstLine; firstDate >= 0; firstDate-- {
		if S[firstDate] == "" {
			firstDate++
			break
		}
	}
	// set number of ship dates
	numDates = firstLine - firstDate
	// Copy shipping dates to correct location
	// Also copy to global var 'header' [bad but convienient]
	for i := 0; i < numDates; i++ {
		S[firstLine+numDates+i] = S[firstLine-numDates+i]
		_, _ = fmt.Fprintf(&header, ",%s", S[firstLine-numDates+i])
	}
	// Remove uneeded first lines
	S = S[firstLine:]
	// Convert shipping from data []string to table [][]string
	for i := 0; i < len(S)-(numDates+config.WIDTH); i += numDates + config.WIDTH {
		for j := 0; j < numDates+config.WIDTH; j++ {
			// Ignore 2nd through 5th columns
			if j == 1 || j == 2 || j == 3 || j == 4 {
				continue
			} else {
				td = S[j+i]
				if strings.Contains(td, "@") {
					td = "0"
				}
				row = append(row, td)
			}
		}
		data = append(data, row)
		row = []string{}
	}
	return data, numDates, header
}

// MakeMap takes the [][]string from ScheduleToTable and returns it as a map
func MakeMap(T [][]string) map[string][]float64 {
	tWidth := len(T[0][0:])
	m := make(map[string][]float64)
	for _, t := range T[1:] {
		if v, ok := m[t[0]]; ok {
			for j := 1; j < tWidth; j++ {
				f, err := strconv.ParseFloat(t[j], config.BITS)
				utility.CheckError("shipping.MakeMap.ParseFloat(if):", err)
				v[j-1] += f
			}
			m[t[0]] = v
		} else {
			d := make([]float64, tWidth-1)
			for j := 1; j < tWidth; j++ {
				f, err := strconv.ParseFloat(t[j], config.BITS)
				utility.CheckError("shipping.MakeMap.ParseFloat(else):", err)
				d[j-1] = f
			}
			m[t[0]] = d
		}
	}
	return m
}

// Validate checks for invalid part numbers corrects them
func Validate(s [][]string) [][]string {
	fmt.Println("Validating...")
	return s
}

// Translate converts Toki units into TLI units
func Translate(s [][]string) [][]string {
	fmt.Println("Translating...")
	return s
}
