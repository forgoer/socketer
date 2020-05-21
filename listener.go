package socketer

import "net"

type ConnectListener func(conn net.Conn)

type ReceiveListener func(data []byte, conn net.Conn)

type CloseListener func(conn net.Conn)
