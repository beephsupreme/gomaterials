package report

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/internal/models"
)

var app *config.AppConfig

// CreateReport creates a strings.Builder object which is written to a csv file
func Build(data []models.Data, backlog, hfr map[string]float64, schedule map[string][]float64) {

	var sb strings.Builder
	sb.Grow(128)

	fmt.Println("Building report...")

	_, _ = fmt.Fprintf(&sb, "%s\n", time.Now().Format("2006-01-02"))
	_, _ = fmt.Fprintf(&sb, "%s\n", app.MainHeader.String())

	for _, r := range data {
		pn := r.PartNum
		_, _ = fmt.Fprintf(&sb, "%s,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f",
			pn,
			r.OnHand,
			backlog[pn],
			backlog[pn]-hfr[pn],
			hfr[pn],
			r.OnOrder,
			r.OnHand+r.OnOrder-backlog[pn],
			r.OnHand+r.OnOrder-backlog[pn]+hfr[pn],
			r.ReOrder)

		if v, ok := schedule[pn]; !ok {
			for i := 0; i < app.NumDates; i++ {
				_, _ = fmt.Fprintf(&sb, "%s", ",")
			}
		} else {
			for i := 0; i < app.NumDates; i++ {
				if v[i] < 1.0 {
					_, _ = fmt.Fprintf(&sb, "%s", ",")
				} else {
					_, _ = fmt.Fprintf(&sb, ",%.0f", v[i])
				}
			}
		}
		_, _ = fmt.Fprintf(&sb, "%s", "\n")
	}

	f, err := os.Create(app.DataPath + app.Outfile)
	helpers.CheckError("[report.Build.Open()] ", err)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			helpers.CheckError("[report.Build.Close()] ", err)
		}
	}(f)
	b, err := f.WriteString(sb.String())
	helpers.CheckError("[report.Build.WriteString(sb)] ", err)
	fmt.Println("Wrote", b, "bytes to disk.")
}

func LoadConfig(a *config.AppConfig) {
	app = a
}
