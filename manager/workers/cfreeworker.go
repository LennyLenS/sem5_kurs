package workers

import (
	"encoding/json"
	"io"
	is "lib/info"
)

func Handler_cfreeworker(reqData io.ReadCloser, managerInfo *is.ManagerInfo, workersPool chan *is.WorkerInfo) {
	var freeWorkerId int
	err := json.NewDecoder(reqData).Decode(&freeWorkerId)
	if err != nil {
		panic("Ошибка парса воркера на кластере, cfreeworker")
	}

	freeWorker := (*managerInfo.WorkersList)[freeWorkerId]
	workersPool <- freeWorker
}
