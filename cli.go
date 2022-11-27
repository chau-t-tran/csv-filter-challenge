package main

import (
	"flag"
)

type Mode int

const (
	IMPLICIT Mode = iota
)

type CLI struct {
	path string
	mode Mode
	args []string
}

func NewCLI() *CLI {
	pathPtr := flag.String("f", "default", ".csv file path")
	implicitPtr := flag.Bool("i", false, "Runs the CLI in implicit mode")
	flag.Parse()
	args := flag.Args()

	var mode Mode

	if *implicitPtr {
		mode = IMPLICIT
	}

	return &CLI{
		path: *pathPtr,
		mode: mode,
		args: args,
	}
}
