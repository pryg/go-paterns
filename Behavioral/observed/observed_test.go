package observed

import "testing"

func TestObserved(t *testing.T) {
	meteoStation := NewMeteoStation()
	meteoStation.AddObserver(new(ConsolObserver))

	// при изменениях оповещаются все слущатели
	meteoStation.SetMeasurements(30, 769)
}
