package simulator

type Machine struct {
}

func (m *Machine) Init() {
	if debug {
		logger.Debug("Initialized machine")
	}
}
