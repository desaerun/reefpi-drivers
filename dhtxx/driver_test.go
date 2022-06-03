package dhtxx

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/d2r2/go-dht"
)

func readDHT11(pin int) (temperature float32, humidity float32, err error) {
	sensorType := dht.DHT11
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(sensorType, pin, false, 10)
	if err != nil {
		return temperature, humidity, err
	}
	if retried < 10 {
		return temperature, humidity, errors.New("read data from sensor failed, retry a few more times")
	}
	return temperature, humidity, nil
}

func readDHT22(pin int) (temperature float32, humidity float32, err error) {
	sensorType := dht.DHT22
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(sensorType, pin, false, 10)
	if err != nil {
		return temperature, humidity, err
	}
	if retried < 10 {
		return temperature, humidity, errors.New("read data from sensor failed, retry a few more times")
	}
	return temperature, humidity, nil
}

func main() {
	for {
		time.Sleep(time.Second * 2)
		temperature, humidity, err := readDHT11(4)
		if err != nil {
			temperature, humidity, err = readDHT22(4)
		}
		if err == nil {
			fahrenheit := math.Round(float64((temperature*9/5+32)*100)) / 100
			celsius := math.Round(float64(temperature*100)) / 100
			fmt.Printf("Temperature = %v*C (%v*F), Humidity = %v%%\n", celsius, fahrenheit, humidity)
		} else {
			fmt.Println(err)
		}
	}
}
