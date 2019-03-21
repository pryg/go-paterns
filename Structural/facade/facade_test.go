package facade

import "testing"

func TestFacade(t *testing.T) {

	/* Это плохо. Много текста.
	power := &Power{}
	power.on()

	dvd := &DVDRoom{}
	dvd.load()

	hdd := &HDD{}
	hdd.copyFromDVD(dvd)
	*/

	// Если скомпоновать в класс Фасада будет короче.
	computer := NewComputer()
	computer.Copy()
}
