package proto

type WireType = uint64

const (
	WireVarint   WireType = iota + 1 // int32, int64, uint32, uint64, bool, enum
	WireBytes                        // Length-delimited   string, bytes, slice, struct
	WireZigzag32                     // 32位负数
	WireZigzag64                     // 64位负数
	WireFixed32                      // 32位定长 float32
	WireFixed64                      // 64位定长 float64
)

func makeWireTag(tag uint64, wt WireType) WireType {
	return uint64(tag)<<3 | uint64(wt)
}

func parseWireTag(wireTag WireType) (tag uint64, wt WireType) {
	tag = wireTag >> 3
	wt = WireType(wireTag & 7)
	return
}
