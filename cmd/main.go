package main

import (
	"os"

	"github.com/danmxli/tree-walking-interpreter/internal/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
