package tests

import (
	"encoding/json"
	"github.com/davyxu/protoplus/proto"
	"github.com/davyxu/protoplus/wire"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

func TestOptional(t *testing.T) {
	bigData := makeMyType()
	data, err := proto.Marshal(&bigData)
	assert.Equal(t, err, nil)
	var output MyTypeMini
	assert.Equal(t, proto.Unmarshal(data, &output), nil)

	t.Logf("%+v", output)
	assert.Equal(t, bigData, output)

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

	input.Struct = MySubType{
		Str: "world",
	}

	input.StructSlice = []MySubType{
		{Int32: 100},
		{Str: "200"},
	}

	return
}

func verifyWire(t *testing.T, raw wire.Struct) {
	data, err := proto.Marshal(raw)
	assert.Equal(t, err, nil)

	t.Log("proto+:", len(data), data)

	newType := reflect.New(reflect.TypeOf(raw).Elem()).Interface().(wire.Struct)

	assert.Equal(t, proto.Unmarshal(data, newType), nil)

	assert.Equal(t, raw, newType)
}

func verifyText(t *testing.T, raw interface{}) {

	tRaw := reflect.TypeOf(raw)

	if tRaw.Kind() != reflect.Ptr {
		panic("expect ptr")
	}

	data := proto.CompactTextString(raw)

	t.Log(data)

	newType := reflect.New(tRaw.Elem()).Interface()

	assert.Equal(t, proto.UnmarshalText(data, newType), nil)

	assert.Equal(t, raw, newType)
}

func TestFull(t *testing.T) {

	input := makeMyType()

	verifyWire(t, &input)

	t.Logf("%v", proto.MarshalTextString(input))
}

func TestIntSlice(t *testing.T) {

	var input MyType
	input.Int32Slice = []int32{-1, 1, 2}

	verifyWire(t, &input)
}

func TestSkipField(t *testing.T) {

	input := makeMyType()

	data, err := proto.Marshal(&input)
	assert.Equal(t, err, nil)

	jsondata, _ := json.Marshal(&input)

	var mini MyTypeMini
	assert.Equal(t, proto.Unmarshal(data, &mini), nil)

	var miniJson MyTypeMini
	json.Unmarshal(jsondata, &miniJson)
	assert.Equal(t, miniJson, mini)
}

func TestPtrField(t *testing.T) {

	input := MyType{}
	data, err := proto.Marshal(&input)
	t.Log(data, err)

}

func TestText(t *testing.T) {

	input := makeMyType()

	verifyText(t, &input)
}

func TestFloat(t *testing.T) {

	type MyFloat struct {
		Value float64
	}

	input := MyFloat{math.MaxFloat64}

	verifyText(t, &input)
}

func TestSlice(t *testing.T) {

	type DummyStruct struct {
		Num int32
	}

	type MyFloat struct {
		Value []int32
		Dummy DummyStruct
	}

	input := MyFloat{
		[]int32{-1, 1, 2},
		DummyStruct{5},
	}

	verifyText(t, &input)
}

func TestEnum(t *testing.T) {

	type MyFloat struct {
		Value []MyEnum
	}

	input := MyFloat{Value: []MyEnum{MyEnum_One}}

	verifyText(t, &input)
}
