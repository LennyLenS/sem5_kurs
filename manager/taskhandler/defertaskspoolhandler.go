package taskhandler

import (
	"fmt"
	"go/ast"
	"go/token"
	is "lib/infostructs"
	rq "lib/requests"
	tb "lib/table"
	ts "lib/tasks"
	tr "lib/trees"
)

func DeferTasksPoolHandler(deferWorkerTaskPool chan ts.WorkerTask, workersPool chan *is.WorkerInfo) {
	for {
		deferTask := <-deferWorkerTaskPool
		if !deferTask.CheckReady() {
			deferWorkerTaskPool <- deferTask
			continue
		}

		worker := <-workersPool
		fset := token.NewFileSet()
		i := 0
		go func() {
			fmt.Println(i, "----------------------------------")
			i++
			fmt.Println("Before")
			ast.Print(fset, deferTask.CWR.Root)
			var resultTable tb.Table
			rq.SendRequest(worker.Port, "wsolveproblem", deferTask.CWR, &resultTable)
			fmt.Println(i, "----------------------------------")
			fmt.Println("Before")
			ast.Print(fset, deferTask.CWR.Root)
			(*deferTask.AllTables)[deferTask.ResultTableName] = resultTable
			node := &tr.TableLeaf{
				TableName: deferTask.ResultTableName,
				Info:      resultTable.Info,
			}
			if deferTask.ParentNode.(*tr.BinaryOp) != deferTask.CWR.Root.(*tr.BinaryOp) {
				fmt.Println("defertask proj   ", deferTask.ParentNode.(*tr.BinaryOp).Left)
				_, ok := deferTask.ParentNode.(*tr.BinaryOp).Left.(*tr.BinaryOp)
				if ok && deferTask.ParentNode.(*tr.BinaryOp).Left.(*tr.BinaryOp) == deferTask.CWR.Root.(*tr.BinaryOp) {
					deferTask.ParentNode.(*tr.BinaryOp).Left = node
				} else {
					deferTask.ParentNode.(*tr.BinaryOp).Right = node
				}
				fmt.Println("After", "defertask after proj   ", deferTask.ParentNode.(*tr.BinaryOp).Left)
				ast.Print(fset, deferTask.ParentNode)
			}
			deferTask.AnalChannel <- true
		}()
	}
}
