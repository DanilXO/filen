package filen

import (
	"flag"
	"golang.org/x/tools/go/analysis"
)

type Runner struct {
	FlagSet     flag.FlagSet
	MaxLinesNum int
	MinLinesNum int
}

func NewAnalyzer(runner *Runner) *analysis.Analyzer {

	return &analysis.Analyzer{
		Name:  "filen",
		Doc:   "checks files size",
		Run:   runner.run,
		Flags: runner.FlagSet,
	}
}

func (r *Runner) run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		fileLen := pass.Fset.Position(f.End()).Line
		fileName := pass.Fset.Position(f.Pos()).Filename

		if fileLen > r.MaxLinesNum {
			pass.Reportf(f.Pos(), "The number of lines in the file %s exceeds the allowed value! maxLinesNum = %d, fileLines = %d",
				fileName, r.MaxLinesNum, fileLen)
		}
		if fileLen < r.MinLinesNum {
			pass.Reportf(f.Pos(), "The number of lines in the file %s less the allowed value! minLinesNum = %d, fileLines = %d",
				fileName, r.MinLinesNum, fileLen)
		}
	}
	return nil, nil
}
