package socketer

import (
	"bufio"
	"io"
	"net"
)

type Server struct {
	listener net.Listener

	address string
	network string

	closeListener   CloseListener
	receiveListener ReceiveListener
	connectListener ConnectListener
}

func NewServer(network, address string) *Server {
	return &Server{
		network: network,
		address: address,
	}
}

func (s *Server) OnConnect(l ConnectListener) {
	s.connectListener = l
}

func (s *Server) OnReceive(l ReceiveListener) {
	s.receiveListener = l
}

func (s *Server) OnClose(l CloseListener) {
	s.closeListener = l
}

func (s *Server) Send(conn net.Conn, data []byte) (int, error) {
	return conn.Write(data)
}

func (s *Server) Close(conn net.Conn) error {
	err := conn.Close()
	if s.closeListener != nil {
		s.connectListener(conn)
	}
	return err
}

func (s *Server) Listen() error {
	var err error
	s.listener, err = net.Listen(s.network, s.address)
	if err != nil {
		return err
	}

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		if s.connectListener != nil {
			s.connectListener(conn)
		}
		go s.accept(conn)
	}
}

func (s *Server) accept(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				s.Close(conn)
			}
			break
		}

		if s.receiveListener != nil {
			go s.receiveListener(message, conn)
		}
	}
}
