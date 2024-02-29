// clearCommitHistory
package main

import (
	"fmt"
	gn "lib/generatelib"
	rq "lib/requests"
	mt "lib/table"
)

const (
	m int = 10000
	n int = 4
)

func getTask() rq.ClientReq {
	expr := "(a / b) * c"
	a := gn.GenerateRandTable(2000, 50)
	b := gn.GenerateRandTable(1000, 50)
	c := gn.GenerateRandTable(20, 50)
	d := gn.GenerateRandTable(m, 50)
	e := gn.GenerateRandTable(m, 50)
	f := gn.GenerateRandTable(n, 50)

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
