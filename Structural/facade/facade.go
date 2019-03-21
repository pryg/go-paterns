// Для сокрытия объемных действий
package facade

import (
	"fmt"
)

type Power struct{}

func (p *Power) on() {
	fmt.Println("Включение питания")
}
func (p *Power) off() {
	fmt.Println("Выключение питания")
}

type DVDRoom struct {
	data bool
}

func (d *DVDRoom) hasData() bool {
	return d.data
}
func (d *DVDRoom) load() {
	d.data = true
}
func (d *DVDRoom) unload() {
	d.data = false
}

type HDD struct{}

func (h *HDD) copyFromDVD(dvd *DVDRoom) {
	if dvd.hasData() {
		fmt.Println("Копируются данные из диска")
	} else {
		fmt.Println("Вставьте диск с ланными")
	}
}

// Класс фасада
type Computer struct {
	power *Power
	dvd   *DVDRoom
	hdd   *HDD
}

func NewComputer() *Computer {
	computer := new(Computer)
	computer.power = new(Power)
	computer.dvd = new(DVDRoom)
	computer.hdd = new(HDD)
	return computer
}

func (c *Computer) Copy() {
	c.power.on()
	c.dvd.load()
	c.hdd.copyFromDVD(c.dvd)
}
