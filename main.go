package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sic-xe/sicer/internal/asm"
	"github.com/sic-xe/sicer/internal/sim"
)

var (
	// Shared flags
	debugFlag bool
	helpFlag  bool

	// Assembler flags
	asmCmd     = flag.NewFlagSet("asm", flag.ExitOnError)
	asmLstFlag = asmCmd.Bool("l", false, "Pretty-print object and assembly code")
	asmInFlag  = asmCmd.String("i", "", "Input assembly file path (.asm) [required]")
	asmOutFlag = asmCmd.String("o", "", "Output object file path (.obj) [required]")

	// Simulator flags
	simCmd    = flag.NewFlagSet("sim", flag.ExitOnError)
	simInFlag = simCmd.String("i", "", "Input object file path (.obj) [required]")
	simBgFlag = simCmd.Bool("b", false, "Run in noninteractive (background) mode")
)

func setupCommonFlags() {
	for _, fs := range []*flag.FlagSet{asmCmd, simCmd} {
		fs.BoolVar(&debugFlag, "d", false, "Show debug info")
		fs.BoolVar(&helpFlag, "h", false, "Show available flags")
	}
}

func help() {
	fmt.Println("Usage: sicer <subcommand> [flags]")
	fmt.Println("Subcommands:")
	fmt.Println("  asm    Assemble .asm files into object code")
	fmt.Println("  sim    Run object code in the simulator")
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		help()
		os.Exit(0)
	}

	setupCommonFlags()

	switch os.Args[1] {
	case "asm":
		asmCmd.Parse(os.Args[2:])
		runAssembler()
	case "sim":
		simCmd.Parse(os.Args[2:])
		runSimulator()
	default:
		fmt.Printf("Unknown subcommand '%s'\n\n", os.Args[1])
		help()
		os.Exit(1)
	}
}

func runAssembler() {
	// Print usage info if requested
	if helpFlag {
		asmCmd.Usage()
		os.Exit(0)
	}

	// Check if the input flag was provided
	if *asmInFlag == "" {
		fmt.Printf("No input file provided!\n\n")
		asmCmd.Usage()
		os.Exit(1)
	}

	// Check if the input file exists
	if _, err := os.Stat(*asmInFlag); os.IsNotExist(err) {
		fmt.Printf("Input file %s does not exist!\n\n", *asmInFlag)
		asmCmd.Usage()
		os.Exit(1)
	}

	// Set output to the same name as input if not specified
	if *asmOutFlag == "" {
		output, found := strings.CutSuffix(*asmInFlag, ".asm")
		if !found {
			fmt.Printf("Input file %s does not have a .asm extension!\n\n",
				*asmInFlag)
		}
		*asmOutFlag = output + ".obj"
	}

	asm.SetDebug(debugFlag)
	asm.SetPrettyPrint(*asmLstFlag)

	code := asm.NewCode()

	// First pass: parse code
	if err := code.ParseFile(*asmInFlag); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Second pass: replace variables with their values
	code.ResolveSymbols()

	if err := code.CreateObjectFile(*asmOutFlag); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *asmLstFlag {
		code.PrettyPrint()
	}
}

func runSimulator() {
	// Print usage info if requested
	if helpFlag {
		simCmd.Usage()
		os.Exit(0)
	}

	// Check if the input flag was provided
	if *simInFlag == "" {
		fmt.Printf("No input file provided!\n\n")
		simCmd.Usage()
		os.Exit(1)
	}

	// Check if the input file exists
	if _, err := os.Stat(*simInFlag); os.IsNotExist(err) {
		fmt.Printf("Input file %s does not exist!\n\n", *simInFlag)
		simCmd.Usage()
		os.Exit(1)
	}

	sim.SetDebug(debugFlag)

	// Clear screen if running in REPL mode (overwritten by debug mode)
	if !*simBgFlag {
		scr := exec.Command("clear")
		scr.Stdout = os.Stdout
		scr.Run()
	}

	// Instantiate a new machine
	var m sim.Machine
	m.New()
	m.SetInteractive(*simBgFlag)

	if err := m.ParseObjFile(*simInFlag); err != nil {
		fmt.Println(err)
	}

	if !*simBgFlag {
		sim.Repl(m)
	} else {
		m.Start()
	}
}
