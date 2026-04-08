package modernfor_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/smeshkov/golangci-lint-plugins/modernfor"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, modernfor.Analyzer, "example")
}
