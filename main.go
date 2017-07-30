package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"time"
)

func main() {
	journal_path := getJournalPath()

	var entry string
	flag.StringVar(&entry, "e", "", "A journal entry")
	flag.Parse()

	t := time.Now()
	path := journal_path + dateString(t)

	file := getFile(path, t)
	if entry != "" {
		writeFlagEntry(entry, *file)
	} else {
		openEditor(path)
	}
	file.Close()
}

func getFile(path string, t time.Time) (file *os.File) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		check(err)
		file.WriteString(beginFile(t))
		return file
	} else {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		check(err)
		file.WriteString(subHeading(t))
		return file
	}
}

func openEditor(path string) {
	cmd := exec.Command("vim", "+ :setf txt", "+norm Go", "+startinsert", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	check(err)
}

func writeFlagEntry(flagstring string, file os.File) {
	var entry bytes.Buffer
	entry.WriteString(flagstring)
	entry.WriteString("\n")

	_, err := file.WriteString(entry.String())
	check(err)

}

func beginFile(t time.Time) string {
	var header bytes.Buffer
	header.WriteString(heading(t))
	header.WriteString(subHeading(t))

	return header.String()
}

func heading(t time.Time) string {
	var heading bytes.Buffer

	heading.WriteString("# ")
	heading.WriteString(prependZero(t.Day()))
	heading.WriteString(" ")
	heading.WriteString(t.Month().String())
	heading.WriteString("\n")

	return heading.String()
}

func subHeading(t time.Time) string {
	var subheading bytes.Buffer

	subheading.WriteString("\n")
	subheading.WriteString("## ")
	subheading.WriteString(prependZero(t.Hour()))
	subheading.WriteString(":")
	subheading.WriteString(prependZero(t.Minute()))
	subheading.WriteString("\n")

	return subheading.String()
}
