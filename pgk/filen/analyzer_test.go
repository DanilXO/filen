package filen

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

type silentTest struct {
	*testing.T
	Errors []string
}

func (t *silentTest) Errorf(format string, args ...interface{}) {
	t.Errors = append(t.Errors, fmt.Sprintf(format, args...))
}

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testCases := []struct {
		name                string
		runConfig           *Runner
		countOfInvalidFiles int
	}{
		{
			name: "all_files_are_valid",
			runConfig: &Runner{
				MaxLinesNum: 500,
				MinLinesNum: 1,
			},
			countOfInvalidFiles: 0,
		},
		{
			name: "one_go_is_too_big",
			runConfig: &Runner{
				MaxLinesNum: 10,
				MinLinesNum: 1,
			},
			countOfInvalidFiles: 1,
		},
		{
			name: "two_go_is_too_small",
			runConfig: &Runner{
				MaxLinesNum: 500,
				MinLinesNum: 6,
			},
			countOfInvalidFiles: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			silentTest := &silentTest{T: t}
			testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
			analysistest.Run(silentTest, testdata, NewAnalyzer(testCase.runConfig), "samples")

			require.Len(t, silentTest.Errors, testCase.countOfInvalidFiles)
		})
	}
}
