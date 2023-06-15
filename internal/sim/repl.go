package sim

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Repl runs the simulator in REPL mode
func Repl(m Machine) {
	header()
	fmt.Println("(interactive mode)")
	replHelp()

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for sc.Scan() {
		switch text := strings.Split(sc.Text(), " "); text[0] {
		case "regs", "r":
			fmt.Println(m.Regs())
		case "mem", "m":
			low, err := strconv.Atoi(text[1])

			if err != nil {
				panic(err)
			}

			high, err := strconv.Atoi(text[2])

			if err != nil {
				panic(err)
			}

			fmt.Println(m.Mem(low, high))
		case "exec", "e":
			if !m.Halted() {
				m.Execute()
			} else {
				fmt.Println("Finished executing program, " +
					"stop trying to break things")
			}
		case "step", "s":
			if !m.Halted() {
				m.Execute()
				fmt.Println(m.Regs())
			} else {
				fmt.Println("Finished executing program, " +
					"stop trying to break things")
			}
		case "word", "w":
			addr, err := strconv.Atoi(text[1])

			if err != nil {
				panic(err)
			}

			word, err := m.Word(addr)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%02X\n", word)
		case "byte", "b":
			addr, err := strconv.Atoi(text[1])

			if err != nil {
				panic(err)
			}

			byt, err := m.Byte(addr)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%02X\n", byt)
		case "setreg", "sr":
			no, err := strconv.Atoi(text[1])

			if err != nil {
				switch text[1] {
				case "a", "A":
					no = 0
				case "x", "X":
					no = 1
				case "l", "L":
					no = 2
				case "b", "B":
					no = 3
				case "s", "S":
					no = 4
				case "t", "T":
					no = 5
				case "f", "F":
					no = 6
				case "pc", "PC":
					no = 8
				case "sw", "SW":
					no = 9
				default:
					fmt.Printf("Invalid register: %s\n", text[2])
					continue
				}
			}

			val, err := strconv.Atoi(text[2])

			if err != nil {
				panic(err)
			}

			if err := m.SetReg(no, val); err != nil {
				panic(err)
			}
		case "begin", "bt":
			if !m.Halted() {
				fmt.Println("Started automatic execution")
				m.Start()
			} else {
				fmt.Println("Finished executing program, " +
					"stop trying to break things")
			}
		case "end", "et":
			if !m.Halted() {
				fmt.Println("Stopped automatic execution")
				m.Stop()
			} else {
				fmt.Println("Finished executing program, " +
					"stop trying to break things")
			}
		default:
			replHelp()
		}

		fmt.Print("> ")
	}
}

func replHelp() {
	fmt.Println("Usage: [command] (options)")
	fmt.Println()
	fmt.Println("  Memory and registers:")
	fmt.Println("    b, byte [addr]           Returns the byte at memory[addr]")
	fmt.Println("    w, word [addr]           Returns the word at memory[addr]")
	fmt.Println("    m, mem [low] [high]      Prints memory contents from " +
		"low to high address")
	fmt.Println("    r, regs                  Prints register values")
	fmt.Println("    sr, setreg [no] [val]    Sets the register [no] to [val]")
	fmt.Println()
	fmt.Println("  Instructions:")
	fmt.Println("    e, exec      Executes the next instruction")
	fmt.Println("    s, step      Executes the next instruction and prints " +
		"register values")
	fmt.Println("    bt, begin    Starts automatically executing instructions")
	fmt.Println("    et, end      Stops automatically executing instructions")
}

func header() {
	fmt.Printf(
		"███████ ██  ██████ ███████ ██ ███    ███\n" +
			"██      ██ ██      ██      ██ ████  ████\n" +
			"███████ ██ ██      ███████ ██ ██ ████ ██\n" +
			"     ██ ██ ██           ██ ██ ██  ██  ██\n" +
			"███████ ██  ██████ ███████ ██ ██      ██\n\n")
}
