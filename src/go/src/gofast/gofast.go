package gofast

import (
	"fmt"
	// "go/ast"
	"go/parser"
	// "go/printer"
	"go/token"
	"log"
	// "os"
)

func Gofast( code string ) string {
  fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("########### Manual Iteration ###########")

  return code
}
