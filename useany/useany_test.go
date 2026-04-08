package useany_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/smeshkov/golangci-lint-plugins/useany"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, useany.Analyzer, "example")
}
