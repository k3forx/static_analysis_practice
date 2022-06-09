package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "_gopher.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal("Error:", err)
	}

	ast.Inspect(f, func(n ast.Node) bool {

		// fmt.Printf("node: %+v\n", n)
		// ast.Print(nil, n)

		// 識別子ではない場合は無視
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}
		fmt.Printf("ident: %+v\n", ident.Name)
		fmt.Printf("detail: %+v\n", ident.IsExported())
		fmt.Println(ident.Name == "Greeting")
		fmt.Println("-----------------------------------")

		// // 識別子がGopherという名前でなければ無視
		// if ident.Name != "Gopher" {
		// 	return true
		// }

		// fmt.Println(fset.Position(ident.Pos()))

		return true
	})
}
