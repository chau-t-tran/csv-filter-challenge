package main

import "os"

type CSVHandler struct {
	file *os.File
}

func NewCSVHandler(path string) (c CSVHandler, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	return CSVHandler{file: file}, nil
}
