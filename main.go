package main

import (
	"fmt"
	"os"

	"github.com/wheniwork/ultronpym/cmd"
)

func main() {
	if err := cmd.Commands().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
