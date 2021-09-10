package utility

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// LoadFile reads a csv file and returns it as a [][]string
func LoadFile(filename string) [][]string {
	fmt.Println("Loading file:", filepath.Base(filename))
	f, err := os.Open(filename)
	CheckError("[utility.LoadFile.Open()] ", err)
	defer func(f *os.File) {
		err := f.Close()
		CheckError("[utility.LoadFile.Close()] ", err)
	}(f)
	r := csv.NewReader(f)
	t, err := r.ReadAll()
	CheckError("[utility.LoadFile.ReadAll()] ", err)
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
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	//fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("Memory used: %v MiB\n", bToMb(m.TotalAlloc))
	//fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
