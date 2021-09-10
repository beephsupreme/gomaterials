package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/internal/report"
	"github.com/beephsupreme/gomaterials/internal/shipping"
	"github.com/beephsupreme/gomaterials/pkg/utility"
)

type Data struct {
	PartNum string
	OnHand  float64
	OnOrder float64
	ReOrder float64
}

func main() {
	start := time.Now()
	fmt.Println("Initializing...")
	var header, dates, out strings.Builder
	var count int
	out.Grow(128)
	_, _ = fmt.Fprintf(&header, "%s", config.HEADER)

	fmt.Println("Loading files...")
	dataTable := utility.LoadFile(config.DATA)
	backlogTable := utility.LoadFile(config.BACKLOG)
	hfrTable := utility.LoadFile(config.HFR)

	fmt.Println("Building tables...")
	data := LoadInventoryData(dataTable)
	backlog := LoadSalesData(backlogTable)
	hfr := LoadSalesData(hfrTable)

	fmt.Println("Retrieving shipping schedule...")
	scheduleList := shipping.GetSchedule(config.URL)

	fmt.Println("Processing shipping schedule...")
	scheduleTable, count, dates := shipping.ScheduleToTable(scheduleList)
	_, _ = fmt.Fprintf(&header, "%s\n", dates.String())

	fmt.Println("Validating shipping schedule...")
	scheduleTable = shipping.ValidateSchedule(scheduleTable)
	scheduleTable = shipping.ConvertUnits(scheduleTable)

	fmt.Println("Building table...")
	schedule := shipping.ScheduleToMap(scheduleTable)

	fmt.Println("Generating report...")
	report.CreateReport(data, backlog, hfr, schedule, count, &out, &header)

	et := time.Since(start)
	fmt.Printf("Task complete! (%.3g seconds)\n", et.Seconds())
}

// LoadInventoryData takes [][]string and returns it as an array of structs
func LoadInventoryData(D [][]string) []Data {
	var d Data
	var data []Data
	var err error
	for i := 1; i < len(D); i++ {
		d.PartNum = D[i][config.PN]
		d.OnHand, err = strconv.ParseFloat(D[i][config.OH], config.BITS)
		utility.CheckError("parse OnHand:", err)
		d.OnOrder, err = strconv.ParseFloat(D[i][config.OO], config.BITS)
		utility.CheckError("parse OnOrder:", err)
		d.ReOrder, err = strconv.ParseFloat(D[i][config.RO], config.BITS)
		utility.CheckError("parse ReOrder:", err)
		data = append(data, d)
	}
	return data
}

// LoadSalesData takes [][]string and returns it as a map
func LoadSalesData(S [][]string) map[string]float64 {
	m := make(map[string]float64)
	for i := 1; i < len(S); i++ {
		pn := S[i][config.PN]
		qty, err := strconv.ParseFloat(S[i][config.QTY], config.BITS)
		utility.CheckError("parse OnHand:", err)
		uom, err := strconv.ParseFloat(S[i][config.UOM], config.BITS)
		utility.CheckError("parse OnOrder:", err)

		if val, ok := m[pn]; !ok {
			m[pn] = qty * uom
		} else {
			m[pn] = val + qty*uom
		}
	}
	return m
}
