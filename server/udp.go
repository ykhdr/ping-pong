package server

import (
	"log"
	"net"
	"strings"
)

func RunUDPServer(address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("Resolve UDP addr error: %v", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("UDP listen error: %v", err)
	}
	defer conn.Close()
	log.Printf("UDP srvr listening on %s", address)

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("ReadFromUDP error: %v", err)
			continue
		}
		msg := strings.TrimSpace(string(buf[:n]))
		log.Printf("Received UDP from %s: %s", clientAddr, msg)
		var resp string
		if msg == "ping" {
			resp = "pong"
		} else {
			resp = "unknown"
		}
		conn.WriteToUDP([]byte(resp), clientAddr)
	}
}
