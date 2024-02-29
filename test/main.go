package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	defAst, _ := parser.ParseExprFrom(fset, "", "(A + B)[id, name]", 0)
	ast.Print(fset, defAst)
}
