package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func makeReadme(filename string) error {
	date := time.Now().Format("2 Jan 2006")

	updated := "<sub>Last updated by GitHub Action on " + date + ".</sub>"
	data := fmt.Sprintf("%s\n", updated)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	makeReadme("../README.md")
}
