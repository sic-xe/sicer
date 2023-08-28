package simulator

import (
	"fmt"
)

func (m *Machine) initRegisters() {
	m.Registers = map[Register]float64{
		A:  0,
		X:  0,
		L:  0,
		B:  0,
		S:  0,
		T:  0,
		F:  0,
		PC: 0,
		SW: 0,
	}
}

// Register returns the value of the register
func (m *Machine) Register(name Register) float64 {
	return m.Registers[name]
}

// SetRegister sets the value of the register
func (m *Machine) SetRegister(name Register, val float64) error {
	if (name != F && !IsWord(val)) || !IsFloat(val) {
		return fmt.Errorf("value %f can't fit in register %d", val, name)
	}

	m.Registers[name] = val
	return nil
}

// SetRegisterLow sets the lowest byte of the register
func (m *Machine) SetRegisterLow(name Register, val float64) error {
	if !IsByte(val) {
		return fmt.Errorf("value %f can't fit in lowest byte of register %d", val, name)
	}

	m.Registers[name] = val
	return nil
}
