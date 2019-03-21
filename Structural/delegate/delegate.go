package delegate

import "fmt"

type Graphics interface {
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

type Painter struct {
	graphics Graphics
}

// Медод "мутатор". Указывает в кого мутирует Painter
func (p *Painter) SetGrafics(g Graphics) {
	p.graphics = g
}
func (p *Painter) Draw() {
	p.graphics.Draw()
}
