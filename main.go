package main

import "rentacar/webserver"

func main() {
	//create a new GIN endpoint server (tudo que for de servidor)
	server := webserver.New()
	server.Run
}
