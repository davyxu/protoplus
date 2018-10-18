package tests

import (
	"github.com/davyxu/protoplus/proto"
	"math"
	"reflect"
	"testing"
)

func TestOptional(t *testing.T) {
	bigData := makeMyType()
	data, err := proto.Marshal(&bigData)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	var output MyTypeMini
	err = proto.Unmarshal(data, &output)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%+v", output)

}

func makeMyType() (input MyType) {

	input.Bool = true
	input.Int32 = 200
	input.UInt32 = math.MaxUint32 - 100
	input.Int64 = -789
	input.UInt64 = 1234567890123456
	input.Str = "hello"
	input.Float32 = 3.14
	input.Float64 = math.MaxFloat64
	input.BoolSlice = []bool{true, false, true}

	input.Struct = &MyType{
		Str: "world",
	}

	return
}

func TestFull(t *testing.T) {

	input := makeMyType()

	t.Logf("size: %d", proto.Size(&input))

	data, err := proto.Marshal(&input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("proto+:", len(data), data)

	var output MyType
	err = proto.Unmarshal(data, &output)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(input, output) {
		t.FailNow()
	}

	t.Logf("%+v", output)
}
