package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/internal/inventory"
	"github.com/beephsupreme/gomaterials/internal/report"
	"github.com/beephsupreme/gomaterials/internal/sales"
	"github.com/beephsupreme/gomaterials/internal/shipping"
	"github.com/beephsupreme/gomaterials/pkg/utility"
)

var app config.AppConfig

func main() {
	start := time.Now()
	fmt.Println("Materials 4.1 \u00A9 2021 Michael N. Rowsey")
	setup()
	run()
	et := time.Since(start)
	fmt.Printf("Task complete! (%.3g seconds)\n", et.Seconds())
	utility.PrintMemUsage()
}

func run() {
	var header, dates, out strings.Builder
	var count int
	out.Grow(128)
	_, _ = fmt.Fprintf(&header, "%s", app.Header)
	data := inventory.LoadData(utility.LoadFile(app.DataPath + app.Data))
	backlog := sales.LoadData(utility.LoadFile(app.DataPath + app.Backlog))
	hfr := sales.LoadData(utility.LoadFile(app.DataPath + app.Hfr))
	validate := utility.LoadFile(app.DataPath + app.Validate)
	translate := utility.LoadFile(app.DataPath + app.Translate)
	scheduleTable, count, dates := shipping.MakeTable(shipping.Retrieve(app.ScheduleURL))
	_, _ = fmt.Fprintf(&header, "%s\n", dates.String())
	scheduleTable = shipping.Validate(scheduleTable, validate)
	scheduleTable = shipping.Translate(scheduleTable, translate, count)
	schedule := shipping.MakeMap(scheduleTable)
	report.Build(data, backlog, hfr, schedule, count, &out, &header)
}

func setup() {
	config.LoadConfig(&app)
	inventory.LoadConfig(&app)
	sales.LoadConfig(&app)
	shipping.LoadConfig(&app)
}
