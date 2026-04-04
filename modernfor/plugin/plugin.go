package plugin

import (
	"golang.org/x/tools/go/analysis"

	"github.com/smeshkov/golangci-lint-plugins/modernfor"
)

// New must be exported and return a slice of analyzers.
// This is the specific signature golangci-lint looks for when building custom modules.
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{modernfor.Analyzer}, nil
}
