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

// CreateReport creates a strings.Builder object which is written to a csv file
func Build(data []models.Data,
	backlog, hfr map[string]float64,
	schedule map[string][]float64,
	count int,
	sb, hdr *strings.Builder) {
	//create materials table using strings.Builder
	fmt.Println("Building report...")
	_, _ = fmt.Fprintf(sb, "%s\n", time.Now().Format("2006-01-02"))
	_, _ = fmt.Fprintf(sb, "%s", hdr.String())
	for _, r := range data {
		pn := r.PartNum
		_, _ = fmt.Fprintf(sb, "%s,%f,%f,%f,%f,%f,%f,%f,%f",
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
			for i := 0; i < count; i++ {
				_, _ = fmt.Fprintf(sb, "%s", ",")
			}
		} else {
			for i := 0; i < count; i++ {
				if v[i] < 1.0 {
					_, _ = fmt.Fprintf(sb, "%s", ",")
				} else {
					_, _ = fmt.Fprintf(sb, ",%f", v[i])
				}
			}
		}
		_, _ = fmt.Fprintf(sb, "%s", "\n")
	}

	f, err := os.Create(config.OUTFILE)
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
