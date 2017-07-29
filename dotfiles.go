package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func getJournalPath() (journal_path string) {
	options := readDotfile()
	journal_path = options.journal_path
	return journal_path
}

func readDotfile() (options Options) {
	path := os.ExpandEnv("$HOME/.jkj")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		makeDotfile(path)
	}
	options = parseDotfile(path)
	return options
}

func makeDotfile(path string) {
	file, err := os.Create(path)
	check(err)
	file.WriteString("journal_path: $HOME/.journal/")
	file.Close()
}

func parseDotfile(path string) (options Options) {
	source, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(source, &options)
	check(err)
	return options
}

type Options struct {
	journal_path string
}
