package cluster

import (
	"encoding/json"
	"io"
	gn "lib/generatelib"
	rq "lib/requests"
	tb "lib/table"
)

func Handler_wsolveproblem(reqData io.ReadCloser) tb.Table {
	var task rq.ClusterWorkerReq
	err := json.NewDecoder(reqData).Decode(&task)
	if err != nil {
		panic("Ошибка распарса задачи на воркере от клиента!")
	}

	newMatrix := gn.GenerateRandTable(3, 25)
	return newMatrix
}
