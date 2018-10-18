package proto

type MyType struct {
	Bool   bool
	Int32  int32
	UInt32 uint32

	Int64  int64
	UInt64 uint64

	Float32 float32
	Float64 float64
	String  string
	Struct  *MyType
}

func (self *MyType) Size() (ret int) {

	ret += SizeBool(0, self.Bool)

	ret += SizeInt32(1, self.Int32)

	ret += SizeUInt32(2, self.UInt32)

	ret += SizeInt64(3, self.Int64)

	ret += SizeUInt64(4, self.UInt64)

	ret += SizeFloat32(5, self.Float32)

	ret += SizeFloat64(6, self.Float64)

	ret += SizeString(7, self.String)

	ret += SizeStruct(8, self.Struct)

	return
}

func (self *MyType) Marshal(buffer *Buffer) error {

	if err := MarshalBool(buffer, 0, self.Bool); err != nil {
		return err
	}

	if err := MarshalInt32(buffer, 1, self.Int32); err != nil {
		return err
	}

	if err := MarshalUInt32(buffer, 2, self.UInt32); err != nil {
		return err
	}

	if err := MarshalInt64(buffer, 3, self.Int64); err != nil {
		return err
	}

	if err := MarshalUInt64(buffer, 4, self.UInt64); err != nil {
		return err
	}

	if err := MarshalFloat32(buffer, 5, self.Float32); err != nil {
		return err
	}

	if err := MarshalFloat64(buffer, 6, self.Float64); err != nil {
		return err
	}

	if err := MarshalString(buffer, 7, self.String); err != nil {
		return err
	}

	if err := MarshalStruct(buffer, 8, self.Struct); err != nil {
		return err
	}

	return nil
}

func (self *MyType) Unmarshal(buffer *Buffer, fieldIndex uint64, wt WireType) error {
	switch fieldIndex {
	case 0:
		return UnmarshalBool(buffer, wt, &self.Bool)
	case 1:
		return UnmarshalInt32(buffer, wt, &self.Int32)
	case 2:
		return UnmarshalUInt32(buffer, wt, &self.UInt32)
	case 3:
		return UnmarshalInt64(buffer, wt, &self.Int64)
	case 4:
		return UnmarshalUInt64(buffer, wt, &self.UInt64)
	case 5:
		return UnmarshalFloat32(buffer, wt, &self.Float32)
	case 6:
		return UnmarshalFloat64(buffer, wt, &self.Float64)
	case 7:
		return UnmarshalString(buffer, wt, &self.String)
	case 8:
		return UnmarshalStruct(buffer, wt, &self.Struct)
	}

	return nil
}
