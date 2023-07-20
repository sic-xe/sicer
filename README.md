# SICer: A SIC(/XE) Example Runner

This repository contains source code for an example implementation of a SIC/XE
assembler and simulator in Go.

SICer is in the process of a complete rewrite since I originally wrote it as a
university course project as a way to learn Go, so it's probably full of bugs
and bad design choices.
The original code is available in the [old](https://github.com/sic-xe/sicer/tree/old)
branch.

## Features

### Simulator

- [ ] Registers
- [ ] Memory
- [ ] I/O devices
- [ ] All addressing modes
- [ ] Integer, floating point and bitwise arithmetic
- [ ] Jumps
- [ ] Load and store
- [ ] Define all available instructions
- [ ] Object code parser
- [ ] All 4 (5) instruction sets
- [ ] Timer for running instructions in sequence
- [ ] Automatic running (start, stop and step modes)
- [ ] System instructions
- [ ] Multi-processor architecture (concurrency)
- [ ] Dissassembler
- [ ] See memory contents during execution
- [ ] See register values during execution
- [ ] Detect halt
- [ ] Adjustable execution speed ("processor speed")
- [ ] Keyboard input
- [ ] Breakpoints and watches
- [ ] Interrupts
- [ ] Basic CLI interface
- [ ] Terminal UI using [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [ ] GUI using [gotk4](https://github.com/diamondburned/gotk4)
- [ ] Undo feature in step-by-step debugger (able to undo up to n operations)
- [ ] Graphical output (Virtual screen) - GUI mode only

### Assembler

- [ ] Parser for assembly code
- [ ] Symbol resolver
- [ ] Object code generator
- [ ] Compiler directives (START, END, ORG, EQU)
- [ ] Compiler base directives (BASE, NOBASE)
- [ ] LTORG compiler directive
- [ ] Full relocation support
- [ ] Control sections
- [ ] Macros
- [ ] Print intermediate code (LST) with comments
- [ ] Various compilter optimizations
    - [ ] Calculate constants during assembly
    - [ ] Remove dead code

### Linker

- [ ] Link multiple object files into one
- [ ] Multiple control sections
- [ ] Adjustable linker settings
- [ ] Inspecting and reodering control sections/symbols
- [ ] Remove or keep symbols in the output file

## Instructions

1. Install prerequisites
    - [Go](https://go.dev/)
1. Build the project
    - Automatically: `make`
    - Manually: `go build -v github.com/sic-xe/sicer`
1. Run the program
    - Run `./sicer`
    - To get additional info run `./sicer -h`

## Licensing

The project is licensed under the [GNU GPLv3](LICENSE) license.
