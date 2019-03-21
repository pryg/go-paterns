package observed

type Observed interface {
	AddObserver(observer Observer)
	RemoteObserver(observer Observer)
	NotifyObservers()
}

type MeteoStation struct {
	temperature int
	pressure    int
	observers   []Observer
}

func (ms *MeteoStation) SetMeasurements(temperature, pressure int) {
	ms.temperature = temperature
	ms.pressure = pressure
	ms.NotifyObservers()
}

func (ms *MeteoStation) AddObserver(observer Observer) {
	ms.observers = append(ms.observers, observer)
}

func (ms *MeteoStation) RemoteObserver(observer Observer) {
	// Как-то удаляем
}

func (ms *MeteoStation) NotifyObservers() {
	for _, observer := range ms.observers {
		observer.HandleEvent(ms.temperature, ms.pressure)
	}
}

func NewMeteoStation() *MeteoStation {
	return &MeteoStation{}
}
