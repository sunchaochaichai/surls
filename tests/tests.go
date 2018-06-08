package tests

import (
	"surls/svc/surlssvc/transports"
)

var surlsGrpcService *transports.GrpcService

func init() {
	surlsGrpcService = transports.NewSUrlsGrpcServer()
}
