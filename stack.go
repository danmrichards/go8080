package go8080

// popPSW is the "Pop Data Off Stack PSW" handler.
//
// The contents of the PSW register pair are restored from two bytes of
// memory indicated by the stack pointer SP.
func (i *Intel8080) popPSW() {
	n := i.stackPop()

	i.r[A] = uint8(n >> 8)
	i.cc.setStatus(uint8(n & 0xff))
}

// pushPSW is the "Push Data Onto Stack PSW" handler.
//
// The contents of the PSW register pair are saved in two bytes of memory
// indicated by the stack pointer SP.
func (i *Intel8080) pushPSW() {
	i.stackAdd(uint16(i.r[A])<<8 | uint16(i.cc.status()))
}

// xthl is the "Exchange Stack" handler.
//
// The contents of the L register are exchanged with the contents of the memory
// byte whose address is held in the stack pointer SP. The contents of the H
// register are exchanged with the contents of the memory byte whose address is
// one greater than that held in the stack pointer.
func (i *Intel8080) xthl() {
	b := uint16(i.mem.Read(i.sp)) | uint16(i.mem.Read(i.sp+1))<<8
	hl := i.hl()

	i.setHL(b)

	i.mem.Write(i.sp, uint8(hl))
	i.mem.Write(i.sp+1, uint8(hl>>8))
}

// sphl is the "Load SP from H and L" handler.
//
// The 16 bits of data held in the H and L registers replace the contents of the
// stack pointer SP. The contents of the H and L registers are unchanged.
func (i *Intel8080) sphl() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := i.hl()

	i.sp = addr
}
