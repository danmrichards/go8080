package go8080

// jmp is the "Jump" handler.
//
// This handler jumps the program counter to a given point in memory.
func (i *Intel8080) jmp() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	i.pc = i.immediateWord()
}

// call is the "Call subroutine" handler.
//
// A call operation is unconditionally performed to subroutine sub.
func (i *Intel8080) call() {
	// We will dump the program to the subroutine indicated by the two immediate
	// bytes in memory.
	addr := i.immediateWord()

	// Update the stack pointer.
	i.stackAdd(i.pc)

	i.pc = addr
}

// ret is the "Return" handler.
//
// A return operation is unconditionally performed.
func (i *Intel8080) ret() {
	i.pc = i.stackPop()
}

// jnz is the "Jump If Not Zero" handler.
//
// If the zero bit is one, program execution continues at the memory address adr.
func (i *Intel8080) jnz() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.z {
		i.pc = addr
	}
}

// jz is the "Jump Zero" handler.
//
// If the zero bit is not one, program execution continues at the memory address
// adr.
func (i *Intel8080) jz() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.z {
		i.pc = addr
	}
}

// jnc is the "Jump Not Carry" handler.
//
// If the carry bit is one, program execution continues at the memory address
// adr.
func (i *Intel8080) jnc() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.cy {
		i.pc = addr
	}
}

// jc is the "Jump Carry" handler.
//
// If the carry bit is not one, program execution continues at the memory
// address adr.
func (i *Intel8080) jc() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.cy {
		i.pc = addr
	}
}

// jpo is the "Jump If Parity Odd" handler.
//
// If the Parity bit is zero (indicating a result with odd parity), program
// execution continues at the memory address adr.
func (i *Intel8080) jpo() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.p {
		i.pc = addr
	}
}

// jpe is the "Jump If Parity Even" handler.
//
// If the Parity bit is one (indicating a result with even parity), program
// execution continues at the memory address adr.
func (i *Intel8080) jpe() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.p {
		i.pc = addr
	}
}

// jp is the "Jump If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), program execution
// continues at the memory address adr.
func (i *Intel8080) jp() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.s {
		i.pc = addr
	}
}

// jm is the "Jump If Minus" handler.
//
// If the Sign bit is one (indicating a positive result), program execution
// continues at the memory address adr.
func (i *Intel8080) jm() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.s {
		i.pc = addr
	}
}

// cz is the "Call If Zero" handler.
//
// If the Zero bit is zero, a call operation is performed to subroutine sub.
func (i *Intel8080) cz() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.z {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// cnz is the "Call If Not Zero" handler.
//
// If the Zero bit is one, a call operation is performed to subroutine sub.
func (i *Intel8080) cnz() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.z {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// cc is the "Call If Carry" handler.
//
// If the Carry bit is zero, a call operation is performed to subroutine sub.
func (i *Intel8080) cic() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.cy {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// cnc is the "Call If Not carry" handler.
//
// If the carry bit is one, a call operation is performed to subroutine sub.
func (i *Intel8080) cnc() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.cy {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// cpo is the "Call If Parity Odd" handler.
//
// If the Parity bit is one (indicating a result with even parity), a call
// operation is performed to subroutine sub.
func (i *Intel8080) cpo() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.p {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// cpe is the "Call If Parity Even" handler.
//
// If the Parity bit is even (indicating a result with even parity), a call
// operation is performed to subroutine sub.
func (i *Intel8080) cpe() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.p {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// cp is the "Call If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), a call operation is
// performed to subroutine sub.
func (i *Intel8080) cp() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if !i.cc.s {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// cp is the "Call If Minus" handler.
//
// If the Sign bit is one (indicating a positive result), a call operation is
// performed to subroutine sub.
func (i *Intel8080) cm() {
	// The address to jump to is two bytes long, so get the next two bytes from
	// memory (most significant first) and merge them.
	addr := i.immediateWord()

	if i.cc.s {
		i.stackAdd(i.pc)
		i.pc = addr
		i.cyc += 6
	}
}

// rnz is the "Return If Not Zero" handler.
//
// If the Zero bit is zero, a return operation is performed.
func (i *Intel8080) rnz() {
	if !i.cc.z {
		i.ret()
		i.cyc += 6
	}
}

// rz is the "Return If Zero" handler.
//
// If the Zero bit is one, a return operation is performed.
func (i *Intel8080) rz() {
	if i.cc.z {
		i.ret()
		i.cyc += 6
	}
}

// rnc is the "Return If Not Carry" handler.
//
// If the Carry bit is zero, a return operation is performed.
func (i *Intel8080) rnc() {
	if !i.cc.cy {
		i.ret()
		i.cyc += 6
	}
}

// rc is the "Return If Carry" handler.
//
// If the Carry bit is one, a return operation is performed.
func (i *Intel8080) rc() {
	if i.cc.cy {
		i.ret()
		i.cyc += 6
	}
}

// rpo is the "Return If Parity Odd" handler.
//
// If the Parity bit is zero (indicating a result with odd parity), a return
// operation is performed.
func (i *Intel8080) rpo() {
	if !i.cc.p {
		i.ret()
		i.cyc += 6
	}
}

// rpe is the "Return If Parity Even" handler.
//
// If the Parity bit is one (indicating a result with event parity), a return
// operation is performed.
func (i *Intel8080) rpe() {
	if i.cc.p {
		i.ret()
		i.cyc += 6
	}
}

// rp is the "Return If Positive" handler.
//
// If the Sign bit is zero (indicating a positive result), a return operation
// is performed.
func (i *Intel8080) rp() {
	if !i.cc.s {
		i.ret()
		i.cyc += 6
	}
}

// rm is the "Return If Minus" handler.
//
// If the Sign bit is one (indicating a negative result), a return operation
// is performed.
func (i *Intel8080) rm() {
	if i.cc.s {
		i.ret()
		i.cyc += 6
	}
}

// pchl is the "Load Program Counter" handler.
//
// The contents of the H register replaces the most significant 8 bits of the
// program counter, and the contents of the L register replace the least
// significant 8 bits of the program counter.
//
// This causes program execution to continue at the address contained in the H
// and L registers.
func (i *Intel8080) pchl() {
	i.pc = i.hl()
}

// rst is the "Restart" handler.
//
// The contents of the program counter are pushed onto the stack, providing a
// return address for later use by a RETURN instruction.
//
// The program execution continues at an address indicated by opc.
func (i *Intel8080) rst(opc byte) {
	i.stackAdd(i.pc)

	i.pc = uint16(opc) & 0x38
}
