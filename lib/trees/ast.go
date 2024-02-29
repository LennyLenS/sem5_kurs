package trees

import (
	// "encoding/json"
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
	Op    token.Token
	Left  ASTNode
	Right ASTNode
}

func ParseExpr(expr string) ASTNode {
	fset := token.NewFileSet()
	defAst, _ := parser.ParseExprFrom(fset, "", expr, 0)
	tree := parseGoAstWithoutSize(defAst)
	return tree
}

func parseGoAstWithoutSize(n ast.Node) ASTNode {
	switch x := n.(type) {
	case *ast.ParenExpr:
		return parseGoAstWithoutSize(x.X)

	case *ast.BinaryExpr:
		left := parseGoAstWithoutSize(x.X)
		right := parseGoAstWithoutSize(x.Y)

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
	}
	return nil
}

// func UpdateTreeStats(node ASTNode, data map[string]tb.Table) {
// 	var dfs func(nd ASTNode)
// 	dfs = func(nd ASTNode) {
// 		switch x := nd.(type) {
// 		case *BinaryOp:
// 			dfs(x.Left)
// 			dfs(x.Right)

// 		case *TableLeaf:
// 			x.Info = data[x.TableName].Info
// 		}
// 	}
// 	dfs(node)
// }

// func GetLeafsNames(root ASTNode) map[string]bool {
// 	answr := map[string]bool{}

// 	var dfs func(nd ASTNode)
// 	dfs = func(nd ASTNode) {
// 		switch x := nd.(type) {
// 		case *BinaryOp:
// 			dfs(x.Left)
// 			dfs(x.Right)

// 		case *TableLeaf:
// 			answr[x.TableName] = true
// 		}
// 	}
// 	dfs(root)

// 	return answr
// }

// func UpparseJson(tree json.RawMessage) ASTNode {
// 	var dfs func(json.RawMessage) ASTNode
// 	dfs = func(node json.RawMessage) ASTNode {
// 		var x map[string]json.RawMessage
// 		err := json.Unmarshal(node, &x)
// 		if err != nil {
// 			panic("ошибка при распарсе на воркере(в dfs)")
// 		}

// 		if _, isLeaf := x["TableName"]; isLeaf {
// 			var newLeaf TableLeaf
// 			json.Unmarshal(x["TableName"], &newLeaf.TableName)
// 			json.Unmarshal(x["Info"], &newLeaf.Info)
// 			return &newLeaf
// 		} else {
// 			left := dfs(x["Left"])
// 			right := dfs(x["Left"])

// 			var newBin BinaryOp
// 			json.Unmarshal(x["Op"], &newBin.Op)
// 			newBin.Left = left
// 			newBin.Right = right
// 			return &newBin
// 		}
// 	}
// 	return dfs(tree)
// }
