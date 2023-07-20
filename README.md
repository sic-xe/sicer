# SICer: A SIC(/XE) Example Runner

This repository contains source code for an example implementation of a SIC/XE
assembler and simulator in Go.

## Features

- Terminal UI (TUI)
- Generate object files from assembly code
- Run generated object files
- Normal or step by step (debug) execution
- View register states
- View memory contents

## Feature wishlist

- [ ] Linking support
- [ ] Support for all SIC/XE features
- [ ] Graphical output (Virtual screen)
- [ ] A nicer TUI using [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [ ] A GUI using [gotk4](https://github.com/diamondburned/gotk4)
- [ ] Undo feature in step-by-step debugger (able to undo up to n operations)

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
