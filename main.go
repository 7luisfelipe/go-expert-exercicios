package main

import (
	webhttp "modcleanarch/app/infrastructure/webHttp"
)

func main() {
	//Roda a REST
	go webhttp.Routes()

	//Roda a aplicação gRPC
	webhttp.WebGrpc()
}
