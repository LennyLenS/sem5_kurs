package requests

import (
	tb "lib/table"
	tr "lib/trees"
)

type ClusterWorkerReq struct {
	Root   tr.ASTNode
	Tables map[string]tb.Table
}
