package taskhandler

import (
	"encoding/json"
	"fmt"
	"io"
	gr "lib/generatelib"
	is "lib/infostructs"

	// tb "lib/table"
	rq "lib/requests"
	ts "lib/tasks"
	tr "lib/trees"
)

func Handler_csolveproblem(workersPool chan *is.WorkerInfo, deferClusterWorkerTaskPool chan ts.ClusterWorkerTask, reqData io.ReadCloser) []byte {
	var inputData rq.ClientReq
	err := json.NewDecoder(reqData).Decode(&inputData)
	if err != nil {
		panic("Ошибка парса задачи на кластере! csolveproblem")
	}

	tables := inputData.Tables
	tree := tr.ParseExpr(inputData.Expr)
	// fmt.Println(tree.(*tr.BinaryOp))

	lastName := ""
	var lastAnalChannel chan bool
	var dfs func(node tr.ASTNode, prevnode tr.ASTNode)
	dfs = func(node, prevnode tr.ASTNode) {
		if x, ok := node.(*tr.BinaryOp); ok {
			dfs(x.Left, node)
			dfs(x.Right, node)
			fmt.Println(x.Left)
			newTask := ts.ClusterWorkerTask{
				CWR:             rq.ClusterWorkerReq{Root: x, Tables: tables},
				ParentNode:      prevnode,
				ResultTableName: gr.GetRandString(100),
				AllTables:       &tables,
				AnalChannel:     make(chan bool, 1),
			}
			deferClusterWorkerTaskPool <- newTask
			lastName = newTask.ResultTableName
			lastAnalChannel = newTask.AnalChannel
		}
	}
	dfs(tree, tree)

	<-lastAnalChannel
	result, err := json.Marshal(tables[lastName])
	if err != nil {
		panic("Ошибка парса результата поддерева на кластере")
	}
	return result
}
