package go8080

// add is the "Add Register to Accumulator" handler.
//
// The given byte is added to the contents of the accumulator and relevant
// condition bits are set.
func (i *Intel8080) add(n byte) {
	i.accumulatorAdd(n, 0)
}

// adi is the "Add Immediate to Accumulator" handler.
//
// The next byte of data from memory is added to the contents of the accumulator
// and relevant condition bits are set.
func (i *Intel8080) adi() {
	i.accumulatorAdd(i.immediateByte(), 0)
}

// addM is the "Add Memory to Accumulator" handler.
//
// The byte pointed to by the HL register pair is added to the contents of the
// accumulator and relevant condition bits are set.
func (i *Intel8080) addM() {
	i.accumulatorAdd(i.mem.Read(i.hl()), 0)
}

// inxSP is the "Increment Stack Pointer" handler.
func (i *Intel8080) inxSP() {
	i.sp++
}

// inr is the "Increment Register" handler.
//
// The specified register is incremented by one.
func (i *Intel8080) inr(n byte) byte {
	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(n) + 1

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = uint8(ans) == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = ans&0x80 != 0x00

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = ans&0xf == 0x00

	// Set the parity bit.
	i.cc.setParity(uint8(ans))

	return byte(ans)
}

// inrM is the "Increment Memory" handler.
//
// The byte pointed to by the HL register pair is incremented by one and
// relevant condition bits are set.
func (i *Intel8080) inrM() {
	addr := i.hl()

	i.mem.Write(addr, i.inr(i.mem.Read(addr)))
}

// dcr is the "Decrement Register" handler.
//
// The specified register is decremented by one.
func (i *Intel8080) dcr(n byte) byte {
	// Perform the arithmetic at higher precision in order to capture the
	// carry out.
	ans := uint16(n) - 1

	// Set the zero condition bit accordingly based on if the result of the
	// arithmetic was zero.
	i.cc.z = uint8(ans) == 0x00

	// Set the sign condition bit accordingly based on if the most
	// significant bit on the result of the arithmetic was set.
	//
	// Determine the result being zero with a bitwise AND operation against
	// 0x80 (10000000 in base 2 and 128 in base 10).
	//
	// 10000000 & 10000000 = 1
	i.cc.s = ans&0x80 != 0x00

	// Set the auxiliary carry condition bit accordingly if the result of
	// the arithmetic has a carry on the third bit.
	i.cc.ac = !(ans&0xf == 0xf)

	// Set the parity bit.
	i.cc.setParity(uint8(ans))

	return uint8(ans)
}

// dcrM is the "Decrement Memory" handler.
//
// The specified register is decremented by one.
func (i *Intel8080) dcrM() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := i.hl()

	i.mem.Write(addr, i.dcr(i.mem.Read(addr)))
}

// dad is the "Double Add" handler.
//
// The 16-bit number in the specified register pair is added to the 16-bit
// number held in the H and L registers using two's complement arithmetic. The
// result replaces the contents of the H and L registers.
func (i *Intel8080) dad(n uint16) {
	ans := uint32(i.hl()) + uint32(n)

	// Set the carry condition bit accordingly.
	i.cc.cy = ans&0x10000 != 0

	i.setHL(uint16(ans))
}

// dad is the "Double Add Stack Pointer" handler.
//
// The 16-bit number in the stack pointer is added to the 16-bit number held in
// the H and L registers using two's complement arithmetic. The result replaces
// the contents of the H and L registers.
func (i *Intel8080) dadSP() {
	ans := uint32(i.hl()) + uint32(i.sp)

	// Set the carry condition bit accordingly.
	i.cc.cy = ans&0x10000 != 0

	i.setHL(uint16(ans))
}

// dcxSP is the "Decrement Stack Pointer" handler.
func (i *Intel8080) dcxSP() {
	i.sp--
}

// daa is the "Decimal Adjust Accumulator" handler.
//
// The eight-bit hexadecimal number in the accumulator is adjusted to form two
// four-bit binary coded decimal digits.
func (i *Intel8080) daa() {
	var (
		a uint8
		c = i.cc.cy
	)

	lsb := i.r[A] & 0x0f
	msb := i.r[A] >> 4

	// If the least significant four bits of the accumulator represents a number
	// greater than 9, or if the Auxiliary Carry bit is equal to one, the
	// accumulator is incremented by six. Otherwise, no incrementing occurs.
	if lsb > 9 || i.cc.ac {
		a += 0x06
	}

	// If the most significant four bits of the accumulator now represent a
	// number greater than 9, or if the normal carry bit is equal to one, the
	// most significant four bits of the accumulator are incremented by six.
	if msb > 9 || i.cc.cy || (msb >= 9 && lsb > 9) {
		a += 0x60
		c = true
	}

	i.accumulatorAdd(a, 0)
	i.cc.setParity(i.r[A])
	i.cc.cy = c
}

// adc is the "Add Register to Accumulator With Carry" handler.
//
// The specified byte plus the content of the Carry bit is added to the contents
// of the accumulator.
func (i *Intel8080) adc(n byte) {
	i.accumulatorAdd(n, i.cc.carryByte())
}

// adcM is the "Add Memory to Accumulator With Carry" handler.
//
// The specified byte plus the content of the Carry bit is added to the contents
// of the accumulator.
//
// The byte pointed to by the HL register pair, plus the content of the Carry
// bit, is added to the contents of the accumulator and relevant condition bits
// are set.
func (i *Intel8080) adcM() {
	i.accumulatorAdd(i.mem.Read(i.hl()), i.cc.carryByte())
}

// sub is the "Subtract Register from Accumulator" handler.
//
// The given byte is subtracted from the contents of the accumulator and
// relevant condition bits are set.
func (i *Intel8080) sub(n byte) {
	i.accumulatorSub(n, 0)
}

// subM is the "Subtract Memory from Accumulator" handler.
//
// The byte pointed to by the HL register pair is subtracted from the contents
// of the accumulator and relevant condition bits are set.
func (i *Intel8080) subM() {
	i.sub(i.mem.Read(i.hl()))
}

// sbb is the "Subtract Register from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the contents of the specified byte. This
// value is then subtracted from the accumulator using two's complement
// arithmetic.
func (i *Intel8080) sbb(n byte) {
	i.accumulatorSub(n, i.cc.carryByte())
}

// sbbM is the "Subtract Memory from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the contents of the byte pointed to by
// the HL register pair. This value is then subtracted from the accumulator
// using two's complement arithmetic.
func (i *Intel8080) sbbM() {
	i.accumulatorSub(i.mem.Read(i.hl()), i.cc.carryByte())
}

// aci is the "Add Immediate to Accumulator With Carry" handler.
//
// The next byte of data from memory, plus the contents of the Carry bit, is
// added to the contents of the accumulator and relevant condition bits are set.
func (i *Intel8080) aci() {
	i.accumulatorAdd(i.immediateByte(), i.cc.carryByte())
}

// sui is the "Subtract Immediate from Accumulator" handler.
//
// The next byte of data from memory is subtracted from the contents of the
// accumulator and relevant condition bits are set.
func (i *Intel8080) sui() {
	i.accumulatorSub(i.immediateByte(), 0)
}

// sbi is the "Subtract Immediate from Accumulator With Borrow" handler.
//
// The Carry bit is internally added to the byte of immediate data. This value
// is then subtracted from the accumulator using two'scomplement arithmetic.
func (i *Intel8080) sbi() {
	i.accumulatorSub(i.immediateByte(), i.cc.carryByte())
}
