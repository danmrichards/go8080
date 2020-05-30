package go8080

// ei is the "enable interrupt" handler.
func (i *Intel8080) ei() {
	i.ie = true
}

// di is the "disable interrupt" handler.
func (i *Intel8080) di() {
	i.ie = false
}

// hlt is the "Halt" handler.
func (i *Intel8080) hlt() {
	i.pc--
	i.halted = true
}

// stc is the "Set Carry" handler.
func (i *Intel8080) stc() {
	i.cc.cy = true
}

// cmc is the "Complement Carry" handler.
//
// If the Carry bit = 0, it is set to 1. If the Carry bit = 1, it is reset to O.
func (i *Intel8080) cmc() {
	i.cc.cy = !i.cc.cy
}
