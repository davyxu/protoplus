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

	input.Struct = MySubType{
		Str: "world",
	}

	input.StructSlice = []MySubType{
		{Int32: 100},
		{Str: "200"},
	}

	return
}

func verify(t *testing.T, raw interface{}) {
	t.Logf("size: %d", proto.Size(raw))

	data, err := proto.Marshal(raw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("proto+:", len(data), data)

	newType := reflect.New(reflect.TypeOf(raw).Elem()).Interface()

	err = proto.Unmarshal(data, newType)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(raw, newType) {
		t.FailNow()
	}
}

func TestFull(t *testing.T) {

	input := makeMyType()

	verify(t, &input)

	t.Logf("%v", input.String())
}

func TestIntSlice(t *testing.T) {

	var input MyType
	input.Int32Slice = []int32{-1, 1, 2}

	verify(t, &input)
}

func testVariant(t *testing.T, raw interface{}) {
	data, err := proto.Marshal(raw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	tRaw := reflect.TypeOf(raw)

	var newValue interface{}
	if tRaw.Kind() == reflect.Slice {

		// go的反射不支持切片取地址, 因此这里需要手动展开
		switch raw.(type) {
		case []int32:
			newValue = &[]int32{}
		case []int64:
			newValue = &[]int64{}
		case []uint32:
			newValue = &[]uint32{}
		case []uint64:
			newValue = &[]uint64{}
		case []float32:
			newValue = &[]float32{}
		case []float64:
			newValue = &[]float64{}
		case []bool:
			newValue = &[]bool{}
		case []string:
			newValue = &[]string{}
		case []byte:
			newValue = &[]byte{}
		default:
			panic("unsupport type")
		}

	} else {
		newValue = reflect.New(tRaw).Interface()
	}

	err = proto.Unmarshal(data, newValue)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	newValue = reflect.ValueOf(newValue).Elem().Interface()

	if !reflect.DeepEqual(raw, newValue) {
		t.FailNow()
	}
}

func TestTopField(t *testing.T) {
	testVariant(t, int32(1))
	testVariant(t, int64(10))
	testVariant(t, uint32(100))
	testVariant(t, uint64(1000))
	testVariant(t, float32(10000.1234))
	testVariant(t, float64(100000.1234))
	testVariant(t, true)
	testVariant(t, "12345")
	testVariant(t, []int32{1, 2, 3, 4})
	testVariant(t, []int64{1, 2, 3, 4})
	testVariant(t, []uint32{1, 2, 3, 4})
	testVariant(t, []uint64{1, 2, 3, 4})
	testVariant(t, []float32{1, 2, 3, 4})
	testVariant(t, []float64{1, 2, 3, 4})
	testVariant(t, []bool{true, false, true})
	testVariant(t, []string{"100", "2", "3"})
}
