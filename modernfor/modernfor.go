// Package modernfor provides a linter that enforces Go 1.22 integer range loop syntax.
package modernfor

import (
	"errors"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var errInspectNotFound = errors.New("inspect analyzer not found")

// Analyzer exports the linter so golangci-lint can use it.
var Analyzer = &analysis.Analyzer{
	Name:     "modernfor",
	Doc:      "enforces Go 1.22 integer range loop syntax (for i := range N)",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errInspectNotFound
	}

	// We only care about traditional `for` loops (not `for ... range`)
	nodeFilter := []ast.Node{
		(*ast.ForStmt)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		forStmt, ok := n.(*ast.ForStmt)
		if !ok {
			return
		}

		if isCountingLoop(forStmt) {
			pass.Reportf(forStmt.Pos(), "use Go 1.22 'for ... := range N' syntax instead of traditional for-loop")
		}
	})

	return nil, nil
}

// isCountingLoop checks whether a for statement is a traditional counting loop
// (e.g., `for i := 0; i < N; i++`) that could use Go 1.22 range syntax.
func isCountingLoop(forStmt *ast.ForStmt) bool {
	if forStmt.Init == nil || forStmt.Cond == nil || forStmt.Post == nil {
		return false
	}

	if _, ok := forStmt.Init.(*ast.AssignStmt); !ok {
		return false
	}

	cond, condOk := forStmt.Cond.(*ast.BinaryExpr)
	if !condOk {
		return false
	}

	inc, incOk := forStmt.Post.(*ast.IncDecStmt)
	if !incOk {
		return false
	}

	return cond.Op == token.LSS && inc.Tok == token.INC
}
