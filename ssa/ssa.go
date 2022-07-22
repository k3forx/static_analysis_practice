package ssa

import (
	"fmt"
	"go/token"
	"go/types"
	"os"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "ssa is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "ssa",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs

	for _, f := range funcs {
		f.WriteTo(os.Stdout)
		for _, b := range f.Blocks {
			if v := binOpErrNil(b, token.NEQ); v != nil {
				fmt.Printf("v: %+v\n", b.Succs[0])
				caller := getIfCallFunc(b.Succs[0])
				fmt.Printf("caller: %+v\n", caller)
			}
			fmt.Println("--------------------")
		}
	}

	return nil, nil
}

var errType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func binOpErrNil(b *ssa.BasicBlock, op token.Token) ssa.Value {
	if len(b.Instrs) == 0 {
		return nil
	}

	ifInst, ok := b.Instrs[len(b.Instrs)-1].(*ssa.If)
	if !ok {
		return nil
	}

	binOp, ok := ifInst.Cond.(*ssa.BinOp)
	if !ok {
		return nil
	}

	if binOp.Op != op {
		return nil
	}

	fmt.Printf("X: %+v, Y: %+v\n", binOp.X.Type(), binOp.Y.Type())

	if !types.Implements(binOp.X.Type(), errType) {
		return nil
	}

	if !types.Implements(binOp.Y.Type(), errType) {
		return nil
	}

	xIsConst, yIsConst := isConst(binOp.X), isConst(binOp.Y)
	fmt.Println(xIsConst, yIsConst)
	switch {
	case !xIsConst && yIsConst: // err != nil (op: token.NEQ) or err == nil
		return binOp.X
	case xIsConst && !yIsConst: // nil != err (op: token.NEQ) or nil == err
		return binOp.Y
	}

	return nil
}

func isConst(v ssa.Value) bool {
	_, ok := v.(*ssa.Const)
	return ok
}

func getIfCallFunc(b *ssa.BasicBlock) *ssa.Call {
	if len(b.Instrs) == 0 {
		return nil
	}

	for _, inst := range b.Instrs {
		call, ok := inst.(*ssa.Call)
		if !ok {
			continue
		} else {
			return call
		}
	}
	return nil
}
