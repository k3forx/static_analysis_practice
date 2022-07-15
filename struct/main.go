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
	// ファイルごとのトークンの位置を記録するFileSetを作成する
	fset := token.NewFileSet()

	// ファイル単位で構文解析を行う
	f, err := parser.ParseFile(fset, "_struct.go", nil, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// 抽象構文木を深さ優先で探索する
	ast.Inspect(f, func(n ast.Node) bool {

		// 識別子ではない場合は無視
		st, ok := n.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range st.Fields.List {
			fieldName := field.Names[0]
			tag := field.Tag.Value

			s := tag
			for _, str := range []string{"`", "\"", "json", ":"} {
				s = strings.ReplaceAll(s, str, "")
			}

			for _, str := range []string{"_", "-"} {
				if strings.Contains(s, str) {
					if len(s) == 1 {
						continue
					}
					fmt.Printf("%s: invalid JSON tag of %s\n", fset.Position(fieldName.Pos()), fieldName.String())
				}
			}
		}

		return true
	})
}
