package factory_metod

import (
	"testing"
	"fmt"
	"patterns/Creational/factory_metod/watch"
)

func TestFactoryMetod(t *testing.T) {
	/* По простому на выходе конкретный объект, что не всегда удобно.
	watch := new(NanoWatch)
	*/

	// с фабричным методом на выходе интерфейс
	// Первый пример, сделать создателя напрямую:
	var watchCreator watch.Creator
	watchCreator = new(watch.DigitalWatchCreator)

	// Второй способ, сделать создателя специальным методом по имени (или перечислению):
	var err error
	watchCreator, err = watch.GetCreatorByName(watch.NanoWatchName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// дальнейший код не меняется для создания любых объектов
	watch := watchCreator.NewWatch()
	watch.ShowTime()
}
