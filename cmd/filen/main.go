package main

import (
	"filen/pgk/filen"
	"golang.org/x/tools/go/analysis/singlechecker"
)

const (
	defaultMaxLinesNum = 500
	defaultMinLinesNum = 5
)

func main() {
	r := &filen.Runner{}
	r.FlagSet.IntVar(&r.MaxLinesNum, "maxLinesNum", defaultMaxLinesNum, "Maximum number of lines in a file")
	r.FlagSet.IntVar(&r.MinLinesNum, "minLinesNum", defaultMinLinesNum, "Minimum number of lines in a file")
	singlechecker.Main(filen.NewAnalyzer(r))
}
