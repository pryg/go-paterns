package complex_builder

import (
	"patterns/Creational/builder/car"
)

// Второй конкретный построитель
type SubaruBuilder struct {
	CarBuilder
}

// переопределяем методы конкретного построителя
func (fb *SubaruBuilder) BuildMake() {
	fb.GetCar().SetMake("Subaru")
}

func (fb *SubaruBuilder) BuildTransmission() {
	fb.GetCar().SetTransmission(car.AUTO)
}

func (fb *SubaruBuilder) BuildMaxSpeed() {
	fb.GetCar().SetMaxSpeed(220)
}

