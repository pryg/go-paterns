package state

import (
	"fmt"
)

type Activity interface {
	DoSomething(human *Human)
}

// Contex
type Human struct {
	state Activity
}

func (h *Human) SetState(activity Activity) {
	h.state = activity
}

func (h *Human) DoSomething() {
	// передает контекст состоянию
	h.state.DoSomething(h)
}

// состояние работы
type Work struct{}

func (w *Work) DoSomething(contex *Human) {
	fmt.Println("Работаем!!!")
	contex.SetState(&WeekEnd{})
}

// состояние отдыха
type WeekEnd struct {
	count int
}

func (w *WeekEnd) DoSomething(context *Human) {
	// отдыхаем три дня
	if w.count < 3 {
		fmt.Println("Отдыхаем Zzz...")
		w.count++
	} else {
		context.SetState(new(Work))
	}
}
