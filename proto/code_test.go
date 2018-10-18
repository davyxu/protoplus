package proto

import (
	"math"
	"testing"
)

func TestCode(t *testing.T) {

	var md MyType
	md.Bool = true
	md.Int32 = 200
	md.UInt32 = math.MaxUint32 - 100
	md.Int64 = -789
	md.UInt64 = 1234567890123456
	md.String = "hello"
	md.Float32 = 3.14
	md.Float64 = math.MaxFloat64

	md.Struct = &MyType{
		String: "world",
	}

	t.Logf("size: %d", Size(&md))

	data, err := Marshal(&md)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("proto+:", len(data), data)
	//bindata, _ := goobjfmt.BinaryWrite(&md)
	//t.Log("binary:", len(bindata), bindata)

	var md2 MyType
	err = Unmarshal(data, &md2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%+v", md2)
}

//func TestPB(t *testing.T) {
//
//	var md MyData
//	md.Name = "hello"
//	md.Age = -1
//	data, err := proto.Marshal(&md)
//	if err != nil {
//		t.Error(err)
//		t.FailNow()
//	}
//
//	t.Log(data)
//
//	var md2 MyData
//	err = proto.Unmarshal(data, &md2)
//	if err != nil {
//		t.Error(err)
//		t.FailNow()
//	}
//}
