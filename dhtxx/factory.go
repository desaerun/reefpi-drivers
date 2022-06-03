package dhtxx

import (
	"errors"
	"fmt"
	"sync"

	"github.com/reef-pi/hal"
)

const pinNumber = "Pin"

type factory struct {
	meta       hal.Metadata
	parameters []hal.ConfigParameter
}

var f *factory
var once sync.Once

// DHT22Factory returns a singleton pH board Driver factory
func DHT22Factory() hal.DriverFactory {
	once.Do(func() {
		f = &factory{
			meta: hal.Metadata{
				Name:         "DHT22",
				Description:  "DHT22 humidity and temperature sensor",
				Capabilities: []hal.Capability{hal.AnalogInput},
			},
			parameters: []hal.ConfigParameter{
				{
					Name:    pinNumber,
					Type:    hal.Integer,
					Order:   0,
					Default: 17,
				},
			},
		}
	})
	return f
}

func (f *factory) Metadata() hal.Metadata {
	return f.meta
}

func (f *factory) GetParameters() []hal.ConfigParameter {
	return f.parameters
}

func (f *factory) ValidateParameters(parameters map[string]interface{}) (bool, map[string][]string) {
	var failures = make(map[string][]string)
	if v, ok := parameters[pinNumber]; ok {
		val, ok := hal.ConvertToInt(v)
		if !ok {
			failure := fmt.Sprint(pinNumber, " is not a number. ", v, " was received.")
			failures[pinNumber] = append(failures[pinNumber], failure)
		}
		if val <= 0 || val >= 28 {
			failure := fmt.Sprint(pinNumber, " is out of range (1 - 27). ", v, " was received.")
			failures[pinNumber] = append(failures[pinNumber], failure)
		}
	} else {
		failure := fmt.Sprint(pinNumber, " is a required parameter, but was not received.")
		failures[pinNumber] = append(failures[pinNumber], failure)
	}

	return len(failures) == 0, failures
}

func (f *factory) NewDriver(parameters map[string]interface{}, hardwareResources interface{}) (hal.Driver, error) {
	if valid, failures := f.ValidateParameters(parameters); !valid {
		return nil, errors.New(hal.ToErrorString(failures))
	}
	intAddress, _ := hal.ConvertToInt(parameters[pinNumber])
	return NewDriver(intAddress, f.meta)
}
