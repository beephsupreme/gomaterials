package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/beephsupreme/gomaterials/internal/config"
	"github.com/beephsupreme/gomaterials/internal/inventory"
	"github.com/beephsupreme/gomaterials/internal/report"
	"github.com/beephsupreme/gomaterials/internal/sales"
	"github.com/beephsupreme/gomaterials/internal/shipping"
)

var app config.AppConfig

func main() {
	start := time.Now()
	fmt.Println("Materials 4.2 \u00A9 2021 Michael N. Rowsey")
	setup()
	run()
	et := time.Since(start)
	fmt.Printf("Task complete! (%.3g seconds)\n", et.Seconds())
	helpers.PrintMemUsage()
}

func run() {
	var header, dates, out strings.Builder
	var count int
	out.Grow(128)
	_, _ = fmt.Fprintf(&header, "%s", app.Header)
	data := inventory.LoadData(helpers.LoadFile(app.Data))
	backlog := sales.LoadData(helpers.LoadFile(app.Backlog))
	hfr := sales.LoadData(helpers.LoadFile(app.Hfr))
	validate := helpers.LoadFile(app.Validate)
	translate := helpers.LoadFile(app.Translate)
	scheduleTable, count, dates := shipping.MakeTable(shipping.Retrieve(app.ScheduleURL))
	_, _ = fmt.Fprintf(&header, "%s\n", dates.String())
	scheduleTable = shipping.Validate(scheduleTable, validate)
	scheduleTable = shipping.Translate(scheduleTable, translate, count)
	schedule := shipping.MakeMap(scheduleTable)
	report.Build(data, backlog, hfr, schedule, count, &out, &header)
}

func setup() {
	config.LoadConfig(&app)
	helpers.LoadConfig(&app)
	inventory.LoadConfig(&app)
	sales.LoadConfig(&app)
	shipping.LoadConfig(&app)
}
