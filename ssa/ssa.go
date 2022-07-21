package ssa

import (
	"os"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
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
		// for _, b := range f.Blocks {
		// 	for _, instr := range b.Instrs {
		// 		fmt.Printf("%+v\n", instr)
		// 		fmt.Println()
		// 	}
		// }
	}

	return nil, nil
}
