package go8080

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type mem []byte

// Read returns the value from memory at the given address.
func (m mem) Read(addr uint16) byte {
	return m[addr]
}

// ReadAll returns the full memory contents.
func (m mem) ReadAll() []byte {
	return m
}

// Write writes the value v into memory at the given address.
func (m mem) Write(addr uint16, v byte) {
	m[addr] = v
}


var debug = flag.Bool("debug", false, "Run the emulator in debug mode")

func TestCPU(t *testing.T) {
	testHarness(t, filepath.Join("testdata", "8080PRE.COM"))
	fmt.Println()

	testHarness(t, filepath.Join("testdata", "TST8080.COM"))
	fmt.Println()

	testHarness(t, filepath.Join("testdata", "CPUTEST.COM"))
	fmt.Println()

	testHarness(t, filepath.Join("testdata", "8080EXM.COM"))
}

func testHarness(t *testing.T, rom string) {
	fmt.Println("*******************")

	// Instantiate 64K of memory.
	mem := make(mem, 65536)

	// Load the test ROM.
	rf, err := os.Open(rom)
	if err != nil {
		t.Fatal(err)
	}
	defer rf.Close()

	// The test ROM assumes the program code starts at 0x100. So read the ROM
	// into memory with this as an offset.
	if _, err = rf.Read(mem[0x100:]); err != nil {
		t.Fatal(err)
	}

	// Manually set the first instruction in the memory to be a JMP to 0x100.
	// This will force the emulation to start at the point where the ROM expects.
	mem.Write(0, 0xc3)
	mem.Write(1, 0)
	mem.Write(2, 0x01)

	// Fix a bug in the test ROM where it does not return from the final success
	// message.
	mem.Write(0x0005, 0xc9)

	var opts []Option
	if *debug {
		opts = append(opts, WithDebugEnabled())
	}
	i80 := NewIntel8080(mem, opts...)

	for {
		if i80.halted {
			t.Fatal("unexpected halt")
		}

		if err := i80.Step(); err != nil {
			t.Fatal(err)
		}

		// Emulate the standard out process implemented in CP/M OS in order to
		// allow us to see the output from the ROM.
		//
		// See: https://en.wikipedia.org/wiki/CP/M
		if i80.pc == 0x05 {
			if i80.r[C] == 0x09 {
				addr := uint16(i80.r[D])<<8 | uint16(i80.r[E])

				for {
					c := mem.Read(addr)

					if fmt.Sprintf("%c", c) == "$" {
						break
					} else {
						addr++
					}

					fmt.Printf("%c", c)
				}
			}
			if i80.r[C] == 0x02 {
				fmt.Printf("%c", i80.r[E])
			}
		}

		if i80.pc == 0x00 {
			break
		}
	}

	fmt.Println()
	fmt.Println("*******************")
}
