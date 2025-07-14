package client

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func RunTCPClient(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("TCP dial error: %v", err)
	}
	defer conn.Close()
	r := bufio.NewReader(conn)
	for i := 0; i < 5; i++ {
		start := time.Now()
		_, err := conn.Write([]byte("ping\n"))
		if err != nil {
			log.Fatalf("Write error: %v", err)
		}
		reply, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("Read error: %v", err)
		}
		latency := time.Since(start)
		log.Printf("Reply TCP: %sLatency: %v", strings.TrimSpace(reply), latency)
		time.Sleep(time.Second)
	}
}
