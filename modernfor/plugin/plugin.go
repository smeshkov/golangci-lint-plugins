// Package plugin registers the modernfor linter with golangci-lint.
package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/smeshkov/golangci-lint-plugins/modernfor"
)

func init() {
	register.Plugin("modernfor", New)
}

// New creates a new modernfor linter plugin.
//
//nolint:ireturn
func New(_ any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

// Plugin implements the golangci-lint plugin interface for modernfor.
type Plugin struct{}

// BuildAnalyzers returns the analysis analyzers for modernfor.
func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{modernfor.Analyzer}, nil
}

// GetLoadMode returns the load mode required by modernfor.
func (p *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
