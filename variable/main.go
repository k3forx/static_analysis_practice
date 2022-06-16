package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "_test1.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// 識別子が定義または利用されてる部分を記録する
	defsOrUses := map[*ast.Ident]types.Object{}
	info := &types.Info{
		Defs: defsOrUses,
		Uses: defsOrUses,
	}

	// // 型チェックを行うための設定
	config := &types.Config{
		Importer: importer.Default(),
	}

	_, err = config.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal("Error:", err)
	}

	result := map[token.Position]bool{}
	ast.Inspect(f, func(n ast.Node) bool {
		if ifStmt, ok := n.(*ast.IfStmt); ok {
			if assignStmt, ok := ifStmt.Init.(*ast.AssignStmt); ok {
				if callExpr, ok := assignStmt.Rhs[0].(*ast.CallExpr); ok {
					if ident, ok := callExpr.Fun.(*ast.Ident); ok {
						if ident.Name == "Translate" {
							result[fset.Position(ident.Pos())] = false
							if ifSt, ok := ifStmt.Body.List[0].(*ast.IfStmt); ok {
								if ce, ok := ifSt.Cond.(*ast.CallExpr); ok {
									if ide, ok := ce.Fun.(*ast.Ident); ok {
										if ide.Name == "IsTypeError" {
											result[fset.Position(ident.Pos())] = true
											return true
										} else {
											return true
										}
									}
								}
							}
							return true
						}
					}
				}
			}
		}

		return true
	})

	for pos, ok := range result {
		if !ok {
			fmt.Printf("%v is not checked!\n", pos)
		}
	}
}
