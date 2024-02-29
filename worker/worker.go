// clearCommitHistory
package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/token"
	is "lib/infostructs"
	rq "lib/requests"
	tb "lib/table"
	"net/http"
	"os"
	oper "worker/operation"
	"worker/winit"
)

var workerInfo *is.WorkerInfo

func wsolveproblem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Запрос на решение от менеджера")
	rq.SendRequest(workerInfo.ManagerPort, "cbusyworker", workerInfo.Id)
	var op oper.Input
	er := json.NewDecoder(r.Body).Decode(&op)
	if er != nil {
		panic("Ошибка парсера")
	}
	fset := token.NewFileSet()
	ast.Print(fset, op.Root)
	var newTable tb.Table
	if er == nil && len(op.Root.Fields) == 0 {
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
		fmt.Println(op.Root.Op)
	} else {
		fmt.Println(op.Root.Left.TableName)
		newTable = oper.Proj(op, workerInfo)
	}
	// fmt.Println(newTable)
	ansjs, _ := json.Marshal(newTable)

	rq.SendRequest(workerInfo.ManagerPort, "cfreeworker", workerInfo.Id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ansjs)
}

func main() {
	args := os.Args
	if len(args) < 5 {
		panic("<порт><порт менеджера><id><кол-во ядер>")
	}
	workerInfo = winit.WorkerInit(args)

	http.HandleFunc("/wsolveproblem", wsolveproblem)
	http.ListenAndServe(fmt.Sprintf(":%s", workerInfo.Port), nil)
}
