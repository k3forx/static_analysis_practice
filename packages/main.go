package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/tools/go/packages"
)

func main() {
	flag.Parse()

	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedName}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Print the names of the source files
	// for each package listed on the command line.
	for _, pkg := range pkgs {
		fmt.Printf("ID: %v, GoFiles: %v, Name: %v\n", pkg.ID, pkg.GoFiles, pkg.Errors)
	}
}
