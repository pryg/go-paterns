package prototype

type Copyable interface {
	Copy() interface{}
}

type Human struct {
	age  int
	name string
}

func NewHuman(age int, name string) *Human {
	return &Human{age: age, name: name}
}

func (h *Human) GetAge() int { return h.age }

func (h *Human) GetName() string { return h.name }

func (h *Human) Copy() interface{} {
	return NewHuman(h.age, h.name)
}

type HumanFactory struct {
	human *Human
}

func CreateHumanFactory() *HumanFactory {
	return &HumanFactory{}
}

func (hf *HumanFactory) SetPrototype(human *Human) {
	hf.human = human
}

func (hf *HumanFactory) MakeCopy() *Human {
	return hf.human.Copy().(*Human)
}
