package simulator

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Byte returns a SIC/XE byte (8 bits) value from memory at address addr
func (m *Machine) Byte(addr float64) (byte, error) {
	if !IsAddress(addr) {
		return 0, fmt.Errorf("address %f is not valid", addr)
	}

	return m.Memory[int(addr)], nil
}

// Word returns a SIC/XE word (3 bytes) value from memory at address addr
func (m *Machine) Word(addr float64) (float64, error) {
	if !IsAddress(addr + 2) {
		return 0, fmt.Errorf("too close to end of memory to get word")
	}

	return math.Float64frombits(
		binary.LittleEndian.Uint64([]byte{
			m.Memory[int(addr)],
			m.Memory[int(addr)+1],
			m.Memory[int(addr)+2],
		})), nil
}

// Float returns a SIC/XE float (6 bytes) from memory at address addr
func (m *Machine) Float(addr float64) (float64, error) {
	if !IsAddress(addr + 5) {
		return 0, fmt.Errorf("too close to end of memory to get float")
	}

	return math.Float64frombits(
		binary.LittleEndian.Uint64([]byte{
			m.Memory[int(addr)],
			m.Memory[int(addr)+1],
			m.Memory[int(addr)+2],
			m.Memory[int(addr)+3],
			m.Memory[int(addr)+4],
			m.Memory[int(addr)+5],
		})), nil
}

// SetByte sets a SIC/XE byte (8 bits) value in memory at address addr
func (m *Machine) SetByte(addr float64, val byte) error {
	if !IsAddress(addr) {
		return fmt.Errorf("address %f is not valid", addr)
	}

	m.Memory[int(addr)] = val

	return nil
}

// SetWord sets a SIC/XE word (3 bytes) value in memory at address addr
func (m *Machine) SetWord(addr float64, val float64) error {
	if !IsAddress(addr + 2) {
		return fmt.Errorf("too close to end of memory to set word")
	}

	bits := math.Float64bits(val)

	m.Memory[int(addr)] = byte(bits)
	m.Memory[int(addr)+1] = byte(bits >> 8)
	m.Memory[int(addr)+2] = byte(bits >> 16)

	return nil
}

// SetFloat sets a SIC/XE float (6 bytes) value in memory at address addr
func (m *Machine) SetFloat(addr float64, val float64) error {
	if !IsAddress(addr + 5) {
		return fmt.Errorf("too close to end of memory to set float")
	}

	bits := math.Float64bits(val)

	m.Memory[int(addr)] = byte(bits)
	m.Memory[int(addr)+1] = byte(bits >> 8)
	m.Memory[int(addr)+2] = byte(bits >> 16)
	m.Memory[int(addr)+3] = byte(bits >> 24)
	m.Memory[int(addr)+4] = byte(bits >> 32)
	m.Memory[int(addr)+5] = byte(bits >> 40)

	return nil
}
