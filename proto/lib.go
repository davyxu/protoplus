package proto

// A Buffer is a buffer manager for marshaling and unmarshaling
// protocol buffers.  It may be reused between invocations to
// reduce memory usage.  It is not necessary to use a Buffer;
// the global functions Marshal and Unmarshal create a
// temporary Buffer and are fine for most applications.
type Buffer struct {
	buf   []byte // encode/decode byte stream
	index int    // read point
}

// NewBuffer allocates a new Buffer and initializes its internal data to
// the contents of the argument slice.
func NewBuffer(e []byte) *Buffer {
	return &Buffer{buf: e}
}

// Reset resets the Buffer, ready for marshaling a new protocol buffer.
func (p *Buffer) Reset() {
	p.buf = p.buf[0:0] // for reading/writing
	p.index = 0        // for reading
}

// SetBuf replaces the internal buffer with the slice,
// ready for unmarshaling the contents of the slice.
func (p *Buffer) SetBuf(s []byte) {
	p.buf = s
	p.index = 0
}

// Bytes returns the contents of the Buffer.
func (p *Buffer) Bytes() []byte { return p.buf }

func (p *Buffer) ConsumeBytes(size int) (ret []byte) {
	ret = p.buf[p.index : p.index+size]

	p.index += size

	return
}

// EncodeVarint writes a varint-encoded integer to the Buffer.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
// protocol buffer types.
func (p *Buffer) EncodeVarint(x uint64) error {
	for x >= 1<<7 {
		p.buf = append(p.buf, uint8(x&0x7f|0x80))
		x >>= 7
	}
	p.buf = append(p.buf, uint8(x))
	return nil
}

// EncodeFixed64 writes a 64-bit integer to the Buffer.
// This is the format for the
// fixed64, sfixed64, and double protocol buffer types.
func (p *Buffer) EncodeFixed64(x uint64) error {
	p.buf = append(p.buf,
		uint8(x),
		uint8(x>>8),
		uint8(x>>16),
		uint8(x>>24),
		uint8(x>>32),
		uint8(x>>40),
		uint8(x>>48),
		uint8(x>>56))
	return nil
}

// EncodeFixed32 writes a 32-bit integer to the Buffer.
// This is the format for the
// fixed32, sfixed32, and float protocol buffer types.
func (p *Buffer) EncodeFixed32(x uint64) error {
	p.buf = append(p.buf,
		uint8(x),
		uint8(x>>8),
		uint8(x>>16),
		uint8(x>>24))
	return nil
}

// EncodeZigzag64 writes a zigzag-encoded 64-bit integer
// to the Buffer.
// This is the format used for the sint64 protocol buffer type.
func (p *Buffer) EncodeZigzag64(x uint64) error {
	// use signed number to get arithmetic right shift.
	return p.EncodeVarint(uint64((x << 1) ^ uint64(int64(x)>>63)))
}

// EncodeZigzag32 writes a zigzag-encoded 32-bit integer
// to the Buffer.
// This is the format used for the sint32 protocol buffer type.
func (p *Buffer) EncodeZigzag32(x uint64) error {
	// use signed number to get arithmetic right shift.
	return p.EncodeVarint(uint64((uint32(x) << 1) ^ uint32(int32(x)>>31)))
}

// EncodeRawBytes writes a count-delimited byte buffer to the Buffer.
// This is the format used for the bytes protocol buffer
// type and for embedded messages.
func (p *Buffer) EncodeRawBytes(b []byte) error {
	p.EncodeVarint(uint64(len(b)))
	p.buf = append(p.buf, b...)
	return nil
}

// EncodeStringBytes writes an encoded string to the Buffer.
// This is the format used for the proto2 string type.
func (p *Buffer) EncodeStringBytes(s string) error {
	p.EncodeVarint(uint64(len(s)))
	p.buf = append(p.buf, s...)
	return nil
}

// errOverflow is returned when an integer is too large to be represented.
var errOverflow = errors.New("proto: integer overflow")

func (p *Buffer) decodeVarintSlow() (x uint64, err error) {
	i := p.index
	l := len(p.buf)

	for shift := uint(0); shift < 64; shift += 7 {
		if i >= l {
			err = io.ErrUnexpectedEOF
			return
		}
		b := p.buf[i]
		i++
		x |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			p.index = i
			return
		}
	}

	// The number is too large to represent in a 64-bit value.
	err = errOverflow
	return
}

