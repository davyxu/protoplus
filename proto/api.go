package proto

import (
	"github.com/davyxu/protoplus/wire"
)

func Marshal(msg wire.Struct) ([]byte, error) {

	l := msg.Size()

	data := make([]byte, 0, l)

	buffer := wire.NewBuffer(data)

	err := msg.Marshal(buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil

}

func Size(msg wire.Struct) int {

	return msg.Size()
}

func Unmarshal(data []byte, msg wire.Struct) (err error) {

	buffer := wire.NewBuffer(data)

	return wire.UnmarshalStructObject(buffer, msg)
}
