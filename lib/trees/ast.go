package trees

import (
	"go/ast"
	"go/parser"
	"go/token"
	tb "lib/table"
)

type ASTNode interface{}

type TableLeaf struct {
	TableName string
	Info      tb.TableInfo
}

func (m *TableLeaf) GetTableInfo() tb.TableInfo {
	return m.Info
}

type BinaryOp struct {
	Op     token.Token
	Left   ASTNode
	Right  ASTNode
	Fields []string
}

func ParseExpr(expr string) ASTNode {
	fset := token.NewFileSet()
	defAst, _ := parser.ParseExprFrom(fset, "", expr, 0)
	tree := parseAst(defAst)
	ast.Print(fset, tree)
	return tree
}

func parseAst(n ast.Node) ASTNode {
	switch x := n.(type) {
	case *ast.ParenExpr:
		return parseAst(x.X)

	case *ast.BinaryExpr:
		left := parseAst(x.X)
		right := parseAst(x.Y)

		newNode := BinaryOp{
			Op:    x.Op,
			Left:  left,
			Right: right,
		}

		return &newNode

	case *ast.Ident:
		newNode := TableLeaf{
			TableName: x.Name,
		}
		return &newNode
	case *ast.IndexListExpr:
		fields := []string{}
		for i := 0; i < len(x.Indices); i++ {
			fields = append(fields, x.Indices[i].(*ast.Ident).Name)
		}
		left := parseAst(x.X)
		newNode := BinaryOp{
			Left:   left,
			Fields: fields,
		}
		return &newNode
	case *ast.IndexExpr:
		fields := []string{}
		fields = append(fields, x.Index.(*ast.Ident).Name)
		left := parseAst(x.X)
		newNode := BinaryOp{
			Left:   left,
			Fields: fields,
		}
		return &newNode
	}
	return nil
}
