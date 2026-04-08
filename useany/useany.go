// Package useany provides a linter that enforces using 'any' instead of 'interface{}'.
package useany

import (
	"errors"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var errInspectNotFound = errors.New("inspect analyzer not found")

// Analyzer exports the linter so golangci-lint can use it.
var Analyzer = &analysis.Analyzer{
	Name:     "useany",
	Doc:      "enforces using 'any' instead of 'interface{}'",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errInspectNotFound
	}

	nodeFilter := []ast.Node{
		(*ast.InterfaceType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		iface, ok := n.(*ast.InterfaceType)
		if !ok {
			return
		}

		// An empty interface has zero methods.
		if iface.Methods != nil && iface.Methods.NumFields() != 0 {
			return
		}

		pass.Reportf(iface.Pos(), "use 'any' instead of 'interface{}'")
	})

	return nil, nil
}
