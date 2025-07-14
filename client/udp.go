package client

import (
	"log"
	"net"
	"strings"
	"time"
)

func RunUDPClient(address string) {
	serverAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("Resolve UDP addr error: %v", err)
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("UDP dial error: %v", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for i := 0; i < 5; i++ {
		start := time.Now()
		_, err := conn.Write([]byte("ping"))
		if err != nil {
			log.Fatalf("WriteToUDP error: %v", err)
		}
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalf("ReadFromUDP error: %v", err)
		}
		latency := time.Since(start)
		log.Printf("Reply UDP: %s Latency: %v", strings.TrimSpace(string(buf[:n])), latency)
		time.Sleep(time.Second)
	}
}
