package main

import (
	"bytes"
	"log"
	"os"
	sc "strconv"
	"time"
)

func main() {
	journal_path := "/Users/jkiely/.journal/"
	t := time.Now()
	path := journal_path + dateString(t)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func dateString(t time.Time) string {
	var datestring bytes.Buffer
	datestring.WriteString(sc.Itoa(t.Year()))
	datestring.WriteString(sc.Itoa(t.YearDay()))

	return datestring.String()
}
