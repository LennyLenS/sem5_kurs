package tasks

import (
	rq "lib/requests"
	tb "lib/table"
	tr "lib/trees"
)

type ClusterWorkerTask struct {
	CWR             rq.ClusterWorkerReq
	ParentNode      tr.ASTNode
	ResultTableName string
	AllTables       *map[string]tb.Table
	AnalChannel     chan bool
}

func (task *ClusterWorkerTask) CheckReady() bool {
	_, ok1 := task.CWR.Root.(*tr.BinaryOp).Left.(*tr.TableLeaf)
	_, ok2 := task.CWR.Root.(*tr.BinaryOp).Right.(*tr.TableLeaf)
	if ok1 && ok2 {
		return true
	}
	return false
}
