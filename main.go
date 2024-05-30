package main

import (
	webhttp "modcleanarch/app/infrastructure/webHttp"
)

func main() {
	//Roda a aplicação REST
	go webhttp.WebRest()

	//Roda a aplicação gRPC
	go webhttp.WebGrpc()

	//Roda a aplicação GraphQl
	webhttp.WebGraphQL()
}
