package main

import (
	"machine"
	"strconv"
	"time"
)

const ThresholdLight = 300

func log(value uint16) {
	val := strconv.Itoa(int(value))
	machine.Serial.Write([]byte(val))
	machine.Serial.Write([]byte{'\n'})
}

func main() {
	machine.InitADC()
	// define pins
	ldrPin := machine.ADC{Pin: machine.PC0}
	rPin := machine.D4
	pirPin := machine.D7

	// configure pins
	rPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pirPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	ldrPin.Configure(machine.ADCConfig{})

	for {
		lightValue := ldrPin.Get() - 65000
		isPersonDetected := pirPin.Get()
		if isPersonDetected {
			rPin.High()
		} else {
			rPin.Low()
		}
		log(lightValue)
		time.Sleep(time.Millisecond * 2000)
	}
}
