package dhtxx

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"time"
)

const (
	RETRIED = 10
)

type DHT22 struct {
	pin              int
	pTemp, pHumidity float32
	pTime            time.Time
}

func (d *DHT22) read() (float32, float32, int) {
	if pTemp, pHumidity, err, _ := dht.ReadDHTxxWithRetry(dht.DHT22, d.pin, false, RETRIED); err != 0 {
		return 0, 0, err
	} else {
		return pTemp, pHumidity, 0
	}
}

func (d *DHT22) ReadSensor() (float32, float32, int) {
	if temp, humidity, err, _ := dht.ReadDHTxxWithRetry(dht.DHT22, d.pin, false, RETRIED); err != 0 {
		return 0, 0, err
	} else {
		d.pTemp = temp
		d.pHumidity = humidity
		d.pTime = time.Now()

		return temp, humidity, 0
	}
}

func (s *DHT22) Temperature() (float64, error) {
	if s.pTime.Before(time.Now().Add(time.Second)) {
		if _, _, err := s.ReadSensor(); err != 0 {
			return 0, fmt.Errorf("error reading temp sensor: %v", err)
		}
	}

	return float64(s.pTemp), nil
}

func (s *DHT22) Humidity() (float64, error) {
	if s.pTime.Before(time.Now().Add(time.Second)) {
		if _, _, err := s.ReadSensor(); err != 0 {
			return 0, fmt.Errorf("error reading humidity sensor: %v", err)
		}
	}
	return float64(s.pHumidity), nil
}
