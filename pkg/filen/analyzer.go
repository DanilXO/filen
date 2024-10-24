package filen

import (
	"flag"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type Runner struct {
	FlagSet        flag.FlagSet
	MaxLines       int
	MinLines       int
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

		if fileLen > cfg.MaxLines {
			pass.Reportf(f.Pos(), "The number of lines exceeds the allowed value. (maxLinesNum = %d, fileLines = %d)",
				cfg.MaxLines, fileLen)
		}

		if fileLen < cfg.MinLines {
			pass.Reportf(f.Pos(), "The number of lines in less the allowed value. (minLinesNum = %d, fileLines = %d)",
				cfg.MinLines, fileLen)
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
