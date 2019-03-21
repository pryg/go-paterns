package state

import "testing"

func TestStateLite(t *testing.T) {
	station := new(Radio7)
	radio := new(Radio)
	radio.SetStation(station)

	// эмитация смены состояния
	for i := 0; i < 10; i++ {
		radio.Play()
		radio.NextStation()
	}
}

func TestState(t *testing.T) {
	human := new(Human)
	human.SetState(new(Work))

	// эмитация смены состояния
	for i := 0; i < 10; i++ {
		human.DoSomething()
	}
}
