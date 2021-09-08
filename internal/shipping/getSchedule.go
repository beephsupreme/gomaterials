package shipping

import (
	"fmt"
	"github.com/beephsupreme/gomaterials/pkg/utility"
	"github.com/gocolly/colly"
)

// GetSchedule downloads a table from the internet and stores is as an array of all cells
func GetSchedule(urlToTable string) []string {
	var td []string
	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("GetSchedule():", err.Error())
	})
	c.OnHTML("td", func(e *colly.HTMLElement) {
		td = append(td, e.Text)
	})
	err := c.Visit(urlToTable)
	utility.CheckError("colly.Visit():", err)

	return td
}
