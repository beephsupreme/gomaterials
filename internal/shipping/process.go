package shipping

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beephsupreme/gomaterials/helpers"
)

// MakeTable converts []string into [][]string
func MakeTable() [][]string {
	var data [][]string
	var row []string
	var td string
	var firstLine, firstDate int

	S := Retrieve()

	// Find start of regular data
	for firstLine, td = range S {
		if td == app.FirstLineText {
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
	app.NumDates = firstLine - firstDate
	// Copy shipping dates to correct location
	for i := 0; i < app.NumDates; i++ {
		S[firstLine+app.ScheduleWidth+i] = S[firstLine-app.NumDates+i]
		_, _ = fmt.Fprintf(&app.MainHeader, ",%s", S[firstLine-app.NumDates+i])
	}

	// Remove uneeded first lines
	S = S[firstLine:]
	// Convert shipping from data []string to table [][]string
	for i := 0; i < len(S)-(app.NumDates+app.ScheduleWidth); i += app.NumDates + app.ScheduleWidth {
		for j := 0; j < app.NumDates+app.ScheduleWidth; j++ {
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
	data = Validate(data)
	data = Translate(data)
	return data
}

// MakeMap takes the [][]string from ScheduleToTable and returns it as a map
func LoadData() map[string][]float64 {

	T := MakeTable()
	tWidth := len(T[0][0:])
	m := make(map[string][]float64)
	for _, t := range T[1:] {
		if v, ok := m[t[0]]; ok {
			for j := 1; j < tWidth; j++ {
				f, err := strconv.ParseFloat(t[j], app.Bits)
				helpers.CheckError("[shipping.MakeMap.ParseFloat(if)] ", err)
				v[j-1] += f
			}
			m[t[0]] = v
		} else {
			d := make([]float64, tWidth-1)
			for j := 1; j < tWidth; j++ {
				f, err := strconv.ParseFloat(t[j], app.Bits)
				helpers.CheckError("[shipping.MakeMap.ParseFloat(else)] ", err)
				d[j-1] = f
			}
			m[t[0]] = d
		}
	}
	return m
}
