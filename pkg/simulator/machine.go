package simulator

type Machine struct {
	registers map[registerName]register
	memory    memory
}

func (m *Machine) Init() {
	m.initRegisters()

	if debug {
		logger.Debug("Initialized machine")
	}
}
