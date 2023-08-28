package simulator

import "fmt"

// Execute decodes an instruction and executes it
func (m *Machine) Execute() error {
	opcode := m.fetch()

	var err error
	switch m.instrType(opcode) {
	case "F1":
		err = m.execInstructionF1(Opcode(opcode))
	case "F2":
		err = m.execInstructionF2(Opcode(opcode), m.fetch())
	case "F3F4":
		// Last two bits of opcode are ni bits, so they have to be masked
		opcode = opcode & 0xFC

		// Check ni bits
		indirect := opcode&0x02 == 0x02
		immediate := opcode&0x01 == 0x01

		err = m.execInstructionSICF3F4(Opcode(opcode), m.fetch(), indirect, immediate)
	default:
		logger.Error("Unknown instruction type")
	}

	return err
}

// fetch fetches and returns a byte from memory at PC
func (m *Machine) fetch() byte {
	pc := m.Registers[PC]

	instr, err := m.Byte(pc)
	if err != nil {
		logger.Error("Could not fetch instruction: %v", err)
		return 0
	}

	m.SetRegister(PC, pc+1)

	return instr
}

// isType returns the type of instruction
func (m *Machine) instrType(op byte) string {
	for _, inst := range InstrF1 {
		if byte(inst) == op {
			return "F1"
		}
	}

	for _, inst := range InstrF2 {
		if byte(inst) == op {
			return "F2"
		}
	}

	for _, inst := range InstrF3F4 {
		if byte(inst) == op {
			return "F3F4"
		}
	}

	return ""
}

// execInstructionF1 executes a format 1 instruction
func (m *Machine) execInstructionF1(op Opcode) error {
	var err error

	switch op {
	case FIX:
		err = fmt.Errorf(NotImplementedError, "FIX")
	case FLOAT:
		err = fmt.Errorf(NotImplementedError, "FLOAT")
	case HIO:
		err = fmt.Errorf(NotImplementedError, "HIO")
	case NORM:
		err = fmt.Errorf(NotImplementedError, "NORM")
	case SIO:
		err = fmt.Errorf(NotImplementedError, "SIO")
	case TIO:
		err = fmt.Errorf(NotImplementedError, "TIO")
	}

	return err
}

// execInstructionF2 executes a format 2 instruction
func (m *Machine) execInstructionF2(op Opcode, operand byte) error {
	var err error

	switch op {
	case ADDR:
		err = fmt.Errorf(NotImplementedError, "ADDR")
	case CLEAR:
		err = fmt.Errorf(NotImplementedError, "CLEAR")
	case COMPR:
		err = fmt.Errorf(NotImplementedError, "COMPR")
	case DIVR:
		err = fmt.Errorf(NotImplementedError, "DIVR")
	case MULR:
		err = fmt.Errorf(NotImplementedError, "MULR")
	case RMO:
		err = fmt.Errorf(NotImplementedError, "RMO")
	case SHIFTL:
		err = fmt.Errorf(NotImplementedError, "SHIFTL")
	case SHIFTR:
		err = fmt.Errorf(NotImplementedError, "SHIFTR")
	case SUBR:
		err = fmt.Errorf(NotImplementedError, "SUBR")
	case SVC:
		err = fmt.Errorf(NotImplementedError, "SVC")
	case TIXR:
		err = fmt.Errorf(NotImplementedError, "TIXR")
	}

	return err
}

// execInstructionF3F4 executes a format 3/4 or SIC instruction
func (m *Machine) execInstructionSICF3F4(op Opcode, operand byte, indirect, immediate bool) error {
	var err error

	switch op {
	case ADD:
		err = fmt.Errorf(NotImplementedError, "ADD")
	case ADDF:
		err = fmt.Errorf(NotImplementedError, "ADDF")
	case AND:
		err = fmt.Errorf(NotImplementedError, "AND")
	case COMP:
		err = fmt.Errorf(NotImplementedError, "COMP")
	case COMPF:
		err = fmt.Errorf(NotImplementedError, "COMPF")
	case DIV:
		err = fmt.Errorf(NotImplementedError, "DIV")
	case DIVF:
		err = fmt.Errorf(NotImplementedError, "DIVF")
	case J:
		err = fmt.Errorf(NotImplementedError, "J")
	case JEQ:
		err = fmt.Errorf(NotImplementedError, "JEQ")
	case JGT:
		err = fmt.Errorf(NotImplementedError, "JGT")
	case JLT:
		err = fmt.Errorf(NotImplementedError, "JLT")
	case JSUB:
		err = fmt.Errorf(NotImplementedError, "JSUB")
	case LDA:
		err = fmt.Errorf(NotImplementedError, "LDA")
	case LDB:
		err = fmt.Errorf(NotImplementedError, "LDB")
	case LDCH:
		err = fmt.Errorf(NotImplementedError, "LDCH")
	case LDF:
		err = fmt.Errorf(NotImplementedError, "LDF")
	case LDL:
		err = fmt.Errorf(NotImplementedError, "LDL")
	case LDS:
		err = fmt.Errorf(NotImplementedError, "LDS")
	case LDT:
		err = fmt.Errorf(NotImplementedError, "LDT")
	case LDX:
		err = fmt.Errorf(NotImplementedError, "LDX")
	case LPS:
		err = fmt.Errorf(NotImplementedError, "LPS")
	case MUL:
		err = fmt.Errorf(NotImplementedError, "MUL")
	case MULF:
		err = fmt.Errorf(NotImplementedError, "MULF")
	case OR:
		err = fmt.Errorf(NotImplementedError, "OR")
	case RD:
		err = fmt.Errorf(NotImplementedError, "RD")
	case RSUB:
		err = fmt.Errorf(NotImplementedError, "RSUB")
	case SSK:
		err = fmt.Errorf(NotImplementedError, "SSK")
	case STA:
		err = fmt.Errorf(NotImplementedError, "STA")
	case STB:
		err = fmt.Errorf(NotImplementedError, "STB")
	case STCH:
		err = fmt.Errorf(NotImplementedError, "STCH")
	case STF:
		err = fmt.Errorf(NotImplementedError, "STF")
	case STI:
		err = fmt.Errorf(NotImplementedError, "STI")
	case STL:
		err = fmt.Errorf(NotImplementedError, "STL")
	case STS:
		err = fmt.Errorf(NotImplementedError, "STS")
	case STSW:
		err = fmt.Errorf(NotImplementedError, "STSW")
	case STT:
		err = fmt.Errorf(NotImplementedError, "STT")
	case STX:
		err = fmt.Errorf(NotImplementedError, "STX")
	case SUB:
		err = fmt.Errorf(NotImplementedError, "SUB")
	case SUBF:
		err = fmt.Errorf(NotImplementedError, "SUBF")
	case TD:
		err = fmt.Errorf(NotImplementedError, "TD")
	case TIX:
		err = fmt.Errorf(NotImplementedError, "TIX")
	case WD:
		err = fmt.Errorf(NotImplementedError, "WD")
	}

	return err
}
