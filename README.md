# ping-pong

Simple ping-pong application

## Usage:
```shell
go run main.go -mode tcp -server -addr :9000 # run tcp server
go run main.go -mode tcp -addr localhost:9000 # run tcp client
go run main.go -mode udp -server -addr :9001 # run udp server
go run main.go -mode udp -addr localhost:9001 # run udp client
```
