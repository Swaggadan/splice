package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: proxyserver listenaddr targetaddr")
	}

	ln, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	client, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	upstream, err := net.Dial("tcp", os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(upstream, client)
	client.Close()
}
