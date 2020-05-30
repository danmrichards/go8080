package go8080

// ana is the "Logical AND Register With Accumulator" handler.
//
// The specified byte is logically ANDed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ana(v uint8) {
	r := i.r[A] & v
	i.cc.cy = false
	i.cc.ac = ((i.r[A] | v) & 0x08) != 0
	i.cc.z = r == 0
	i.cc.s = r&0x80 != 0
	i.cc.setParity(r)
	i.r[A] = r
}

// xra is the "Logical Exclusive-Or Register With Accumulator" handler.
//
// The specified byte is EXCLUSIVE-ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) xra(v uint8) {
	i.r[A] ^= v
	i.cc.cy = false
	i.cc.ac = false
	i.cc.z = i.r[A] == 0
	i.cc.s = i.r[A]&0x80 != 0
	i.cc.setParity(i.r[A])
}

// ora is the "Logical OR Register With Accumulator" handler.
//
// The specified byte is logically ORed bit by bit with the contents of the
// accumulator. The Carry bit is reset to zero.
func (i *Intel8080) ora(v uint8) {
	i.r[A] |= v
	i.cc.cy = false
	i.cc.ac = false
	i.cc.z = i.r[A] == 0
	i.cc.s = i.r[A]&0x80 != 0
	i.cc.setParity(i.r[A])
}

// cmp is the "Compare Register With Accumulator" handler.
//
// The specified byte is compared to the contents of the accumulator. The
// comparison is performed by internally subtracting the contents of REG from
// the accumulator (leaving both unchanged) and setting the condition bits
// according to the result.
//
// In particular, the Zero bit is set if the quantities are equal, and reset if
// they are unequal. Since a subtract operation is performed, the Carry bit will
// be set if there is no carry out of bit 7, indicating that the contents of REG
// are greater than the contents of the accumulator, and reset otherwise.
func (i *Intel8080) cmp(v uint8) {
	r := int16(i.r[A]) - int16(v)
	i.cc.cy = r&0x100 != 0
	i.cc.ac = ^(i.r[A]^uint8(r)^v)&0x10 != 0
	i.cc.z = r&0xff == 0
	i.cc.s = r&0x80 != 0
	i.cc.setParity(byte(r))
}

// rlc is the "Rotate Accumulator Left" handler.
//
// The Carry bit is set equal to the high-order bit of the accumulator. The
// contents of the accumulator are rotated one bit position to the left, with
// the high-order bit being transferred to the low-order bit position of the
// accumulator.
func (i *Intel8080) rlc() {
	i.cc.cy = i.r[A]&0x80 != 0
	i.r[A] <<= 1
	if i.cc.cy {
		i.r[A] |= 0x01
	}
}

// rrc is the "Rotate Accumulator Right" handler.
//
// The carry bit is set equal to the low-order bit of the accumulator. The
// contents of the accumulator are rotated one bit position to the right, with
// the low-order bit being transferred to the high-order bit position of the
// accumulator.
func (i *Intel8080) rrc() {
	i.cc.cy = i.r[A]&0x01 != 0
	i.r[A] >>= 1
	if i.cc.cy {
		i.r[A] |= 0x80
	}
}

// ral is the "Rotate Accumulator Left Through Carry" handler.
//
// The contents of the accumulator are rotated one bit position to the left.
//
// The high-order bit of the accumulator replaces the carry bit, while the carry
// bit replaces the high-order bit of the accumulator.
func (i *Intel8080) ral() {
	cy := i.cc.cy
	i.cc.cy = i.r[A]&0x80 != 0
	i.r[A] <<= 1
	if cy {
		i.r[A] |= 0x01
	}
}

// rar is the "Rotate Accumulator Right Through Carry" handler.
//
// The contents of the accumulator are rotated one bit position to the right.
//
// The low-order bit of the accumulator replaces the carry bit, while the carry
// bit replaces the high-order bit of the accumulator.
func (i *Intel8080) rar() {
	cy := i.cc.cy
	i.cc.cy = i.r[A]&0x01 != 0
	i.r[A] >>= 1
	if cy {
		i.r[A] |= 0x80
	}
}

// cma is the "Compliment Accumulator" handler.
//
// Each bit of the contents of the accumulator is complemented (producing the
// one's complement).
//
// E.g. 01010001 -> 10101110
func (i *Intel8080) cma() {
	i.r[A] ^= 0xff
}
