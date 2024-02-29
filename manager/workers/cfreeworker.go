package workers

import (
	"encoding/json"
	"io"
	is "lib/infostructs"
)

func Handler_cfreeworker(reqData io.ReadCloser, clusterInfo *is.ManagerInfo, workersPool chan *is.WorkerInfo) {
	var freeWorkerId int
	err := json.NewDecoder(reqData).Decode(&freeWorkerId)
	if err != nil {
		panic("Ошибка парса воркера на кластере, cfreeworker")
	}

	freeWorker := (*clusterInfo.WorkersList)[freeWorkerId]
	workersPool <- freeWorker
}
