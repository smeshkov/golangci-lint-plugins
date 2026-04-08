// Package plugin registers the useany linter with golangci-lint.
package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/smeshkov/golangci-lint-plugins/useany"
)

func init() {
	register.Plugin("useany", New)
}

// New creates a new useany linter plugin.
//
//nolint:ireturn
func New(_ any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

// Plugin implements the golangci-lint plugin interface for useany.
type Plugin struct{}

// BuildAnalyzers returns the analysis analyzers for useany.
func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{useany.Analyzer}, nil
}

// GetLoadMode returns the load mode required by useany.
func (p *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
