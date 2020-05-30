package go8080

import "fmt"

// handleOp dispatches the appropriate handler for the given opcode.
func (i *Intel8080) handleOp(opc byte) error {
	switch opc {
	case 0x00, 0x10, 0x20, 0x30, 0x08, 0x18, 0x28, 0x38:
		// NOP and ignore opcodes.

	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c,
		0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a,
		0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68,
		0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d,
		0x7f:
		i.movRR(opc)

	case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x7e:
		i.movMR(opc)

	case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
		i.movRM(opc)

	case 0x3e, 0x06, 0x0e, 0x16, 0x1e, 0x26, 0x2e:
		i.mvi(opc)

	case 0x36:
		i.mviM()

	case 0x0a:
		i.ldax(i.bc())

	case 0x1a:
		i.ldax(i.de())

	case 0x3a:
		i.ldax(i.immediateWord())

	case 0x02:
		i.stax(i.bc())

	case 0x12:
		i.stax(i.de())

	case 0x32:
		i.stax(i.immediateWord())

	case 0x01:
		// LXI B, W
		i.setBC(i.immediateWord())

	case 0x11:
		// LXI D, W
		i.setDE(i.immediateWord())

	case 0x21:
		// LXI H, W
		i.setHL(i.immediateWord())

	case 0x31:
		// LXI SP, W
		i.sp = i.immediateWord()

	case 0x2a:
		i.lhld()

	case 0x22:
		i.shld()

	case 0xf9:
		i.sphl()

	case 0xeb:
		i.xchg()

	case 0xe3:
		i.xthl()

	case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87:
		i.add(i.opcRegVal(opc))

	case 0x86:
		i.addM()

	case 0xc6:
		i.adi()

	case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f:
		i.adc(i.opcRegVal(opc))

	case 0x8e:
		i.adcM()

	case 0xce:
		i.aci()

	case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97:
		i.sub(i.opcRegVal(opc))

	case 0x96:
		i.subM()

	case 0xd6:
		i.sui()

	case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f:
		i.sbb(i.opcRegVal(opc))

	case 0x9e:
		i.sbbM()

	case 0xde:
		i.sbi()

	case 0x09:
		i.dad(i.bc())

	case 0x19:
		i.dad(i.de())

	case 0x29:
		i.dad(i.hl())

	case 0x39:
		i.dadSP()

	case 0xf3:
		i.di()

	case 0xfb:
		i.ei()

	case 0x76:
		i.hlt()

	case 0x4, 0xc, 0x14, 0x1c, 0x24, 0x2c, 0x3c:
		d := (opc >> 3) & 0x7
		i.r[d] = i.inr(i.r[d])

	case 0x34:
		i.inrM()

	case 0x05, 0x0d, 0x15, 0x1d, 0x25, 0x2d, 0x3d:
		d := (opc >> 3) & 0x7
		i.r[d] = i.dcr(i.r[d])

	case 0x35:
		i.dcrM()

	case 0x03:
		// INX B
		i.setBC(i.bc() + 1)

	case 0x13:
		// INX D
		i.setDE(i.de() + 1)

	case 0x23:
		// INX H
		i.setHL(i.hl() + 1)

	case 0x33:
		// INX SP
		i.inxSP()

	case 0x0b:
		// DCX B
		i.setBC(i.bc() - 1)

	case 0x1b:
		// DCX D
		i.setDE(i.de() - 1)

	case 0x2b:
		// DCX H
		i.setHL(i.hl() - 1)

	case 0x3b:
		// DCX SP
		i.dcxSP()

	case 0x27:
		i.daa()

	case 0x2f:
		i.cma()

	case 0x37:
		i.stc()

	case 0x3f:
		i.cmc()

	case 0x07:
		i.rlc()

	case 0x0f:
		i.rrc()

	case 0x17:
		i.ral()

	case 0x1f:
		i.rar()

	case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7:
		// ANA r
		d := opc & 0x7
		i.ana(i.r[d])

	case 0xa6:
		// ANA M
		i.ana(i.mem.Read(i.hl()))

	case 0xe6:
		// ANI
		i.ana(i.immediateByte())

	case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf:
		// XRA r
		d := opc & 0x7
		i.xra(i.r[d])

	case 0xae:
		// XRA M
		i.xra(i.mem.Read(i.hl()))

	case 0xee:
		i.xra(i.immediateByte())

	case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7:
		// ORA r
		d := opc & 0x7
		i.ora(i.r[d])

	case 0xb6:
		// ORA M
		i.ora(i.mem.Read(i.hl()))

	case 0xf6:
		// ORI
		i.ora(i.immediateByte())

	case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
		// CMP r
		d := opc & 0x7
		i.cmp(i.r[d])

	case 0xbe:
		// CMP M
		i.cmp(i.mem.Read(i.hl()))

	case 0xfe:
		// CPI
		i.cmp(i.immediateByte())

	case 0xc3:
		i.jmp()

	case 0xc2:
		i.jnz()

	case 0xca:
		i.jz()

	case 0xd2:
		i.jnc()

	case 0xda:
		i.jc()

	case 0xe2:
		i.jpo()

	case 0xea:
		i.jpe()

	case 0xf2:
		i.jp()

	case 0xfa:
		i.jm()

	case 0xe9:
		i.pchl()

	case 0xcd:
		i.call()

	case 0xc4:
		i.cnz()

	case 0xcc:
		i.cz()

	case 0xd4:
		i.cnc()

	case 0xdc:
		i.cic()

	case 0xe4:
		i.cpo()

	case 0xec:
		i.cpe()

	case 0xf4:
		i.cp()

	case 0xfc:
		i.cm()

	case 0xc9, 0xd9:
		i.ret()

	case 0xc0:
		i.rnz()

	case 0xc8:
		i.rz()

	case 0xd0:
		i.rnc()

	case 0xd8:
		i.rc()

	case 0xe0:
		i.rpo()

	case 0xe8:
		i.rpe()

	case 0xf0:
		i.rp()
	case 0xf8:
		i.rm()

	case 0xc7, 0xcf, 0xd7, 0xdf, 0xe7, 0xef, 0xf7, 0xff:
		i.rst(opc)

	case 0xc5:
		// PUSH B
		i.stackAdd(i.bc())

	case 0xd5:
		// PUSH D
		i.stackAdd(i.de())

	case 0xe5:
		// PUSH H
		i.stackAdd(i.hl())

	case 0xf5:
		// PUSH PSW
		i.pushPSW()

	case 0xc1:
		// POP B
		i.setBC(i.stackPop())

	case 0xd1:
		// POP D
		i.setDE(i.stackPop())

	case 0xe1:
		// POP H
		i.setHL(i.stackPop())

	case 0xf1:
		// POP PSW
		i.popPSW()

	case 0xdb:
		i.in()

	case 0xd3:
		i.out()

	default:
		return fmt.Errorf(
			"unsupported opcode 0x%02x at program counter %04x", opc, i.pc,
		)
	}

	return nil
}
