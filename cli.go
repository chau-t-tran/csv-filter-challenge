package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Mode int

const (
	IMPLICIT Mode = iota
	EXPLICIT
	PROMPT
)

type CLI struct {
	path       string
	mode       Mode
	args       []string
	csvHandler *CSVHandler
}

func NewCLI() (*CLI, error) {
	pathPtr := flag.String("f", "default", ".csv file path")
	implicitPtr := flag.Bool("i", false, "Runs the CLI in implicit mode")
	explicitPtr := flag.Bool("e", false, "Runs the CLI in explicit mode")
	promptPtr := flag.Bool("p", false, "Runs the CLI in prompt mode")
	flag.Parse()

	csvHandler, err := NewCSVHandler(*pathPtr)
	if err != nil {
		return nil, err
	}

	var filterArgs []string
	var mode Mode

	if *implicitPtr {
		mode = IMPLICIT
		filterArgs = flag.Args()
	}

	if *explicitPtr {
		mode = EXPLICIT
		nameArgMap := make(map[string]string, 10)
		for _, arg := range flag.Args() {
			split := strings.Split(arg, "=")
			nameArgMap[split[0]] = split[1]
		}

		colNames, err := csvHandler.NextRow()
		if err != nil {
			return nil, err
		}
		for _, name := range colNames {
			filterArgs = append(filterArgs, nameArgMap[name])
		}
	}

	if *promptPtr {
		mode = PROMPT
		scanner := bufio.NewScanner(os.Stdin)

		colNames, err := csvHandler.NextRow()
		if err != nil {
			return nil, err
		}
		for _, name := range colNames {
			fmt.Printf("Enter filter for %s: ", name)
			if scanner.Scan() {
				filterArgs = append(filterArgs, scanner.Text())
			}
			fmt.Printf("\n")
		}
	}

	return &CLI{
		path: *pathPtr,
		mode: mode,
		args: filterArgs,
	}, nil
}
