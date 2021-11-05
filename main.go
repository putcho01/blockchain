package main

import (
	"os"

	"blockchain-sample/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}

	cmd.Run()

}
