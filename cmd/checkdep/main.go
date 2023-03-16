package main

import (
	"checkdep"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkdep.Analyzer) }
