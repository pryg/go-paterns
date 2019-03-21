package watch

import (
	"fmt"
	"time"
)

type NanoWatch struct {
}

func (w *NanoWatch) ShowTime() {
	fmt.Printf("RFC3339Nano: ", time.Now().Format(time.RFC3339Nano))
}

type NanoWatchCreator struct{
}
// фабричный метод
func (m *NanoWatchCreator) NewWatch() Watch {
	return new(NanoWatch)
}
