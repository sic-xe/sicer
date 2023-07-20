package simulator

import (
	"fmt"

	"github.com/sic-xe/sicer/pkg/common"
)

type register = int

var (
	A  register = 0
	X  register = 1
	L  register = 2
	B  register = 3
	S  register = 4
	T  register = 5
	F  register = 6
	PC register = 8
	SW register = 9
)

func (m *Machine) initRegisters() {
	m.registers = map[register]int{
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

// Register returns the value of the register with the given ID.
func (m *Machine) Register(r register) (int, error) {
	if !common.IsRegister(r) {
		return 0, fmt.Errorf("register with ID %d does not exist", r)
	}

	return m.registers[r], nil
}

// SetRegister sets the value of the register with the given ID.
func (m *Machine) SetRegister(r register, i int) error {
	if !common.IsRegister(r) {
		return fmt.Errorf("register with ID %d does not exist", r)
	}

	// If it's the F register it can hold two words, otherwise just one
	if (r == 6 && !common.IsTwoWords(i)) || !common.IsWord(i) {
		return fmt.Errorf("value %d can't fit in the register", i)
	}

	m.registers[r] = i

	return nil
}
