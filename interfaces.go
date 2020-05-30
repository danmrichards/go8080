package go8080

// MemReader is the interface that wraps the basic Read and ReadAll methods.
//
// Read returns the value from memory at the given address.
//
// ReadAll returns the full memory contents.
type MemReader interface {
	Read(addr uint16) byte
	ReadAll() []byte
}

// MemWriter is the interface that wraps the basic Write method.
//
// Write writes the value v into memory at the given address.
type MemWriter interface {
	Write(addr uint16, v byte)
}

// MemReadWriter is the interface that groups the basic Read and Write methods.
type MemReadWriter interface {
	MemReader
	MemWriter
}
