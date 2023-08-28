package simulator

import (
	"bufio"
	"fmt"
)

type DeviceType uint8

const (
	Unknown DeviceType = iota
	Reader
	Writer
)

type Device struct {
	t DeviceType
	r *bufio.Reader
	w *bufio.Writer
}

// Test checks if a device is initialized for reading or writing
func (d *Device) Test() bool {
	return d.t != Unknown && (d.t == Reader && d.r != nil || d.t == Writer && d.w != nil)
}

// Read reads a byte from the device
func (d *Device) Read() (float64, error) {
	if d.t != Reader || d.r == nil {
		return 0, fmt.Errorf("device is not initialized for reading")
	}

	return 0, nil
}

// Write writes a byte to the device
func (d *Device) Write(val float64) error {
	if d.t != Writer || d.w == nil {
		return fmt.Errorf("evice is not initialized for writing")
	}

	if !IsByte(val) {
		return fmt.Errorf("value %f is not a byte", val)
	}

	return nil
}

// InitDevice initializes a device for reading or writing
func InitDevice(t DeviceType) (*Device, error) {
	if t == Unknown {
		return nil, fmt.Errorf("device must be a reader or a writer")
	}

	return &Device{t: t}, nil
}
