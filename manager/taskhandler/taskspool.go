package taskhandler

import (
	is "lib/info"
	rq "lib/requests"
	tb "lib/table"
	ts "lib/tasks"
	tr "lib/trees"
)

func TasksPool(WorkerTaskPool chan ts.WorkerTask, workersPool chan *is.WorkerInfo) {
	for {
		Task := <-WorkerTaskPool
		if !Task.CheckReady() {
			WorkerTaskPool <- Task
			continue
		}

		worker := <-workersPool
		go func() {
			var resultTable tb.Table
			rq.SendRequest(worker.Port, "wsolveproblem", Task.CWR, &resultTable)
			(*Task.AllTables)[Task.ResultTableName] = resultTable
			node := &tr.TableLeaf{
				TableName: Task.ResultTableName,
				Info:      resultTable.Info,
			}
			if Task.ParentNode.(*tr.BinaryOp) != Task.CWR.Root.(*tr.BinaryOp) {
				_, ok := Task.ParentNode.(*tr.BinaryOp).Left.(*tr.BinaryOp)
				if ok && Task.ParentNode.(*tr.BinaryOp).Left.(*tr.BinaryOp) == Task.CWR.Root.(*tr.BinaryOp) {
					Task.ParentNode.(*tr.BinaryOp).Left = node
				} else {
					Task.ParentNode.(*tr.BinaryOp).Right = node
				}
			}
			Task.AnalChannel <- true
		}()
	}
}
