package factory

import "math"

// интерфейсы семейства объектов
type Mouse interface {
	Click()
	DBClick()
	Scroll(detection int)
}

type Keyboard interface {
	Print()
	Println()
}

type Touchpad interface {
	Track(deltaX, deltaY int)
}

// Первое семейство объектов (Ru)
type RuMouse struct{}

func (m *RuMouse) Click() {
	P("Щелчок мышью")
}

func (m *RuMouse) DBClick() {
	P("Двойной щелчок мышью")
}

func (m *RuMouse) Scroll(detection int) {
	if detection > 0 {
		P("Скролим вверх")
	} else if detection < 0 {
		P("Скролим вниз")
	} else {
		P("Не скролим")
	}
}

type RuKeyboard struct{}

func (k *RuKeyboard) Print() {
	P("Печатаем строку")
}

func (k *RuKeyboard) Println() {
	P("Печатаем строку с переводом")
}

type RuTouchpad struct{}

func (t *RuTouchpad) Track(deltaX, deltaY int) {
	s := int(math.Sqrt(math.Pow(float64(deltaX), 2) + math.Pow(float64(deltaY), 2)))
	P("Передвинулись на ", s, " пикселей.")
}

// Второе семейство объектов (Ru)
type EnMouse struct{}

func (m *EnMouse) Click() {
	P("Mouse click")
}

func (m *EnMouse) DBClick() {
	P("Mouse double cklick")
}

func (m *EnMouse) Scroll(detection int) {
	if detection > 0 {
		P("Scroll Up")
	} else if detection < 0 {
		P("Scroll Down")
	} else {
		P("No scrolling")
	}
}

type EnKeyboard struct{}

func (k *EnKeyboard) Print() {
	P("Print")
}

func (k *EnKeyboard) Println() {
	P("Print line")
}

type EnTouchpad struct{}

func (t *EnTouchpad) Track(deltaX, deltaY int) {
	s := int(math.Sqrt(math.Pow(float64(deltaX), 2) + math.Pow(float64(deltaY), 2)))
	P("Moved ", s, " pixels")
}
