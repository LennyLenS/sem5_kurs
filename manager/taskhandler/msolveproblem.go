package taskhandler

import (
	"encoding/json"
	"io"
	gr "lib/generatelib"
	rq "lib/requests"
	ts "lib/tasks"
	tr "lib/trees"
)

func Handler_msolveproblem(workerTaskPool chan ts.WorkerTask, reqData io.ReadCloser) []byte {
	var inputData rq.ClientReq
	err := json.NewDecoder(reqData).Decode(&inputData)
	if err != nil {
		panic("Ошибка парса задачи на менеджере!")
	}

	tables := inputData.Tables
	tree := tr.ParseExpr(inputData.Expr)
	lastName := ""
	var lastAnalChannel chan bool
	var dfs func(node tr.ASTNode, prevnode tr.ASTNode)
	i := 0
	dfs = func(node, prevnode tr.ASTNode) {
		if x, ok := node.(*tr.BinaryOp); ok && len(node.(*tr.BinaryOp).Fields) == 0 {
			i++
			dfs(x.Left, node)
			dfs(x.Right, node)
			newTask := ts.WorkerTask{
				CWR:             rq.WorkerReq{Root: x, Tables: tables},
				ParentNode:      prevnode,
				ResultTableName: gr.GetRandString(100),
				AllTables:       &tables,
				AnalChannel:     make(chan bool, 1),
			}
			workerTaskPool <- newTask
			lastName = newTask.ResultTableName
			lastAnalChannel = newTask.AnalChannel
		} else if x, ok := node.(*tr.BinaryOp); ok {
			i++
			dfs(x.Left, node)
			newTask := ts.WorkerTask{
				CWR:             rq.WorkerReq{Root: x, Tables: tables},
				ParentNode:      prevnode,
				ResultTableName: gr.GetRandString(100),
				AllTables:       &tables,
				AnalChannel:     make(chan bool, 1),
			}
			workerTaskPool <- newTask
			lastName = newTask.ResultTableName
			lastAnalChannel = newTask.AnalChannel
		}
	}
	dfs(tree, tree)

	<-lastAnalChannel
	result, err := json.Marshal(tables[lastName])
	if err != nil {
		panic("Ошибка парса результата поддерева на менеджере")
	}
	return result
}
