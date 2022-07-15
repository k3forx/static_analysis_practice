package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func main() {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "_example.go", nil, 0)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		f, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		params := f.Type.Params.List
		for _, p := range params {
			variableName := p.Names[0]
			selector, ok := (p.Type).(*ast.SelectorExpr)
			if !ok {
				return true
			}

			ident, ok := selector.X.(*ast.Ident)
			if !ok {
				return true
			}

			if (ident.Name == strings.ToLower(selector.Sel.Name)) && (variableName.Name != "ctx") {
				fmt.Printf("%s: variable name of context.Context is invalid\n", fset.Position(variableName.Pos()))
			}
		}
		return true
	})
}
