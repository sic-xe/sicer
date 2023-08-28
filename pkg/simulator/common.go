package simulator

import (
	"math"

	"github.com/sic-xe/sicer/pkg/common"
)

var (
	debug  bool
	logger common.Logger
)

type (
	Register uint8
	Opcode   byte
)

const (
	MAX_ADDRESS = 1048576 // 2^20

	A Register = iota
	X
	L
	B
	S
	T
	F
	_
	PC
	SW

	// Opcodes
	ADD    Opcode = 0x18
	ADDF   Opcode = 0x58
	ADDR   Opcode = 0x90
	AND    Opcode = 0x40
	CLEAR  Opcode = 0xB4
	COMP   Opcode = 0x28
	COMPF  Opcode = 0x88
	COMPR  Opcode = 0xA0
	DIV    Opcode = 0x24
	DIVF   Opcode = 0x64
	DIVR   Opcode = 0x9C
	FIX    Opcode = 0xC4
	FLOAT  Opcode = 0xC0
	HIO    Opcode = 0xF4
	J      Opcode = 0x3C
	JEQ    Opcode = 0x30
	JGT    Opcode = 0x34
	JLT    Opcode = 0x38
	JSUB   Opcode = 0x48
	LDA    Opcode = 0x00
	LDB    Opcode = 0x68
	LDCH   Opcode = 0x50
	LDF    Opcode = 0x70
	LDL    Opcode = 0x08
	LDS    Opcode = 0x6C
	LDT    Opcode = 0x74
	LDX    Opcode = 0x04
	LPS    Opcode = 0xD0
	MUL    Opcode = 0x20
	MULF   Opcode = 0x60
	MULR   Opcode = 0x98
	NORM   Opcode = 0xC8
	OR     Opcode = 0x44
	RD     Opcode = 0xD8
	RMO    Opcode = 0xAC
	RSUB   Opcode = 0x4C
	SHIFTL Opcode = 0xA4
	SHIFTR Opcode = 0xA8
	SIO    Opcode = 0xF0
	SSK    Opcode = 0xEC
	STA    Opcode = 0x0C
	STB    Opcode = 0x78
	STCH   Opcode = 0x54
	STF    Opcode = 0x80
	STI    Opcode = 0xD4
	STL    Opcode = 0x14
	STS    Opcode = 0x7C
	STSW   Opcode = 0xE8
	STT    Opcode = 0x84
	STX    Opcode = 0x10
	SUB    Opcode = 0x1C
	SUBF   Opcode = 0x5C
	SUBR   Opcode = 0x94
	SVC    Opcode = 0xB0
	TD     Opcode = 0xE0
	TIO    Opcode = 0xF8
	TIX    Opcode = 0x2C
	TIXR   Opcode = 0xB8
	WD     Opcode = 0xDC
)

// Opcodes by type
var (
	InstrF1   = []Opcode{FIX, FLOAT, HIO, NORM, SIO, TIO}
	InstrF2   = []Opcode{ADDR, CLEAR, COMPR, DIVR, MULR, RMO, SHIFTL, SHIFTR, SUBR, SVC, TIXR}
	InstrF3F4 = []Opcode{ADD, ADDF, AND, COMP, COMPF, DIV, DIVF, J, JEQ, JGT, JLT, JSUB, LDA,
		LDB, LDCH, LDF, LDL, LDS, LDT, LDX, LPS, MUL, MULF, OR, RD, RSUB, SSK, STA, STB, STCH,
		STF, STI, STL, STS, STSW, STT, STX, SUB, SUBF, TD, TIX, WD}
)

// IsByte checks if the value fits in a SIC/XE byte (8 bits)
func IsByte(val float64) bool {
	return val >= 0 && val <= math.Pow(2, 8)-1
}

// IsWord checks if the value fits in a SIC/XE word (3 bytes)
func IsWord(val float64) bool {
	return val >= -math.Pow(2, 23) && val <= math.Pow(2, 24)-1
}

// IsFloat checks if the value fits in a SIC/XE float (6 bytes)
func IsFloat(val float64) bool {
	return val >= -math.Pow(2, 47) && val <= math.Pow(2, 48)-1
}

// IsAddress checks if the value is a valid SIC/XE address
func IsAddress(val float64) bool {
	return val >= 0 && val < MAX_ADDRESS
}
