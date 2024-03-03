package minit

import (
	is "lib/info"
	ts "lib/tasks"
	"manager/taskhandler"
)

func ManagerInit(args []string) (*is.ManagerInfo, chan *is.WorkerInfo, chan ts.WorkerTask) {
	managerPort := args[1]
	workersPool := make(chan *is.WorkerInfo, 50)
	workerTaskPool := make(chan ts.WorkerTask, 1000000)
	managerInfo := &is.ManagerInfo{
		Port:        managerPort,
		WorkersList: &map[int]*is.WorkerInfo{},
	}
	go taskhandler.TasksPool(workerTaskPool, workersPool)

	return managerInfo, workersPool, workerTaskPool
}
