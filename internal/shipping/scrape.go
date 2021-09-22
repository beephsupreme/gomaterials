package shipping

import (
	"fmt"

	"github.com/beephsupreme/gomaterials/helpers"
	"github.com/gocolly/colly"
)

// Retrieve scrapes the Toki Corp. shipping schedule.
// It grabs all 'td' elements an stores themas a []string
func Retrieve(urlToTable string) []string {
	var td []string
	fmt.Println("Retrieving schedule...")
	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("[shipping.Retrieve.colly()] ", err.Error())
	})
	c.OnHTML("td", func(e *colly.HTMLElement) {
		td = append(td, e.Text)
	})
	err := c.Visit(urlToTable)
	helpers.CheckError("[shipping.Retrieve.colly.Visit()] ", err)
	return td
}
