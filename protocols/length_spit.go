package protocols

import "encoding/binary"

type LengthSpitProtocol struct {
	ByteOrder binary.ByteOrder
	Offset uint32
	BodyLength uint32
}
