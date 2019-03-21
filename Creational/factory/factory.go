// Для производства семейства(набора) объектов.
// Еще гарантирует взаимодействие между объектами одного семейства.
package factory

import (
	"errors"
	"fmt"
)

var P = fmt.Println

type NameFactory uint

// Enum
const (
	RU NameFactory = iota
	EN
)

// интерфейс фабрики
type DeviceFactory interface {
	GetMouse() Mouse
	GetKeyboard() Keyboard
	GetTouchpad() Touchpad
}

// Фабрика для создания первого семейства
type RuDeviceFactory struct{}

func (f *RuDeviceFactory) GetMouse() Mouse {
	return new(EnMouse)
}

func (f *RuDeviceFactory) GetKeyboard() Keyboard {
	return new(EnKeyboard)
}

func (f *RuDeviceFactory) GetTouchpad() Touchpad {
	return new(EnTouchpad)
}

// Фабрика для создания второго семейства
type EnDeviceFactory struct{}

func (f *EnDeviceFactory) GetMouse() Mouse {
	return new(EnMouse)
}

func (f *EnDeviceFactory) GetKeyboard() Keyboard {
	return new(EnKeyboard)
}

func (f *EnDeviceFactory) GetTouchpad() Touchpad {
	return new(EnTouchpad)
}

func GetFactoryByCountryCode(lang NameFactory) (DeviceFactory, error) {
	switch lang {
	case RU:
		return new(RuDeviceFactory), nil
	case EN:
		return new(EnDeviceFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported Country Code: %d", lang))
	}
}
