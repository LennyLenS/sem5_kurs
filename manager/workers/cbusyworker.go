package workers

import (
	"encoding/json"
	"io"
	is "lib/info"
)

func Handler_cbusyworker(reqData io.ReadCloser, managerInfo *is.ManagerInfo, workersPool chan *is.WorkerInfo) {
	var busyWorkerId int
	err := json.NewDecoder(reqData).Decode(&busyWorkerId)
	if err != nil {
		panic("Ошибка парса воркера на кластере, cbusyworker")
	}
}
