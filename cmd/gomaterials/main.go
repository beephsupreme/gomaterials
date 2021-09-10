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

func main() {
	start := time.Now()
	fmt.Println("Materials 12.0, (c) 2021 Michael N. Rowsey")
	var header, dates, out strings.Builder
	var count int
	out.Grow(128)
	_, _ = fmt.Fprintf(&header, "%s", config.HEADER)
	data := inventory.LoadData(utility.LoadFile(config.DATA))
	backlog := sales.LoadData(utility.LoadFile(config.BACKLOG))
	hfr := sales.LoadData(utility.LoadFile(config.HFR))
	scheduleList := shipping.GetSchedule(config.URL)

	fmt.Println("Processing shipping schedule...")
	scheduleTable, count, dates := shipping.ScheduleToTable(scheduleList)
	_, _ = fmt.Fprintf(&header, "%s\n", dates.String())

	fmt.Println("Validating shipping schedule...")
	scheduleTable = shipping.ValidateSchedule(scheduleTable)
	scheduleTable = shipping.ConvertUnits(scheduleTable)
	schedule := shipping.ScheduleToMap(scheduleTable)

	fmt.Println("Generating report...")
	report.CreateReport(data, backlog, hfr, schedule, count, &out, &header)

	et := time.Since(start)
	fmt.Printf("Task complete! (%.3g seconds)\n", et.Seconds())
}
