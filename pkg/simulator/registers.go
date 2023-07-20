package simulator

import (
	"fmt"
)

type registerName string

const (
	A  registerName = "A"
	X  registerName = "X"
	L  registerName = "L"
	B  registerName = "B"
	S  registerName = "S"
	T  registerName = "T"
	F  registerName = "F"
	PC registerName = "PC"
	SW registerName = "SW"
)

type register struct {
	name  registerName
	value float64
}

func (m *Machine) initRegisters() {
	m.registers = map[registerName]register{
		A:  {name: A},
		X:  {name: X},
		L:  {name: L},
		B:  {name: B},
		S:  {name: S},
		T:  {name: T},
		F:  {name: F},
		PC: {name: PC},
		SW: {name: SW},
	}
}

func (r *register) Value() float64 {
	return r.value
}

func (r *register) SetValue(v float64) error {
	if (r.name != F && !IsWord(v)) || !IsTwoWords(v) {
		return fmt.Errorf("value %f can't fit in register %s", v, r.name)
	}
	r.value = v

	return nil
}

func (r *register) SetLowValue(v float64) error {
	if !IsByte(v) {
		return fmt.Errorf("value %f can't fit in register %s", v, r.name)
	}
	r.value = v

	return nil
}
