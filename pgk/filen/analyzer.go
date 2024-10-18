package filen

import (
	"flag"
	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var (
	maxLinesNum int
	minLinesNum int
)

const (
	defaultMaxLinesNum = 500
	defaultMinLinesNum = 5
)

func NewAnalyzer() *analysis.Analyzer {
	var flagSet flag.FlagSet

	flagSet.IntVar(&maxLinesNum, "maxLinesNum", defaultMaxLinesNum, "Maximum number of lines in a file")
	flagSet.IntVar(&minLinesNum, "minLinesNum", defaultMinLinesNum, "Minimum number of lines in a file")

	return &analysis.Analyzer{
		Name:  "filen",
		Doc:   "checks files size",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		fileLen := pass.Fset.Position(f.End()).Line
		fileName := pass.Fset.Position(f.Pos()).Filename

		if fileLen > maxLinesNum {
			pass.Reportf(f.Pos(), "The number of lines in the file %s exceeds the allowed value! maxLinesNum = %d, fileLines = %d",
				fileName, maxLinesNum, fileLen)
		}
		if fileLen < minLinesNum {
			pass.Reportf(f.Pos(), "The number of lines in the file %s less the allowed value! minLinesNum = %d, fileLines = %d",
				fileName, minLinesNum, fileLen)
		}
	}
	return nil, nil
}
