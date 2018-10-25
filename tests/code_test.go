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
	input.Int32Slice = []int32{1, 2, 3, 4}
	input.UInt32Slice = []uint32{100, 200, 300, 400}
	input.Int64Slice = []int64{1, 2, 3, 4}
	input.UInt64Slice = []uint64{100, 200, 300, 400}
	input.StrSlice = []string{"genji", "dva", "bastion"}
	input.Float32Slice = []float32{1.1, 2.1, 3.2, 4.5}
	input.Float64Slice = []float64{1.1, 2.1, 3.2, 4.5}
	input.BytesSlice = []byte("bytes")
	input.Enum = MyEnum_Two
	input.EnumSlice = []MyEnum{MyEnum_Two, MyEnum_One, MyEnum_Zero}

	input.Struct = &MyType{
		Str: "world",
	}

	input.StructSlice = []*MyType{
		{Int32: 100},
		{Str: "200"},
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

	t.Logf("%v", output.String())
}
