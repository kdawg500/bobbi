package main

import (
	"os"

	"github.com/kdawg500/bobbi/link"
)

func main() {
	args := os.Args[1:]
	u := link.Validate(args[0])
}
