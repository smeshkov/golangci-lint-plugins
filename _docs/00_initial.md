Here is a complete example of how to build a custom `golangci-lint` plugin to enforce the Go 1.22+ `for i := range N` syntax over the traditional `for i := 0; i < N; i++` loop. 

We will use the **Module Plugin System** (custom runner) as it is the most reliable and modern way to extend `golangci-lint`.

### Step 1: Initialize the Go Module
Create a new directory for your linter and initialize a Go module.
```bash
mkdir modernfor-linter
cd modernfor-linter
go mod init example.com/modernfor
go get golang.org/x/tools/go/analysis
```

### Step 2: Write the Analyzer Logic
Custom linters rely on the `golang.org/x/tools/go/analysis` framework. We will inspect the Abstract Syntax Tree (AST) to find `ast.ForStmt` nodes that look like a basic counting loop.

Create `modernfor.go`:

```go
package modernfor

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer exports the linter so golangci-lint can use it.
var Analyzer = &analysis.Analyzer{
	Name:     "modernfor",
	Doc:      "enforces Go 1.22 integer range loop syntax (for i := range N)",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// We only care about traditional `for` loops (not `for ... range`)
	nodeFilter := []ast.Node{
		(*ast.ForStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		forStmt := n.(*ast.ForStmt)

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
```

### Step 3: Expose the Plugin for Golangci-Lint
For `golangci-lint` to inject your module properly using its custom builder, you need to provide an exported function that returns your `Analyzer`. 

Create `plugin/plugin.go`:

```go
package plugin

import (
	"golang.org/x/tools/go/analysis"
	"example.com/modernfor"
)

// New must be exported and return a slice of analyzers.
// This is the specific signature golangci-lint looks for when building custom modules.
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{modernfor.Analyzer}, nil
}
```

### Step 4: Build the Custom `golangci-lint` Binary
Instead of fighting with strict `.so` plugin compilation limits, `golangci-lint` provides a tool called `custom-gcl` to build a custom binary that statically compiles your linter inside.

Install the `custom-gcl` builder tool:
```bash
go install github.com/golangci/golangci-lint/cmd/custom-gcl@latest
```

In the root of the project you want to lint (or within your linter project for testing), create a `.custom-gcl.yml` file. If you are testing locally, use `replace` to point to your local directory:

```yaml
# .custom-gcl.yml
version: v1.56.2 # The base version of golangci-lint you want to use
plugins:
  - module: 'example.com/modernfor'
    import: 'example.com/modernfor/plugin'
    version: v1.0.0
replace:
  - example.com/modernfor=../path/to/your/modernfor-linter # Path to your local module
```

Build the binary:
```bash
custom-gcl build .custom-gcl.yml
```
*This generates a binary named `custom-gcl` in your directory.*

### Step 5: Configure and Run
Now, configure your project's standard `.golangci.yml` file to turn on your new private linter. Because it is injected as a module plugin, `golangci-lint` reads it as if it were natively built in.

```yaml
# .golangci.yml
linters-settings:
  custom:
    modernfor:
      type: module
      description: Enforces Go 1.22 range over int
      original-url: example.com/modernfor

linters:
  enable:
    - modernfor
```

Finally, run your custom binary on your code! 
```bash
./custom-gcl run
```

If it encounters code like `for i := 0; i < 5; i++ { ... }`, it will successfully output:
`main.go:10:2: modernfor: use Go 1.22 'for ... := range N' syntax instead of traditional for-loop`