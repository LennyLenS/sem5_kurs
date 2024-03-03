package workers

import (
	"encoding/json"
	"io"
	is "lib/info"
)

func Handler_caddworker(reqData io.ReadCloser, managerInfo *is.ManagerInfo, workersPool chan *is.WorkerInfo) {
	var newWorker is.WorkerInfo
	err := json.NewDecoder(reqData).Decode(&newWorker)
	if err != nil {
		panic("Ошибка парса воркера на кластере caddworker")
	}

	(*managerInfo.WorkersList)[newWorker.Id] = &newWorker
	workersPool <- &newWorker
}
