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

	file, err := parser.ParseFile(fset, "_example.go", nil, 0)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	_, err = config.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		log.Fatal("Error:", err)
	}

	p, err := config.Importer.Import("context")
	if err != nil {
		log.Fatal("Error:", err)
	}
	ctxType := p.Scope().Lookup("Context").Type()
	it, ok := ctxType.Underlying().(*types.Interface)
	if !ok {
		log.Fatal("should be found Context interface")
	}

	ast.Inspect(file, func(n ast.Node) bool {
		f, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		params := f.Type.Params.List
		for _, p := range params {
			// ast.Print(fset, p)
			varIdent := p.Names[0]

			obj := info.ObjectOf(varIdent)
			if obj == nil {
				return true
			}

			if types.Implements(obj.Type(), it) && (varIdent.Name != "ctx") {
				fmt.Printf("%s: variable name of context.Context is invalid\n", fset.Position(varIdent.Pos()))
			}
		}
		return true
	})
}
