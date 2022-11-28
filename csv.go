package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var (
	ErrEOF                error = errors.New("EOF")
	ErrFilterIncompatible error = errors.New("FilterIncompatible")
	ErrFilterMismatch     error = errors.New("FilterMismatch")
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
			err = ErrEOF
		}
		return
	}

	text := c.scanner.Text()
	row = strings.Split(text, ",")
	return
}

func (c *CSVHandler) NextRowWithFilter(filter []string) (row []string, err error) {
	row, err = c.NextRow()
	if err != nil {
		return
	}

	if len(filter) != len(row) {
		err = ErrFilterMismatch
		return
	}

	for i, arg := range filter {
		match, _ := regexp.MatchString(arg, row[i])
		if !match {
			err = ErrFilterIncompatible
			return
		}
	}
	return
}
