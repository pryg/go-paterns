package complex_builder

import "patterns/Creational/builder/car"

// Абстрактный класс объявляем через интерфейс и
type AbstractCarBuilder interface {
	CreateCar()
	GetCar() *car.Car
	BuildMake()
	BuildTransmission()
	BuildMaxSpeed()
}

// и "абстракного типа" построителя, который реализует часть статичных методов
type CarBuilder struct {
	Car *car.Car
}

func (b *CarBuilder) CreateCar() {
	b.Car = &car.Car{}
}

func (b *CarBuilder) GetCar() *car.Car {
	return b.Car
}


