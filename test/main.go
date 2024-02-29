package main

import (
	"go/ast"
	"go/token"
	tr "lib/trees"
)

func main() {
	defAst := tr.ParseExpr("A*B")

}
