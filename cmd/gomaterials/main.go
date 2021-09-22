package main

import (
	"fmt"
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
	setup()
	fmt.Println(app.AppVersion)
	run()
	et := time.Since(start)
	if app.Debug {
		fmt.Printf("ET: %.3g seconds\n", et.Seconds())
		helpers.PrintMemUsage()
	}
}

func run() {
	data := inventory.LoadData()
	backlog := sales.LoadData(app.Backlog)
	hfr := sales.LoadData(app.Hfr)
	scheduleTable := shipping.MakeTable(shipping.Retrieve(app.ScheduleURL))
	scheduleTable = shipping.Validate(scheduleTable)
	scheduleTable = shipping.Translate(scheduleTable)
	schedule := shipping.MakeMap(scheduleTable)
	report.Build(data, backlog, hfr, schedule)
}

func setup() {
	config.LoadConfig(&app)
	helpers.LoadConfig(&app)
	inventory.LoadConfig(&app)
	sales.LoadConfig(&app)
	shipping.LoadConfig(&app)
	report.LoadConfig(&app)
}
