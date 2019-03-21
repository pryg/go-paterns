package complex_builder

import "patterns/Creational/builder/car"

// Первый конкретный построитель
type FordBuilder struct {
	CarBuilder
}

// переопределяем методы конкретного построителя
func (fb *FordBuilder) BuildMake() {
	fb.GetCar().SetMake("Ford")
}

func (fb FordBuilder) BuildTransmission() {
	fb.GetCar().SetTransmission(car.AUTO)
}

func (fb FordBuilder) BuildMaxSpeed() {
	fb.GetCar().SetMaxSpeed(250)
}
