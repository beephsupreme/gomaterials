package shipping

import (
	"fmt"
	"time"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/gocolly/colly"
)

// Retrieve scrapes the Toki Corp. shipping schedule.
// It grabs all 'td' elements an stores themas a []string
func Retrieve() []string {
	var td []string
	tb := time.Now()
	fmt.Println("Retrieving", app.ScheduleURL)
	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("[shipping.Retrieve.colly()] ", err.Error())
	})
	c.OnHTML("td", func(e *colly.HTMLElement) {
		td = append(td, e.Text)
	})
	err := c.Visit(app.ScheduleURL)
	helpers.CheckError("[shipping.Retrieve.colly.Visit()] ", err)
	te := time.Since(tb)
	fmt.Printf("Time to retrieve: %.3g seconds\n", te.Seconds())
	return td
}
