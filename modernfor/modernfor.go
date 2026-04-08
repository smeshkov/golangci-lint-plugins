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
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errInspectNotFound
	}

	// We only care about traditional `for` loops (not `for ... range`)
	nodeFilter := []ast.Node{
		(*ast.ForStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		forStmt, ok := n.(*ast.ForStmt)
		if !ok {
			return
		}

		// A traditional counting loop has an Init, Cond, and Post statement.
		if forStmt.Init == nil || forStmt.Cond == nil || forStmt.Post == nil {
			return
		}

		// 1. Check Init: e.g., `i := 0`
		_, okInit := forStmt.Init.(*ast.AssignStmt)
		// 2. Check Cond: e.g., `i < N`
		cond, okCond := forStmt.Cond.(*ast.BinaryExpr)
		// 3. Check Post: e.g., `i++`
		inc, okPost := forStmt.Post.(*ast.IncDecStmt)

		if !okInit || !okCond || !okPost {
			return
		}

		// Simplified check: If the condition uses `<` and the post increments `++`,
		// flag it as a candidate for the new Go 1.22 range syntax.
		if cond.Op == token.LSS && inc.Tok == token.INC {
			pass.Reportf(forStmt.Pos(), "use Go 1.22 'for ... := range N' syntax instead of traditional for-loop")
		}
	})

	return nil, nil
}
