package main

import (
	"github.com/mahiro72/checkdeps"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkdeps.Analyzer) }
