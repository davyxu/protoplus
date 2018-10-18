package proto

func Marshal(raw interface{}) ([]byte, error) {

	msg := raw.(Struct)

	l := msg.Size()

	data := make([]byte, 0, l)

	buffer := NewBuffer(data)

	err := msg.Marshal(buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func Size(raw interface{}) int {
	msg := raw.(Struct)

	return msg.Size()
}

func Unmarshal(data []byte, raw interface{}) (err error) {

	msg := raw.(Struct)

	buffer := NewBuffer(data)

	return rawUnmarshalStruct(buffer, msg)

}
