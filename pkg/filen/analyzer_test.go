package filen_test

import (
	"testing"

	"github.com/DanilXO/filen/pkg/filen"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestName(t *testing.T) {
	testdata := analysistest.TestData()

	testCases := []struct {
		desc   string
		config *filen.Runner
		pkg    string
	}{
		{
			desc: "no report",
			pkg:  "a",
			config: &filen.Runner{
				MaxLines:       500,
				MinLines:       5,
				IgnoreComments: false,
			},
		},
		{
			desc: "default max lines exceeded",
			pkg:  "b",
			config: &filen.Runner{
				MaxLines:       500,
				MinLines:       5,
				IgnoreComments: false,
			},
		},
		{
			desc: "default min lines exceeded",
			pkg:  "c",
			config: &filen.Runner{
				MaxLines:       500,
				MinLines:       5,
				IgnoreComments: false,
			},
		},
		{
			desc: "no report when ignore comments",
			pkg:  "d",
			config: &filen.Runner{
				MaxLines:       500,
				MinLines:       5,
				IgnoreComments: true,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			analysistest.RunWithSuggestedFixes(t, testdata, filen.NewAnalyzer(test.config), test.pkg)
		})
	}
}
