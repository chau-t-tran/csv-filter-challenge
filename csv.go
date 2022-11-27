package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type CSVHandler struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewCSVHandler(path string) (c CSVHandler, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)

	return CSVHandler{file: file, scanner: scanner}, nil
}

func (c *CSVHandler) NextRow() (row []string, err error) {
	if !c.scanner.Scan() {
		err = c.scanner.Err()
		if err == nil {
			err = errors.New("EOF")
		}
		return
	}

	text := c.scanner.Text()
	row = strings.Split(text, ",")
	return
}
