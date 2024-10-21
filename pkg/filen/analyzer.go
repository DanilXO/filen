package filen

import (
	"flag"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type Runner struct {
	FlagSet        flag.FlagSet
	MaxLinesNum    int
	MinLinesNum    int
	IgnoreComments bool
}

func NewAnalyzer(runner *Runner) *analysis.Analyzer {

	return &analysis.Analyzer{
		Name:  "filen",
		Doc:   "checks files size",
		Run:   runner.run,
		Flags: runner.FlagSet,
	}
}

func (cfg *Runner) run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		fileLen := getLengthOfFile(f, pass.Fset, cfg.IgnoreComments)
		fileName := pass.Fset.Position(f.Pos()).Filename

		if fileLen > cfg.MaxLinesNum {
			pass.Reportf(f.Pos(), "The number of lines in the file %s exceeds the allowed value! maxLinesNum = %d, fileLines = %d",
				fileName, cfg.MaxLinesNum, fileLen)
		}
		if fileLen < cfg.MinLinesNum {
			pass.Reportf(f.Pos(), "The number of lines in the file %s less the allowed value! minLinesNum = %d, fileLines = %d",
				fileName, cfg.MinLinesNum, fileLen)
		}

	}
	return nil, nil
}

func getLengthOfFile(file *ast.File, fset *token.FileSet, ignoreComments bool) int {
	fileLen := fset.Position(file.End()).Line

	if ignoreComments {
		return fileLen - len(file.Comments)
	}

	return fileLen
}
