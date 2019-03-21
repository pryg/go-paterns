package composite

import "fmt"

type Shape interface {
	Draw()
}

type Triangle struct{}

func (*Triangle) Draw() {
	fmt.Println("Рисуем квадрат")
}

type Square struct{}

func (*Square) Draw() {
	fmt.Println("Рисуем прямоугольник")
}

type Circle struct{}

func (*Circle) Draw() {
	fmt.Println("Рисуем круг")
}

type Composite struct {
	components []Shape
}

func (c *Composite) AddComponent(component Shape) {
	c.components = append(c.components, component)
}

func (c *Composite) RemoveComponent(component Shape) {
	// Как-то удаляем
}

func (c *Composite) Draw() {
	for _, component := range c.components {
		component.Draw()
	}
}