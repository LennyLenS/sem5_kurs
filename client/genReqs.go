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

func getFormedTask() rq.ClientReq {
	expr := "(b + a) * a"
	a := gn.GenerateRandTable(2, 50)
	b := gn.GenerateRandTable(10, 50)
	fmt.Println(a)
	fmt.Println(b)
	c := gn.GenerateRandTable(n, 50)
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
	Req := getFormedTask()
	var Ans mt.Table
	rq.SendRequest(port, "msolveproblem", Req, &Ans)
	fmt.Println(Ans)
}
