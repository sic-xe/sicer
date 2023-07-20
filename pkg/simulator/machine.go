package simulator

type Machine struct {
	registers map[registerName]register
}

func (m *Machine) Init() {
	m.initRegisters()

	if debug {
		logger.Debug("Initialized machine")
	}
}
