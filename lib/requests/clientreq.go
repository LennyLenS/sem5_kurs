package requests

import (
	mt "lib/table"
)

type ClientReq struct {
	Expr   string
	Tables map[string]mt.Table
}
