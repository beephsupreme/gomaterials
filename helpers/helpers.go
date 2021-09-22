package helpers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/beephsupreme/gomaterials/internal/config"
)

var app *config.AppConfig

// LoadFile reads a csv file and returns it as a [][]string
func LoadFile(filename string) [][]string {

	fmt.Println("Loading file:", filename)
	f, err := os.Open(app.DataPath + filename)
	CheckError("[helpers.LoadFile.Open()] ", err)
	defer func(f *os.File) {
		err := f.Close()
		CheckError("[helpers.LoadFile.Close()] ", err)
	}(f)
	r := csv.NewReader(f)
	t, err := r.ReadAll()
	CheckError("[helpers.LoadFile.ReadAll()] ", err)
	return t
}

// CheckError checks for an error and halts program if error found
func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory used: %v MiB\n", bToMb(m.TotalAlloc))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func LoadConfig(a *config.AppConfig) {
	app = a
}
