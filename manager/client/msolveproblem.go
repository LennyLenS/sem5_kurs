package client

import (
	"io"
	is "lib/infostructs"
	rq "lib/requests"
)

func Handler_msolveproblem(clientData io.ReadCloser, freeManagerReq chan chan *is.ManagerInfo) []byte {
	sendManagerChan := make(chan *is.ManagerInfo, 1)
	freeManagerReq <- sendManagerChan
	cluster := <-sendManagerChan

	var answerToClient []byte
	rq.SendRequest(cluster.Port, "csolveproblem", clientData, &answerToClient)
	return answerToClient
}
