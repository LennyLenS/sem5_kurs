package client

import (
	"io"
	is "lib/infostructs"
	rq "lib/requests"
)

func Handler_msolveproblem(clientData io.ReadCloser, freeClusterReq chan chan *is.ManagerInfo) []byte {
	sendClusterChan := make(chan *is.ManagerInfo, 1)
	freeClusterReq <- sendClusterChan
	cluster := <-sendClusterChan

	var answerToClient []byte
	rq.SendRequest(cluster.Port, "csolveproblem", clientData, &answerToClient)
	return answerToClient
}
