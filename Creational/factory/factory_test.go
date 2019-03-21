package factory

import "testing"

func TestAbstractFactory(t *testing.T) {

	factory, err := GetFactoryByCountryCode(RU)
	// Далее клиентский код не меняется для любой фабрики и любого семейства объектов.
	if err != nil {
		P(err.Error())
		return
	}
	mouse := factory.GetMouse()
	keyboard := factory.GetKeyboard()
	touchpad := factory.GetTouchpad()

	mouse.Click()
	keyboard.Println()
	touchpad.Track(20, 45)
}
