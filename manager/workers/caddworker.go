package workers

import (
	"encoding/json"
	"io"
	is "lib/infostructs"
)

func Handler_caddworker(reqData io.ReadCloser, clusterInfo *is.ManagerInfo, workersPool chan *is.WorkerInfo) {
	var newWorker is.WorkerInfo
	err := json.NewDecoder(reqData).Decode(&newWorker)
	if err != nil {
		panic("Ошибка парса воркера на кластере caddworker")
	}

	(*clusterInfo.WorkersList)[newWorker.Id] = &newWorker
	workersPool <- &newWorker
}
