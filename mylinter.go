package mylinter

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "mylinter is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "mylinter",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}

	reports := map[token.Pos]bool{}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if ifStmt, ok := n.(*ast.IfStmt); ok {
			if assignStmt, ok := ifStmt.Init.(*ast.AssignStmt); ok {
				if callExpr, ok := assignStmt.Rhs[0].(*ast.CallExpr); ok {
					if ident, ok := callExpr.Fun.(*ast.Ident); ok {
						if ident.Name == "Translate" {
							reports[ident.NamePos] = false
							if ifSt, ok := ifStmt.Body.List[0].(*ast.IfStmt); ok {
								if ce, ok := ifSt.Cond.(*ast.CallExpr); ok {
									if ide, ok := ce.Fun.(*ast.Ident); ok {
										if ide.Name == "IsTypeError" {
											reports[ident.NamePos] = true
											// pass.Reportf(ide.NamePos, "is not checke...")
										} else {
											reports[ident.NamePos] = false
											pass.Reportf(ide.NamePos, "not checked...")
										}
									}
								}
							}
						}
					}
				}
			}
		}
		// switch n := n.(type) {
		// case *ast.Ident:
		// 	if n.Name == "gopher" {
		// 		pass.Reportf(n.Pos(), "identifier is gopher")
		// 	}
		// }
	})

	fmt.Println("----------------------------------")
	for _, ok := range reports {
		if !ok {
		}
	}

	return nil, nil
}
