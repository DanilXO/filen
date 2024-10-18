package main

import (
	"filen/pgk/filen"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(filen.NewAnalyzer()) }
