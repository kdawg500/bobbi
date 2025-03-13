package main

import (
	"fmt"
	"os"

	"github.com/kdawg500/bobbi/link"
)

func main() {
	args := os.Args[1:]
	u := link.Validate(args[0])
	fmt.Println(u)
}
