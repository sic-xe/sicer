package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sic-xe/sicer/pkg/simulator"
)

var (
	debugFlag bool
	helpFlag  bool

	simCmd    = flag.NewFlagSet("sim", flag.ExitOnError)
	simInFlag = simCmd.String("i", "", "Input object file path (.obj) [required]")
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		help()
		os.Exit(0)
	}

	for _, fs := range []*flag.FlagSet{simCmd} {
		fs.BoolVar(&debugFlag, "d", false, "Show debug info")
		fs.BoolVar(&helpFlag, "h", false, "Show available flags")
	}

	switch os.Args[1] {
	case "sim":
		simCmd.Parse(os.Args[2:])
		runSimulator()
	default:
		fmt.Printf("Unknown subcommand '%s'\n\n", os.Args[1])
		help()
		os.Exit(1)
	}
}

func help() {
	fmt.Println(`Usage: sicer <subcommand> [flags]

Subcommands:
	sim    Run object code in the simulator

Global flags:
	-d    Show debug info
	-h    Show available flags`)
}

func runSimulator() {
	// Print usage info if requested
	if helpFlag {
		simCmd.Usage()
		os.Exit(0)
	}

	// Check if an input file was provided
	if simInFlag == nil || *simInFlag == "" {
		fmt.Printf("No input file provided!\n\n")
		simCmd.Usage()
		os.Exit(1)
	}

	// Check if the input file exists
	inFile := *simInFlag
	if _, err := os.Stat(inFile); os.IsNotExist(err) {
		fmt.Printf("Input file %s does not exist!\n\n", inFile)
		simCmd.Usage()
		os.Exit(1)
	}

	simulator.Start(debugFlag, inFile)
}
