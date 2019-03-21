package car

import "fmt"

var P = fmt.Sprintf

type Transmission int
const (
	MANUAL Transmission = iota
	AUTO
)

type Car struct {
	mark         string
	transmission Transmission
	maxSpeed     int
}

func NewCar() *Car {
	return &Car{"Default", MANUAL, 150}
}

func (c *Car) SetMake(make string) {
	c.mark = make
}

func (c *Car) SetTransmission(transmission Transmission) {
	c.transmission = transmission
}

func (c *Car) SetMaxSpeed(maxSpeed int) {
	c.maxSpeed = maxSpeed
}

func (c *Car) ToString() string {
	return P("Car {mark: %s, transmission: %v, maxSpeed: %d}", c.mark, c.transmission, c.maxSpeed)
}
