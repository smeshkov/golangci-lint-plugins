package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/smeshkov/golangci-lint-plugins/modernfor"
)

func init() {
	register.Plugin("modernfor", New)
}

func New(conf any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

type Plugin struct{}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{modernfor.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
