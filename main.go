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
	var file os.File
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString(beginFile(t))
	} else {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString(subHeading(t))
	}
	file.Close()
}

func dateString(t time.Time) string {
	var datestring bytes.Buffer
	datestring.WriteString(sc.Itoa(t.Year()))
	datestring.WriteString(sc.Itoa(t.YearDay()))

	return datestring.String()
}

func beginFile(t time.Time) string {
	var header bytes.Buffer
	header.WriteString(heading(t))
	header.WriteString(subHeading(t))

	return header.String()
}

func heading(t time.Time) string {
	var heading bytes.Buffer

	heading.WriteString(sc.Itoa(t.Day()))
	heading.WriteString(" ")
	heading.WriteString(t.Month().String())
	heading.WriteString("\n============")
	heading.WriteString("\n")

	return heading.String()
}

func subHeading(t time.Time) string {
	var subheading bytes.Buffer

	subheading.WriteString("\n")
	subheading.WriteString(sc.Itoa(t.Hour()))
	subheading.WriteString(":")
	subheading.WriteString(sc.Itoa(t.Minute()))
	subheading.WriteString("\n------------")
	subheading.WriteString("\n")

	return subheading.String()
}
