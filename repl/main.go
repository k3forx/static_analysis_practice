package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func eval1pls1(expr *ast.BinaryExpr) (int64, error) {
	xLit, ok := expr.X.(*ast.BasicLit)
	if !ok {
		return 0, errors.New("left operand is not BasicLit")
	}

	yLit, ok := expr.Y.(*ast.BasicLit)
	if !ok {
		return 0, errors.New("right operand is not BasicLit")
	}

	if expr.Op != token.ADD {
		return 0, errors.New("operator is not +")
	}

	x, err := strconv.ParseInt(xLit.Value, 10, 64)
	if err != nil {
		return 0, err
	}

	y, err := strconv.ParseInt(yLit.Value, 10, 64)
	if err != nil {
		return 0, err
	}

	return x + y, nil
}

func main() {
	expr, err := parser.ParseExpr("1+1")
	if err != nil {
		panic(err)
	}

	if be, ok := expr.(*ast.BinaryExpr); ok {
		if v, err := eval1pls1(be); err == nil {
			fmt.Println(v)
		} else {
			fmt.Println("Error:", err)
		}
	}

}
