package go8080

const (
	// Define the named working registers.
	B = iota
	C
	D
	E
	H
	L
	_ // This would be the 'F' register (conditions) but we're not using it.
	A
)

// opcRegVal returns the register value indicated by the given opcode.
func (i *Intel8080) opcRegVal(opc byte) byte {
	return i.r[opc&0x7]
}

// bc returns the data stored in the BC register pair.
//
// The data is two bytes long, so merge the two bytes stored in each side of the
// register pair.
func (i *Intel8080) bc() uint16 {
	return uint16(i.r[B])<<8 | uint16(i.r[C])
}

// de returns the data stored in the DE register pair.
//
// The data is two bytes long, so merge the two bytes stored in each side of the
// register pair.
func (i *Intel8080) de() uint16 {
	return uint16(i.r[D])<<8 | uint16(i.r[E])
}

// hl returns the data stored in the HL register pair.
//
// The data is two bytes long, so merge the two bytes stored in each side of the
// register pair.
func (i *Intel8080) hl() uint16 {
	return uint16(i.r[H])<<8 | uint16(i.r[L])
}

// setBC sets the contents of the BC register pair.
func (i *Intel8080) setBC(v uint16) {
	i.r[B] = byte(v >> 8)
	i.r[C] = byte(v)
}

// setDE sets the contents of the DE register pair.
func (i *Intel8080) setDE(v uint16) {
	i.r[D] = byte(v >> 8)
	i.r[E] = byte(v)
}

// setHL sets the contents of the HL register pair.
func (i *Intel8080) setHL(v uint16) {
	i.r[H] = byte(v >> 8)
	i.r[L] = byte(v)
}
