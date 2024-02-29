// clearCommitHistory
package main

import (
	"fmt"
	is "lib/infostructs"
	ts "lib/tasks"
	mi "manager/minit"
	"manager/taskhandler"
	"manager/workers"
	"net/http"
	"os"
)

var ManagerPort string
var Workers map[int]*is.WorkerInfo

var managerInfo *is.ManagerInfo
var workersPool chan *is.WorkerInfo
var deferWorkerTaskPool chan ts.WorkerTask

func msolveproblem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Запрос на решение от клиента")
	result := taskhandler.Handler_csolveproblem(deferWorkerTaskPool, r.Body)
	fmt.Println("Воркер решил задачу, отправка ответа клиенту")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func maddworker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Добавление воркера в менеджеру")

	workers.Handler_caddworker(r.Body, managerInfo, workersPool)
}

func mfreeworker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Воркер выполнил свою работу и сообщил менеджеру")

	workers.Handler_cfreeworker(r.Body, managerInfo, workersPool)
}

func mbusyworker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Воркер взял свою работу и сообщил менеджеру")

	workers.Handler_cbusyworker(r.Body, managerInfo, workersPool)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("<порт>")
	}

	managerInfo, workersPool, deferWorkerTaskPool = mi.ManagerInit(args)

	http.HandleFunc("/msolveproblem", msolveproblem)

	http.HandleFunc("/caddworker", maddworker)
	http.HandleFunc("/cfreeworker", mfreeworker)
	http.HandleFunc("/cbusyworker", mbusyworker)

	http.ListenAndServe(fmt.Sprintf(":%s", managerInfo.Port), nil)
}
