package builder

import (
	"fmt"
	"patterns/Creational/builder/simple_builder"
	"testing"
	"patterns/Creational/builder/complex_builder"
)

func TestSimpleBuilder(t *testing.T) {
	// Удобство в том, что можно пропустить любой из методов построения.
	// Тогда будет использовано значение по умолчанию.
	car := simple_builder.NewCarBuilder().
		BuildMark("Lada Kalina").
		BuildMaxSpeed(200).
		Build()
	fmt.Printf(car.ToString())
}

func TestComplexBuilder(t *testing.T) {

	director := &complex_builder.Director{}
	director.SetBuilder(new(complex_builder.SubaruBuilder))
	subaru := director.BuildCar()
	fmt.Printf(subaru.ToString())
	fmt.Println()

	director.SetBuilder(new(complex_builder.FordBuilder))
	ford := director.BuildCar()
	fmt.Printf(ford.ToString())
	fmt.Println()
}

