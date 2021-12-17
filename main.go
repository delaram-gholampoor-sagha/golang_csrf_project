package main

import "log"

var host = "localhost"
var port = "9000"

func main() {
	db.InitDB()

	iwtErr := myJwt.InitJWT()
	if myJwt != nil {
		log.Panicln("Error initializing the jwt")
		log.Fatal(jwtErr)
	}

	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Print("error starting the server")
		log.Fatal(serverErr)
	}

}
