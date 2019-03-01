package main

import (
	"fmt"
	"os"

	"github.com/ninedraft/flex/pkg/cli"
)

func main() {
	if err := cli.Main().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
