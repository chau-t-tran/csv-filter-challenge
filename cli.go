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

	nameArgMap := make(map[string]string, 10)
	var filterArgs []string
	var mode Mode

	colNames, err := csvHandler.NextRow()
	if err != nil {
		return nil, err
	}

	if *implicitPtr {
		mode = IMPLICIT
		args := flag.Args()
		for i, name := range colNames {
			if i >= len(args) {
				nameArgMap[name] = "*"
				continue
			}
			nameArgMap[name] = args[i]
		}
	}

	if *explicitPtr {
		mode = EXPLICIT
		for _, arg := range flag.Args() {
			split := strings.Split(arg, "=")
			nameArgMap[split[0]] = split[1]
		}
	}

	if *promptPtr {
		mode = PROMPT
		scanner := bufio.NewScanner(os.Stdin)
		for _, name := range colNames {
			fmt.Printf("Enter filter for %s: ", name)
			if scanner.Scan() {
				nameArgMap[name] = scanner.Text()
			}
			fmt.Printf("\n")
		}
	}

	for _, name := range colNames {
		arg, ok := nameArgMap[name]
		if !ok {
			filterArgs = append(filterArgs, "*")
		}
		filterArgs = append(filterArgs, arg)
	}

	return &CLI{
		path:       *pathPtr,
		mode:       mode,
		args:       filterArgs,
		csvHandler: &csvHandler,
	}, nil
}

func (c *CLI) Run() {
	for {
		row, err := c.csvHandler.NextRowWithFilter(c.args)
		if err != nil {
			if err == ErrEOF {
				break
			}
			continue
		}
		fmt.Println(strings.Join(row, ","))
	}
}
