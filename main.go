package main

import (
	"fmt"
	"os"
	"os/exec"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("I'm sorry, what?")
	}
}
