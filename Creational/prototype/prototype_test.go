package prototype

import (
	"fmt"
	"testing"
)

func TestProtoType(t *testing.T) {

	// Нужны копии объектов
	humanOriginal := NewHuman(25, "Andrey")
	fmt.Printf("HumanOriginal{age: %d, name: %s}\n", humanOriginal.GetAge(), humanOriginal.GetName())

	// Первый способ по простому с преобразованием типа
	humanCopy1 := humanOriginal.Copy().(*Human)
	fmt.Printf("HumanCopy1{age: %d, name: %s}\n", humanCopy1.GetAge(), humanCopy1.GetName())

	// Второй способ с фабрикой
	humanFactory := CreateHumanFactory()
	humanFactory.SetPrototype(humanOriginal)
	humanCopy2 := humanFactory.MakeCopy()

	fmt.Printf("HumanCopy2{age: %d, name: %s}\n", humanCopy2.GetAge(), humanCopy2.GetName())
}
