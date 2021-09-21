package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/beephsupreme/gomaterials/pkg/utility"
)

type AppConfig struct {
	DataPath        string `json:"DataPath"`
	Backlog         string `json:"Backlog"`
	Data            string `json:"Data"`
	Hfr             string `json:"Hfr"`
	Outfile         string `json:"Outfile"`
	Translate       string `json:"Translate"`
	Validate        string `json:"Validate"`
	ScheduleURL     string `json:"ScheduleURL"`
	Header          string `json:"Header"`
	FirstLineText   string `json:"FirstLineText"`
	FirstLineNumber int
	Factor          int `json:"Factor"`
	Bits            int `json:"Bits"`
	ScheduleWidth   int `json:"ScheduleWidth"`
	NumDates        int
	PN              int `json:"PN"`
	OH              int `json:"OH"`
	OO              int `json:"OO"`
	RO              int `json:"RO"`
	QTY             int `json:"QTY"`
	TLI             int `json:"TLI"`
	TOKI            int `json:"TOKI"`
	UOM             int `json:"UOM"`
}

func RunConfig(app *AppConfig) {

	// then config file settings
	//var a AppConfig
	content, err := os.ReadFile("/Users/michael/Dropbox/gomaterials/assets/data/config.json")
	if err != nil {
		utility.CheckError("opening config file", err)
	}

	err = json.Unmarshal(content, app)
	if err != nil {
		utility.CheckError("unmarshalling config file", err)
	}

	fmt.Println("DataPath: ", app.DataPath)
	fmt.Println("Backlog: ", app.Backlog)
	fmt.Println("Data: ", app.Data)
	fmt.Println("Hfr: ", app.Hfr)
	fmt.Println("Outfile: ", app.Outfile)
	fmt.Println("Translate: ", app.Translate)
	fmt.Println("Validate: ", app.Validate)
	fmt.Println("ScheduleURL: ", app.ScheduleURL)
	fmt.Println("Header: ", app.Header)
	fmt.Println("FirstLineText: ", app.FirstLineText)
	fmt.Println("Factor: ", app.Factor)
	fmt.Println("Bits: ", app.Bits)
	fmt.Println("ScheduleWidth: ", app.ScheduleWidth)
	fmt.Println("PN: ", app.PN)
	fmt.Println("OH: ", app.OH)
	fmt.Println("OO: ", app.OO)
	fmt.Println("RO: ", app.RO)
	fmt.Println("QTY: ", app.QTY)
	fmt.Println("TLI: ", app.TLI)
	fmt.Println("TOKI: ", app.TOKI)
	fmt.Println("UOM: ", app.UOM)

}
