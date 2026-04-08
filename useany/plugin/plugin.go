package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/smeshkov/golangci-lint-plugins/useany"
)

func init() {
	register.Plugin("useany", New)
}

func New(conf any) (register.LinterPlugin, error) { //nolint:ireturn
	return &Plugin{}, nil
}

type Plugin struct{}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{useany.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
