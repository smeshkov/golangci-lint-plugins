package useany

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer exports the linter so golangci-lint can use it.
var Analyzer = &analysis.Analyzer{
	Name:     "useany",
	Doc:      "enforces using 'any' instead of 'interface{}'",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.InterfaceType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		iface := n.(*ast.InterfaceType)

		// An empty interface has zero methods.
		if iface.Methods != nil && iface.Methods.NumFields() != 0 {
			return
		}

		pass.Reportf(iface.Pos(), "use 'any' instead of 'interface{}'")
	})

	return nil, nil
}
