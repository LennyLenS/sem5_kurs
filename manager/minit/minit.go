package minit

import (
	is "lib/infostructs"
	ts "lib/tasks"
	"manager/taskhandler"
)

func ManagerInit(args []string) (*is.ManagerInfo, chan *is.WorkerInfo, chan ts.ClusterWorkerTask) {
	managerPort := args[1]
	workersPool := make(chan *is.WorkerInfo, 50)
	deferClusterWorkerTaskPool := make(chan ts.ClusterWorkerTask, 1000000)
	managerInfo := &is.ManagerInfo{
		Port:        managerPort,
		WorkersList: &map[int]*is.WorkerInfo{},
	}
	go taskhandler.DeferTasksPoolHandler(deferClusterWorkerTaskPool, workersPool)

	return managerInfo, workersPool, deferClusterWorkerTaskPool
}
