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
	run()
	et := time.Since(start)
	fmt.Printf("Total runtime: %.3g seconds\n", et.Seconds())
	helpers.PrintMemUsage()
}

// run executes each step of building the materials report
func run() {
	fmt.Println(app.AppVersion)
	data := inventory.LoadData()
	backlog := sales.LoadData(app.Backlog)
	hfr := sales.LoadData(app.Hfr)
	schedule := shipping.LoadData()
	report.Build(data, backlog, hfr, schedule)
}

// setup loads config.AppConfig and passes a reference to each package
// to facilitate data sharing throughout the app
func setup() {
	config.LoadConfig(&app)
	helpers.LoadConfig(&app)
	inventory.LoadConfig(&app)
	sales.LoadConfig(&app)
	shipping.LoadConfig(&app)
	report.LoadConfig(&app)
}
