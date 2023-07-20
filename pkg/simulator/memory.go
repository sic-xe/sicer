package simulator

import (
	"encoding/binary"
	"fmt"
	"math"
)

// TODO: Think about if bytes/words should be signed or unsigned

type memory struct {
	bytes []byte
}

// Byte returns a SIC/XE byte (8 bits) from memory at address addr
func (m *memory) Byte(addr uint32) (float64, error) {
	if !IsAddress(addr) {
		return 0, fmt.Errorf("address %d is not valid", addr)
	}

	return float64(m.bytes[addr]), nil
}

// Word returns a SIC/XE word (3 bytes) from memory at address addr
func (m *memory) Word(addr uint32) (float64, error) {
	if !IsAddress(addr) {
		return 0, fmt.Errorf("address %d is not valid", addr)
	}

	// Check that addr + 2 is a valid address
	if !IsAddress(addr + 2) {
		return 0, fmt.Errorf("too close to end of memory to get word")
	}

	return math.Float64frombits(
		binary.BigEndian.Uint64([]byte{
			m.bytes[addr],
			m.bytes[addr+1],
			m.bytes[addr+2],
		})), nil
}

// Float returns a SIC/XE float (6 bytes) from memory at address addr
func (m *memory) Float(addr uint32) (float64, error) {
	if !IsAddress(addr) {
		return 0, fmt.Errorf("address %d is not valid", addr)
	}

	// Check that addr + 5 is a valid address
	if !IsAddress(addr + 5) {
		return 0, fmt.Errorf("too close to end of memory to get word")
	}

	return math.Float64frombits(
		binary.BigEndian.Uint64([]byte{
			m.bytes[addr],
			m.bytes[addr+1],
			m.bytes[addr+2],
			m.bytes[addr+3],
			m.bytes[addr+4],
			m.bytes[addr+5],
		})), nil
}
