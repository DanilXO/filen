package main

import (
	"github.com/DanilXO/filen/pkg/filen"
	"golang.org/x/tools/go/analysis/singlechecker"
)

const (
	defaultMaxLines       = 500
	defaultMinLines       = 5
	defaultIgnoreComments = false
)

func main() {
	r := &filen.Runner{}

	r.FlagSet.IntVar(&r.MaxLines, "maxLines", defaultMaxLines, "Maximum number of lines in a file")
	r.FlagSet.IntVar(&r.MinLines, "minLines", defaultMinLines, "Minimum number of lines in a file")
	r.FlagSet.BoolVar(&r.IgnoreComments, "ignoreComments", defaultIgnoreComments, "Ignore comment lines or not")

	singlechecker.Main(filen.NewAnalyzer(r))
}
