package config

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
