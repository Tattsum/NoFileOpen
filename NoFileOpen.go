package NoFileOpen

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "NoFileOpen checks bad used os.FileOpen"

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
		}
	}

	if ospkg == nil {
		return nil, nil
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node){
		switch n := n.(type) {
		case *ast.CallExpr:

			// check function is os.OpenFile
			caller, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}

			if pass.TypesInfo.ObjectOf(caller.Sel).Pkg() != ospkg {
				return
			}

			if caller.Sel.Name != "OpenFile" {
				return
			}

			pass.Reportf(caller.Pos(), "don't use os.OpenFile. Use os.Create")
		}
	})


	return nil, nil
}
