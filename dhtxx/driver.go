package dhtxx

import (
	"fmt"

	"github.com/reef-pi/hal"
)

type Driver struct {
	meta     hal.Metadata
	channels []hal.AnalogInputPin
}

func NewDriver(pin int, meta hal.Metadata) (*Driver, error) {
	s := &DHT22{
		pin: pin,
	}
	ch1, err := newChannel(s, 0)
	if err != nil {
		return nil, err
	}
	ch2, err := newChannel(s, 1)
	if err != nil {
		return nil, err
	}
	return &Driver{
		meta:     meta,
		channels: []hal.AnalogInputPin{ch1, ch2},
	}, nil
}

func (d *Driver) Metadata() hal.Metadata {
	return d.meta
}

func (d *Driver) Pins(cap hal.Capability) ([]hal.Pin, error) {
	if cap == hal.AnalogInput {
		return []hal.Pin{d.channels[0], d.channels[1]}, nil
	}
	return nil, fmt.Errorf("unsupported capability: %s", cap.String())
}

func (d *Driver) AnalogInputPins() []hal.AnalogInputPin {
	return d.channels
}

func (d *Driver) AnalogInputPin(n int) (hal.AnalogInputPin, error) {
	if n < 0 || n > 1 {
		return nil, fmt.Errorf("DHT22 board does not have channel %d", n)
	}
	return d.channels[n], nil
}

func (d *Driver) Close() error {
	return nil
}
