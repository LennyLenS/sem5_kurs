// clearCommitHistory
package main

import (
	"encoding/json"
	"fmt"
	"go/token"
	is "lib/infostructs"
	rq "lib/requests"
	tb "lib/table"
	"net/http"
	"os"
	"worker/winit"
	oper "worker/workpool"
)

var workerInfo *is.WorkerInfo

func wsolveproblem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Запрос на решение от кластера")
	rq.SendRequest(workerInfo.ClusterPort, "cbusyworker", workerInfo.Id)
	var op oper.Input
	er := json.NewDecoder(r.Body).Decode(&op)
	if er != nil {
		panic("ошибка парсера")
	}
	var newTable tb.Table
	fmt.Println(op.Root.Op)
	fmt.Println(op.Root.Left)
	fmt.Println(op.Root.Right)
	switch op.Root.Op {
	case token.ADD:
		newTable = oper.SUM(op, workerInfo)
	case token.MUL:
		newTable = oper.MUL(op, workerInfo)
	case token.QUO:
		newTable = oper.QUO(op, workerInfo)
	case token.SUB:
		newTable = oper.SUB(op, workerInfo)
	}
	ansjs, _ := json.Marshal(newTable)

	rq.SendRequest(workerInfo.ClusterPort, "cfreeworker", workerInfo.Id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ansjs)
}

func main() {
	args := os.Args
	if len(args) < 5 {
		panic("<порт><порт кластера><id><кол-во ядер>")
	}
	workerInfo = winit.WorkerInit(args)

	http.HandleFunc("/wsolveproblem", wsolveproblem)
	http.ListenAndServe(fmt.Sprintf(":%s", workerInfo.Port), nil)
}
