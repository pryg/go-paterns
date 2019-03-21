// для постоения объектов не показывая конструктор
package simple_builder

import (
	"patterns/Creational/builder/car"
)

// Построитель (попроще)
type CarBuilder struct {
	m string
	s int
}

// Приходится делать конструктор, т.к. в go нельзя задавать значения по умолчанию в структуре
func NewCarBuilder() *CarBuilder {
	// значения по умолчанию отличные от значеий в конструкторе объекта
	// это только для примера, простой построитель может переопределять и не все свойства
	return &CarBuilder{"Lada", 150}
}

// Методы построения возвращают самого же строителя
func (b *CarBuilder) BuildMark(make string) *CarBuilder {
	b.m = make
	return b
}

func (b *CarBuilder) BuildMaxSpeed(maxSpeed int) *CarBuilder {
	b.s = maxSpeed
	return b
}

// В методе построения присваиваются значения по умолчанию
func (b *CarBuilder) Build() *car.Car {
	car := car.NewCar()
	car.SetMake(b.m)
	car.SetMaxSpeed(b.s)
	return car
}
