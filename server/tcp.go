package server

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func RunTCPServer(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("TCP listen error: %v", err)
	}
	log.Printf("TCP server listening on %s", address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go handleTCPConn(conn)
	}
}

func handleTCPConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}
		msg = strings.TrimSpace(msg)
		log.Printf("Received TCP: %s", msg)
		if msg == "ping" {
			conn.Write([]byte("pong\n"))
		} else {
			conn.Write([]byte("unknown\n"))
		}
	}
}
