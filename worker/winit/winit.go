package winit

import (
	is "lib/infostructs"
	rq "lib/requests"
	"strconv"
)

func WorkerInit(args []string) *is.WorkerInfo {
	workerPort := args[1]
	managerPort := args[2]
	id, err := strconv.Atoi(args[3])
	if err != nil {
		panic("Ошибка парса Id")
	}
	cores, err := strconv.Atoi(args[4])
	if err != nil {
		panic("Ошибка парса кол-во ядер")
	}

	workerInfo := is.WorkerInfo{
		Port:        workerPort,
		ManagerPort: managerPort,
		Id:          id,
		Cores:       cores,
	}

	rq.SendRequest(workerInfo.ManagerPort, "caddworker", workerInfo)
	return &workerInfo
}
