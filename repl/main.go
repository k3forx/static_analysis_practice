package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/constant"
	"go/parser"
)

func evalBinaryExpr(expr *ast.BinaryExpr) (constant.Value, error) {

	xLit, ok := expr.X.(*ast.BasicLit)
	if !ok {
		return constant.MakeUnknown(), errors.New("left operand is not BasicLit")
	}

	yLit, ok := expr.Y.(*ast.BasicLit)
	if !ok {
		return constant.MakeUnknown(), errors.New("right operand is not BasicLit")
	}

	x := evalBasicLit(xLit)
	y := evalBasicLit(yLit)
	return constant.BinaryOp(x, expr.Op, y), nil
}

func evalBasicLit(expr *ast.BasicLit) constant.Value {
	return constant.MakeFromLiteral(expr.Value, expr.Kind, 0)
}

func main() {
	expr, err := parser.ParseExpr("2*10")
	if err != nil {
		panic(err)
	}

	if be, ok := expr.(*ast.BinaryExpr); ok {
		if v, err := evalBinaryExpr(be); err == nil {
			fmt.Println(v)
		} else {
			fmt.Println("Error:", err)
		}
	}

}
