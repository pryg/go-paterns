package watch

import (
	"fmt"
	"time"
)

type DigitalWatch struct {
}

func (w *DigitalWatch) ShowTime() {
	fmt.Printf("RFC3339: ", time.Now().Format(time.RFC3339))
}

type DigitalWatchCreator struct{
}
// фабричный метод
func (m *DigitalWatchCreator) NewWatch() Watch {
	return new(DigitalWatch)
}