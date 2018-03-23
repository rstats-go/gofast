package gofast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func Gofast( code string ) []string {
  fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

  functions := []string{} ;

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Functions
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			functions = append( functions, fn.Name.Name )
		}
		return true
	})

  return functions
}
