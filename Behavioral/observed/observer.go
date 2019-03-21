package observed

import "fmt"

type Observer interface {
	HandleEvent(temperature, pressure int)
}

type ConsolObserver struct{}

func (co *ConsolObserver) HandleEvent(temperature, pressure int) {
	fmt.Printf("Погода измеилась. {Температура: %d, Давление: %d}\n", temperature, pressure)
}

func NewConsoleObserver() Observer {
	return new(ConsolObserver)
}
