package tests

import "surls/svc/surlssvc/handlers"

var surlsGrpcService *handlers.GrpcService

func init() {
	surlsGrpcService = handlers.NewSUrlsGrpcServer()
}
