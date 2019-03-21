package state

import (
	"fmt"
	"reflect"
)

// State
type Station interface {
	Play()
}

// State 1
type Radio7 struct{}

func (s *Radio7) Play() {
	fmt.Println("Радио 7...")
}

// State 2
type RadioDFM struct{}

func (s *RadioDFM) Play() {
	fmt.Println("Радио DFM...")
}

// State 3
type VestiFM struct{}

func (s *VestiFM) Play() {
	fmt.Println("Радио Вести ФМ..")
}

// Contex set its state
type Radio struct {
	station Station
}

func (r *Radio) SetStation(station Station) {
	r.station = station
}

// переключает станции - состояния, т.е. радио меняет состояние
func (r *Radio) NextStation() {
	if IsInstanceOf(r.station, (*Radio7)(nil)) {
		r.SetStation(new(RadioDFM))
	} else if IsInstanceOf(r.station, (*RadioDFM)(nil)) {
		r.SetStation(new(VestiFM))
	} else if IsInstanceOf(r.station, (*VestiFM)(nil)) {
		r.SetStation(new(Radio7))
	}
}

func (r *Radio) Play() {
	r.station.Play()
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	//	fmt.Printf("TypeOf(objectPtr) = %v, TypeOf(typePtr) = %v\n", reflect.TypeOf(objectPtr), reflect.TypeOf(typePtr))
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}
