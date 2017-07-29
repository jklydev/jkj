package main

import (
	"bytes"
	"log"
	sc "strconv"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dateString(t time.Time) string {
	var datestring bytes.Buffer
	datestring.WriteString(sc.Itoa(t.Year()))
	datestring.WriteString(prependZero(int(t.Month())))
	datestring.WriteString(prependZero(t.Day()))

	return datestring.String()
}

func prependZero(timeint int) string {
	timestring := sc.Itoa(timeint)
	if len(timestring) == 2 {
		return timestring
	} else {
		return "0" + timestring
	}
}
