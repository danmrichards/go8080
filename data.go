package go8080

// movRR is the "Move Register to Register" handler.
//
// One byte of data is moved from the register specified by src (the source
// register) to the register specified by dst (the destination register).
func (i *Intel8080) movRR(opc byte) {
	d := (opc >> 3) & 0x7
	s := opc & 0x7
	i.r[d] = i.r[s]
}

// movRR is the "Move Memory to Register" handler.
//
// One byte of data is moved from the memory address pointed by the HL register
// pair, to the given register r.
func (i *Intel8080) movMR(opc byte) {
	d := (opc >> 3) & 0x7
	a := i.hl()
	i.r[d] = i.mem.Read(a)
}

// movRR is the "Move Register to Memory" handler.
//
// One byte of data is moved from the register specified by r (the source
// register) to the memory address pointed by the HL register pair.
func (i *Intel8080) movRM(opc byte) {
	s := opc & 0x7
	a := i.hl()
	i.mem.Write(a, i.r[s])
}

// mvi is the "Move Immediate Data" handler.
//
// The byte of immediate data is stored in the specified register.
func (i *Intel8080) mvi(opc byte) {
	d := (opc >> 3) & 0x7
	v := i.immediateByte()
	i.r[d] = v
}

// mvi is the "Move Immediate Data Memory" handler.
//
// The byte of immediate data is stored in the register specified by the byte
// pointed by the HL register pair.
func (i *Intel8080) mviM() {
	// Determine the address of the byte pointed by the HL register pair.
	// The address is two bytes long, so merge the two bytes stored in each
	// side of the register pair.
	addr := i.hl()

	i.mem.Write(addr, i.immediateByte())
}

// ldax is the "Load Accumulator" handler.
//
// The contents of the memory location addressed by registers B and C, or by
// registers D and E, replace the contents of the accumulator.
func (i *Intel8080) ldax(addr uint16) {
	i.r[A] = i.mem.Read(addr)
}

// stax is the "Store Accumulator" handler.
//
// The contents of the accumulator are stored in the memory location addressed
// by registers B an dC, or by registers 0 and E.
func (i *Intel8080) stax(addr uint16) {
	i.mem.Write(addr, i.r[A])
}

// shld is the "Store H and L Direct" handler.
//
// The contents of the L register are stored at the memory address formed by
// concatenating HI ADD with LOW ADD. The contents of the H register are stored
// at the next higher memory address.
func (i *Intel8080) shld() {
	addr := i.immediateWord()

	hl := i.hl()

	i.mem.Write(addr, byte(hl&0xff))
	i.mem.Write(addr+1, byte(hl>>8))
}

// lhld is the "Load H and L Direct" handler.
//
// The byte at the memory address formed by concatenating HI ADD with LOW ADD
// replaces the contents of the L register. The byte at the next higher memory
// address replaces the contents of the H register.
func (i *Intel8080) lhld() {
	addr := i.immediateWord()

	b := uint16(i.mem.Read(addr)) | uint16(i.mem.Read(addr+1))<<8

	i.setHL(b)
}

// xchg is the "Exchange Registers" handler.
//
// The 16 bits of data held in the H and L registers are exchanged with the 16
// bits of data held in the D and E registers.
func (i *Intel8080) xchg() {
	de, hl := i.de(), i.hl()
	i.setDE(hl)
	i.setHL(de)
}
