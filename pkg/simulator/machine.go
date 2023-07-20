package simulator

type Machine struct {
	registers map[register]int
}

func (m *Machine) Init() {
	m.initRegisters()

	if debug {
		logger.Debug("Initialized machine")
	}
}
