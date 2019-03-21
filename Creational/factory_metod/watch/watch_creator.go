// Возвращает экземпляр конкретного объекта
package watch

import (
	"errors"
)

const (
	DigitalWatchName = "DigitalWatch"
	NanoWatchName = "NanoWatch"
)

type Watch interface {
	ShowTime()
}
// Создатель не обязательно должен быть абстрактным.
type Creator interface {
	NewWatch() Watch	// фабричный метод
}

func GetCreatorByName(name string) (Creator, error) {
	switch name {
	case DigitalWatchName:
		return new(DigitalWatchCreator), nil
	case NanoWatchName:
		return new(NanoWatchCreator), nil
	default:
		return nil, errors.New("Такие часы не выпускаются!")
	}
}
