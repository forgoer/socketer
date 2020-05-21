package socketer

import (
	"bytes"
	"github.com/forgoer/socketer/packet"
	"net"
)

type Client struct {
	Conn net.Conn

	address string
	network string

	closeListener   CloseListener
	receiveListener ReceiveListener
	connectListener ConnectListener

	packet Packet
}

func NewClient(network, address string) *Client {
	client := &Client{
		address: address,
		network: network,
		packet:  packet.EOF,
	}
	return client
}

func (c *Client) Dial() error {
	return c.dial()
}

func (c *Client) dial() error {
	if c.Conn != nil {
		return nil
	}

	var err error
	c.Conn, err = net.Dial(c.network, c.address)
	if err == nil && c.connectListener != nil {
		c.connectListener(c.Conn)
	}

	return err
}

func (c *Client) SetPacket(packet Packet) {
	c.packet = packet
}

func (c *Client) Send(data interface{}) (int, error) {
	var err error
	buffer := bytes.NewBufferString("")
	err = c.packet.Pack(buffer, data)
	if err != nil {
		return 0, err
	}

	return c.Conn.Write(buffer.Bytes())
}

func (c *Client) OnConnect(l ConnectListener) {
	c.connectListener = l
}

func (c *Client) OnReceive(l ReceiveListener) {
	c.receiveListener = l
}

func (c *Client) OnClose(l CloseListener) {
	c.closeListener = l
}

func (c *Client) Close() error {
	err := c.Conn.Close()
	if c.closeListener != nil {
		c.closeListener(c.Conn)
	}
	return err
}

func (c *Client) Recv(dst interface{}) error {
	err := c.packet.UnPack(c.Conn, dst)
	return err
}

func (c *Client) Open() error {
	var err error
	defer c.Conn.Close()

	for {
		var buffer []byte
		err = c.packet.UnPack(c.Conn, &buffer)
		if err != nil {
			break
		}

		if c.receiveListener != nil {
			go c.receiveListener(buffer, c.Conn)
		}
	}

	return err
}
