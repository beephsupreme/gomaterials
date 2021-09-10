package shipping

import (
	"fmt"

	"github.com/beephsupreme/gomaterials/pkg/utility"
	"github.com/gocolly/colly"
)

// Retrieve downloads an html table stores the 'td's as an []string
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
	utility.CheckError("[shipping.Retrieve.colly.Visit()] ", err)
	return td
}
