package simulator

func Start(isDebug bool, inputFile string) {
	debug = isDebug
	if debug {
		logger.Debug("Starting simulator")
	}

	var m Machine
	m.Init()
}
