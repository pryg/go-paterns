package complex_builder

import "patterns/Creational/builder/car"

// Директор имеет ссылку только на абстрактных построителей
type Director struct {
	builder AbstractCarBuilder
}

// Приказывает, что строить
func (d *Director) SetBuilder(builder AbstractCarBuilder) {
	d.builder = builder
}

func (d *Director) BuildCar() *car.Car {
	d.builder.CreateCar()
	d.builder.BuildMake()
	d.builder.BuildTransmission()
	d.builder.BuildMaxSpeed()
	return d.builder.GetCar()
}
