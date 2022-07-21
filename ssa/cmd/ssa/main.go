package main

import (
	"github.com/k3forx/static_analysis_practice/ssa"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ssa.Analyzer) }
