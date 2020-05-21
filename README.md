Socketer
====

## Installation

```
go get -u github.com/forgoer/socketer
```

## Usage

Server:
```go
    serv := socketer.NewServer("tcp", "127.0.0.1:6000")

    serv.OnReceive(func(data []byte, conn net.Conn) {
        fmt.Println("Server Received: ", string(data))
        serv.Send(conn, []byte("ok\n"))
    })
	
    serv.Listen()
```


Client:

```go
    client := socketer.NewClient("tcp", "127.0.0.1:6000")

    client.OnReceive(func(data []byte, conn net.Conn) {
        fmt.Println("Client Received: " + string(data))
    })

    client.Dial()
```