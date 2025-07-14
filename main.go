package main

import (
	"flag"
	"fmt"
	"os"
	"ping-pong/client"
	"ping-pong/server"
	"strings"
)

var (
	mode = flag.String("mode", "tcp", "protocol to use: tcp or udp")
	srvr = flag.Bool("server", false, "run in server mode")
	addr = flag.String("addr", ":9000", "address to bind or connect to (host:port)")
)

func main() {
	flag.Parse()

	switch strings.ToLower(*mode) {
	case "tcp":
		if *srvr {
			server.RunTCPServer(*addr)
		} else {
			client.RunTCPClient(*addr)
		}
	case "udp":
		if *srvr {
			server.RunUDPServer(*addr)
		} else {
			client.RunUDPClient(*addr)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown mode: %s\n", *mode)
		os.Exit(1)
	}
}
