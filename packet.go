package socketer

import "io"

type Packet interface {
	// Pack
	Pack(writer io.Writer, src interface{}) error

	// UnPack
	UnPack(reader io.Reader, dst interface{}) error
}
