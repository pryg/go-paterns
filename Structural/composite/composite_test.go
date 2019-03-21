package composite

import "testing"

func TestComposite(t *testing.T) {

	var square1 Shape = new(Square)
	var square2 Shape = new(Square)
	var triangle1 Shape = new(Triangle)

	var square3 Shape = new(Square)
	var circle1 Shape = new(Circle)
	var circle2 Shape = new(Circle)
	var circle3 Shape = new(Circle)

	composite1 := new(Composite)
	composite2 := new(Composite)

	composite := new(Composite)

	composite1.AddComponent(square1)
	composite1.AddComponent(square2)
	composite1.AddComponent(triangle1)

	composite2.AddComponent(square3)
	composite2.AddComponent(circle1)
	composite2.AddComponent(circle2)
	composite2.AddComponent(circle3)

	// Прелесть в матрёшках
	composite.AddComponent(composite1)
	composite.AddComponent(composite2)
	composite.AddComponent(new(Triangle))

	// будут рисовать все объекты
	composite.Draw()
}