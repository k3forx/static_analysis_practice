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

	// 型チェックを行うための設定
	config := &types.Config{
		Importer: importer.Default(),
	}

	// 型チェックを行う
	_, err = config.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal("Error:", err)
	}

	var res bool
	var cnt int
	ast.Inspect(f, func(n ast.Node) bool {
		// 識別子ではない場合は無視
		blockStmt, ok := n.(*ast.BlockStmt)
		if !ok {
			return true
		}
		if len(blockStmt.List) == 1 {
			return true
		}

		var targetIndex int
		for index, stmt := range blockStmt.List {
			assignStmt, ok := stmt.(*ast.AssignStmt)
			if !ok {
				continue
			}
			callExpr, ok := assignStmt.Rhs[0].(*ast.CallExpr)
			if !ok {
				continue
			}
			ident, ok := callExpr.Fun.(*ast.Ident)
			if !ok {
				continue
			}
			if ident.Name == "Translate" {
				targetIndex = index
				break
			}
		}

		targetStmt := blockStmt.List[targetIndex+1]
		ast.Inspect(targetStmt, func(n ast.Node) bool {
			ifStmt, ok := targetStmt.(*ast.IfStmt)
			if !ok {
				return true
			}

			for _, l := range ifStmt.Body.List {
				ifStmt, ok := l.(*ast.IfStmt)
				if !ok {
					return true
				}
				callExpr, ok := ifStmt.Cond.(*ast.CallExpr)
				if !ok {
					return true
				}
				ident, ok := callExpr.Fun.(*ast.Ident)
				if !ok {
					return true
				}
				if ident.Name == "IsTypeError" {
					cnt++
					res = true
					return true
				}
			}
			return true
		})

		return true
	})

	fmt.Println("------------------")
	fmt.Println(res, cnt)
	fmt.Println("finish")
}
