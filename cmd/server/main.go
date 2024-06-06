package main

import (
	"log"
	"mytcpapp/internal/handlers"
	"mytcpapp/pkg/utils"
	"net"
)

func main() {
	/* 	cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	   	listener, err := tls.Listen("tcp", "127.0.0.1:8080", config)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	defer listener.Close()
	   	utils.LogMessage("TLS server is listening on port 8080")

	   	for {
	   		conn, err := listener.Accept()
	   		if err != nil {
	   			utils.LogMessage("Error accepting connection: " + err.Error())
	   			continue
	   		}
	   		go handlers.HandleConnection(conn)
	   	} */
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	utils.LogMessage("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.LogMessage("Error accepting connection: " + err.Error())
			continue
		}
		go handlers.HandleConnection(conn)
	}
}
