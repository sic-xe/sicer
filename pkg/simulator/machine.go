package simulator

type Machine struct {
	Registers map[Register]float64
	Memory    [MAX_ADDRESS]byte
	Devices   map[uint8]*Device
}

func (m *Machine) Init() {
	m.initRegisters()

	// Initialize devices
	m.Devices = make(map[uint8]*Device, 256)
	m.Devices[0], _ = InitDevice(Reader) // stdin
	m.Devices[1], _ = InitDevice(Writer) // stdout
	m.Devices[2], _ = InitDevice(Writer) // stderr

	if debug {
		logger.Debug("Initialized machine")
	}
}
