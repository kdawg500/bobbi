package main

import (
	"os"
	"time"

	"github.com/fatih/color"
)

var visited = map[string]bool{}

func main() {
	start := time.Now()
	args := os.Args[1:]
	input := args[0]

	duration := time.Since(start)
	color.Red(input)
	color.Cyan(duration.String())
}
