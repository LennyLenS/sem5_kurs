package infostructs

type ManagerInfo struct {
	Port        string
	WorkersList *map[int]*WorkerInfo
}
