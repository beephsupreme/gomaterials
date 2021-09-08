package main

import (
	"fmt"
	"github.com/beephsupreme/gomaterials/internal/shipping"
	"github.com/beephsupreme/gomaterials/pkg/utility"
	"os"
	"strconv"
	"strings"
	"time"
)

type Data struct {
	PartNum string
	OnHand  float64
	OnOrder float64
	ReOrder float64
}

const (
	PN      = 0
	QTY     = 1
	OH      = 1
	UOM     = 2
	OO      = 2
	RO      = 3
	BITS    = 64
	DATA    = "data.csv"
	BACKLOG = "bl.csv"
	HFR     = "hfr.csv"
	URL     = "https://www.toki.co.jp/purchasing/TLIHTML.files/sheet001.htm"
	OUTFILE = "materials.csv"
	HEADER  = "Part Number,On Hand,Backlog,Released,HFR,On Order,T-Avail,R-Avail,Reorder"
)

var header, dates, out strings.Builder
var count int

func main() {
	start := time.Now()
	fmt.Println("Initializing...")
	out.Grow(128)
	_, _ = fmt.Fprintf(&header, "%s", HEADER)
	fmt.Println("Loading data.csv...")
	dataTable := utility.LoadFile(DATA)
	fmt.Println("Formatting data...")
	data := LoadInventoryData(dataTable)
	fmt.Println("Loading bl.csv...")
	backlogTable := utility.LoadFile(BACKLOG)
	fmt.Println("Formatting bl.csv...")
	backlog := LoadSalesData(backlogTable)
	fmt.Println("Loading hfr.csv...")
	hfrTable := utility.LoadFile(HFR)
	fmt.Println("Formatting hfr.csv...")
	hfr := LoadSalesData(hfrTable)
	fmt.Println("Retrieving shipping...")
	scheduleList := shipping.GetSchedule(URL)
	fmt.Println("Parsing shipping...")
	scheduleTable, count, dates := shipping.ScheduleToTable(scheduleList)
	_, _ = fmt.Fprintf(&header, "%s\n", dates.String())
	fmt.Println("Validating shipping...")
	scheduleTable = shipping.ValidateSchedule(scheduleTable)
	fmt.Println("Converting units...")
	scheduleTable = shipping.ConvertUnits(scheduleTable)
	fmt.Println("Formatting shipping...")
	schedule := shipping.ScheduleToMap(scheduleTable)
	fmt.Println("Generating report...")
	CreateReport(data, backlog, hfr, schedule, count)
	et := time.Since(start)
	fmt.Printf("Task complete! (%.3g seconds)\n", et.Seconds())

	// need to add part number validation on shipping
	// need to add unit conversion to shipping
}

// LoadInventoryData takes [][]string and returns it as an array of structs
func LoadInventoryData(D [][]string) []Data {
	var d Data
	var data []Data
	var err error
	for i := 1; i < len(D); i++ {
		d.PartNum = D[i][PN]
		d.OnHand, err = strconv.ParseFloat(D[i][OH], BITS)
		utility.CheckError("parse OnHand:", err)
		d.OnOrder, err = strconv.ParseFloat(D[i][OO], BITS)
		utility.CheckError("parse OnOrder:", err)
		d.ReOrder, err = strconv.ParseFloat(D[i][RO], BITS)
		utility.CheckError("parse ReOrder:", err)
		data = append(data, d)
	}
	return data
}

// LoadSalesData takes [][]string and returns it as a map
func LoadSalesData(S [][]string) map[string]float64 {
	m := make(map[string]float64)
	for i := 1; i < len(S); i++ {
		pn := S[i][PN]
		qty, err := strconv.ParseFloat(S[i][QTY], BITS)
		utility.CheckError("parse OnHand:", err)
		uom, err := strconv.ParseFloat(S[i][UOM], BITS)
		utility.CheckError("parse OnOrder:", err)

		if val, ok := m[pn]; !ok {
			m[pn] = qty * uom
		} else {
			m[pn] = val + qty*uom
		}
	}
	return m
}

// CreateReport creates a strings.Builder object which is written to a csv file
func CreateReport(data []Data, backlog, hfr map[string]float64, schedule map[string][]float64, count int) {
	// create materials table using strings.Builder
	_, _ = fmt.Fprintf(&out, "%s\n", header.String())
	for _, r := range data {
		pn := r.PartNum
		_, _ = fmt.Fprintf(&out, "%s,%f,%f,%f,%f,%f,%f,%f,%f",
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
				_, _ = fmt.Fprintf(&out, "%s", ",")
			}
		} else {
			for i := 0; i < count; i++ {
				if v[i] < 1.0 {
					_, _ = fmt.Fprintf(&out, "%s", ",")
				} else {
					_, _ = fmt.Fprintf(&out, ",%f", v[i])
				}
			}
		}
		_, _ = fmt.Fprintf(&out, "%s", "\n")
	}

	f, err := os.Create(OUTFILE)
	utility.CheckError("os.Create()", err)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			utility.CheckError("os.Create()", err)
		}
	}(f)
	b, err := f.WriteString(out.String())
	utility.CheckError("WriteString()", err)
	fmt.Println("Wrote", b, "bytes to disk.")
}
