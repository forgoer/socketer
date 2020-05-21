package packet

import (
	"bufio"
	"errors"
	"io"
	"reflect"
)

const delim = '\n'

var EOF = &EOFpacket{Delim: delim}

type EOFpacket struct {
	Delim byte
}

func (p *EOFpacket) Pack(writer io.Writer, src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("Pack failed: src must be of type byte.")
	}
	_, err := writer.Write(b)
	return err
}

func (p *EOFpacket) UnPack(reader io.Reader, dst interface{}) error {
	b, err := bufio.NewReader(reader).ReadBytes(p.Delim)

	v := reflect.ValueOf(dst).Elem()

	v.SetBytes(b)

	return err
}
