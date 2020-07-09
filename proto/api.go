package proto

import (
	"github.com/davyxu/protoplus/wire"
)

type Struct = wire.Struct

func Marshal(msg Struct) ([]byte, error) {

	l := msg.Size()

	data := make([]byte, 0, l)

	buffer := wire.NewBuffer(data)

	err := msg.Marshal(buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil

}

func Size(msg Struct) int {

	return msg.Size()
}

func Unmarshal(data []byte, msg Struct) (err error) {

	buffer := wire.NewBuffer(data)

	return wire.UnmarshalStructObject(buffer, msg)
}
