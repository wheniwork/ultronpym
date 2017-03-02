package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func success(v ...interface{}) {
	color.New(color.FgHiGreen).Println(v...)
}
func fatal(v ...interface{}) {
	color.New(color.FgRed, color.Bold).Println(v...)
	os.Exit(1)
}

func printHash(file string, hash string, changed bool) {
	green := color.New(color.FgMagenta).SprintfFunc()
	if changed {
		blue := color.New(color.FgHiBlue).SprintfFunc()
		fmt.Printf("%s: %s %s\n", green(file), hash, blue("[changed]"))
	} else {
		fmt.Printf("%s: %s\n", green(file), hash)
	}
}
