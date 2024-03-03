package info

type ManagerInfo struct {
	Port        string
	WorkersList *map[int]*WorkerInfo
}

type WorkerInfo struct {
	Port        string
	ManagerPort string
	Id          int
	Cores       int
}
