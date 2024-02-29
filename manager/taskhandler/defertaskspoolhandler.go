package taskhandler

import (
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
		go func() {
			var resultTable tb.Table
			rq.SendRequest(worker.Port, "wsolveproblem", deferTask.CWR, &resultTable)
			(*deferTask.AllTables)[deferTask.ResultTableName] = resultTable
			node := &tr.TableLeaf{
				TableName: deferTask.ResultTableName,
				Info:      resultTable.Info,
			}
			if deferTask.ParentNode.(*tr.BinaryOp) != deferTask.CWR.Root.(*tr.BinaryOp) {
				_, ok := deferTask.ParentNode.(*tr.BinaryOp).Left.(*tr.BinaryOp)
				if ok && deferTask.ParentNode.(*tr.BinaryOp).Left.(*tr.BinaryOp) == deferTask.CWR.Root.(*tr.BinaryOp) {
					deferTask.ParentNode.(*tr.BinaryOp).Left = node
				} else {
					deferTask.ParentNode.(*tr.BinaryOp).Right = node
				}
			}
			deferTask.AnalChannel <- true
		}()
	}
}
