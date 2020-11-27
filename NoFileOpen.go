package NoFileOpen

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "NoFileOpen is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "NoFileOpen",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var ospkg *types.Package

func run(pass *analysis.Pass) (interface{}, error) {
	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Path() == "os" {
			ospkg = pkg
			fmt.Println(ospkg.Path())
		}
	}

	if ospkg == nil {
		return nil, nil
	}

	


	return nil, nil
}
