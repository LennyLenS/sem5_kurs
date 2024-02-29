package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	defAst, _ := parser.ParseExprFrom(fset, "", "(A[id, name] * B[id])", 0)
	ast.Print(fset, defAst)
}
