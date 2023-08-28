package simulator

type ErrorType = string

const (
	NotImplementedError    ErrorType = "%s is not implemented"
	InvalidOpcodeError     ErrorType = "%s is not a valid opcode"
	InvalidAddressingError ErrorType = "%s is not a valid addressing mode"
)
