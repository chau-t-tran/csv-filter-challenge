package main

import "fmt"

func main() {
	fmt.Println("Welcome!")
	cli, err := NewCLI()
	fmt.Println(cli)
	if err != nil {
		panic(err)
	}

	cli.Run()
}
