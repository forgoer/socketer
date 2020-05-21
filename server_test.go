package socketer

import (
	"net"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	serv := NewServer("tcp", "127.0.0.1:6000")

	serv.OnReceive(func(data []byte, conn net.Conn) {
		t.Log("Server Received: ", string(data))
		serv.Send(conn, []byte("ok\n"))
	})
	go func() {
		err := serv.Listen()
		panic(err)
	}()

	time.Sleep(1 * time.Second)

	client := NewClient("tcp", "127.0.0.1:6000")

	defer client.Close()
	client.OnReceive(func(data []byte, conn net.Conn) {
		t.Log("Client Received: ", string(data))
	})

	client.Dial()

	client.Send([]byte("啊哈哈\n"))

	var data []byte
	err := client.Recv(&data)
	if err != nil {
		panic(err)
	}

	t.Log("Client Received: ", string(data))

	//client.Open()
}
