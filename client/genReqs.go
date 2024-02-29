// clearCommitHistory
package main

import (
	"fmt"
	gn "lib/generatelib"
	rq "lib/requests"
	mt "lib/table"
)

const (
	m int = 1000
	n int = 4
)

func getTask() rq.ClientReq {
	expr := "(((a + b) / c) * ((d - e) / f))[id, name]"
	a := gn.GenerateRandTable(100, 50)
	b := gn.GenerateRandTable(1000, 50)
	c := gn.GenerateRandTable(20, 50)
	d := gn.GenerateRandTable2(m, 25)
	e := gn.GenerateRandTable2(m, 25)
	f := gn.GenerateRandTable2(m, 25)

	h := map[string]mt.Table{
		"a": a,
		"b": b,
		"c": c,
		"d": d,
		"e": e,
		"f": f,
	}

	req := rq.ClientReq{
		Expr:   expr,
		Tables: h,
	}
	return req
}

func SendRequest(port string) {
	Req := getTask()
	var Ans mt.Table
	rq.SendRequest(port, "msolveproblem", Req, &Ans)
	fmt.Println(Ans)
}
