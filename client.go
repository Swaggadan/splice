package main

import (
	"bytes"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: proxyclient proxyaddr")
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data := bytes.Repeat([]byte{'x'}, 1024*1024)

	go func() {
		for {
			_, err := conn.Write(data)
			if err != nil {
				return
			}
		}
	}()
	<-time.After(1 * time.Minute)
	conn.Close()
}
