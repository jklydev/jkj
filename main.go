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
	file.WriteString(beginFile(t))
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
	heading.WriteString("\n\n")

	return heading.String()
}

func subHeading(t time.Time) string {
	var subheading bytes.Buffer

	subheading.WriteString(sc.Itoa(t.Hour()))
	subheading.WriteString(":")
	subheading.WriteString(sc.Itoa(t.Minute()))
	subheading.WriteString("\n------------")
	subheading.WriteString("\n")

	return subheading.String()
}
