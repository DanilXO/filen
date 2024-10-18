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

	silentTest := &silentTest{T: t}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	result := analysistest.Run(silentTest, testdata, NewAnalyzer(), "samples")
	require.Len(t, result, 1)
	require.Len(t, result[0].Diagnostics, 2)
}