// DecodeVarint reads a varint-encoded integer from the Buffer.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
// protocol buffer types.
func (p *Buffer) DecodeVarint() (x uint64, err error) {
	i := p.index
	buf := p.buf

	if i >= len(buf) {
		return 0, io.ErrUnexpectedEOF
	} else if buf[i] < 0x80 {
		p.index++
		return uint64(buf[i]), nil
	} else if len(buf)-i < 10 {
		return p.decodeVarintSlow()
	}

	var b uint64
	// we already checked the first byte
	x = uint64(buf[i]) - 0x80
	i++

	b = uint64(buf[i])
	i++
	x += b << 7
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 7

	b = uint64(buf[i])
	i++
	x += b << 14
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 14

	b = uint64(buf[i])
	i++
	x += b << 21
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 21

	b = uint64(buf[i])
	i++
	x += b << 28
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 28

	b = uint64(buf[i])
	i++
	x += b << 35
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 35

	b = uint64(buf[i])
	i++
	x += b << 42
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 42

	b = uint64(buf[i])
	i++
	x += b << 49
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 49

	b = uint64(buf[i])
	i++
	x += b << 56
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 56

	b = uint64(buf[i])
	i++
	x += b << 63
	if b&0x80 == 0 {
		goto done
	}

	return 0, errOverflow

done:
	p.index = i
	return x, nil
}

// DecodeFixed64 reads a 64-bit integer from the Buffer.
// This is the format for the
// fixed64, sfixed64, and double protocol buffer types.
func (p *Buffer) DecodeFixed64() (x uint64, err error) {
	// x, err already 0
	i := p.index + 8
	if i < 0 || i > len(p.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	p.index = i

	x = uint64(p.buf[i-8])
	x |= uint64(p.buf[i-7]) << 8
	x |= uint64(p.buf[i-6]) << 16
	x |= uint64(p.buf[i-5]) << 24
	x |= uint64(p.buf[i-4]) << 32
	x |= uint64(p.buf[i-3]) << 40
	x |= uint64(p.buf[i-2]) << 48
	x |= uint64(p.buf[i-1]) << 56
	return
}

// DecodeFixed32 reads a 32-bit integer from the Buffer.
// This is the format for the
// fixed32, sfixed32, and float protocol buffer types.
func (p *Buffer) DecodeFixed32() (x uint64, err error) {
	// x, err already 0
	i := p.index + 4
	if i < 0 || i > len(p.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	p.index = i

	x = uint64(p.buf[i-4])
	x |= uint64(p.buf[i-3]) << 8
	x |= uint64(p.buf[i-2]) << 16
	x |= uint64(p.buf[i-1]) << 24
	return
}

// DecodeZigzag64 reads a zigzag-encoded 64-bit integer
// from the Buffer.
// This is the format used for the sint64 protocol buffer type.
func (p *Buffer) DecodeZigzag64() (x uint64, err error) {
	x, err = p.DecodeVarint()
	if err != nil {
		return
	}
	x = (x >> 1) ^ uint64((int64(x&1)<<63)>>63)
	return
}

// DecodeZigzag32 reads a zigzag-encoded 32-bit integer
// from  the Buffer.
// This is the format used for the sint32 protocol buffer type.
func (p *Buffer) DecodeZigzag32() (x uint64, err error) {
	x, err = p.DecodeVarint()
	if err != nil {
		return
	}
	x = uint64((uint32(x) >> 1) ^ uint32((int32(x&1)<<31)>>31))
	return
}

// DecodeRawBytes reads a count-delimited byte buffer from the Buffer.
// This is the format used for the bytes protocol buffer
// type and for embedded messages.
func (p *Buffer) DecodeRawBytes(alloc bool) (buf []byte, err error) {
	n, err := p.DecodeVarint()
	if err != nil {
		return nil, err
	}

	nb := int(n)
	if nb < 0 {
		return nil, fmt.Errorf("proto: bad byte length %d", nb)
	}
	end := p.index + nb
	if end < p.index || end > len(p.buf) {
		return nil, io.ErrUnexpectedEOF
	}

	if !alloc {
		// todo: check if can get more uses of alloc=false
		buf = p.buf[p.index:end]
		p.index += nb
		return
	}

	buf = make([]byte, nb)
	copy(buf, p.buf[p.index:])
	p.index += nb
	return
}

// DecodeStringBytes reads an encoded string from the Buffer.
// This is the format used for the proto2 string type.
func (p *Buffer) DecodeStringBytes() (s string, err error) {
	buf, err := p.DecodeRawBytes(false)
	if err != nil {
		return
	}
	return string(buf), nil
}
