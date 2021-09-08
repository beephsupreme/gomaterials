package utility

import (
	"encoding/csv"
	"log"
	"os"
)

// LoadFile reads a csv file and returns it as a [][]string
func LoadFile(filename string) [][]string {
	f, err := os.Open(filename)
	CheckError("os.Open():", err)
	defer func(f *os.File) {
		err := f.Close()
		CheckError("f.Close():", err)
	}(f)
	r := csv.NewReader(f)
	t, err := r.ReadAll()
	CheckError("d.ReadAll()", err)
	return t
}

// CheckError checks for an error and halts program if error found
func CheckError(source string, err error) {
	if err != nil {
		log.Fatal(source, err.Error())
	}
}
