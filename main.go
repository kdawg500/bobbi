package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kdawg500/bobbi/link"
)

func main() {
	start := time.Now()
	args := os.Args[1:]
	link.Validate(args[0])

	duration := time.Since(start)
	fmt.Println(duration)
}
