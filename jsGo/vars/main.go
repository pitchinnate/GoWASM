package vars

import "github.com/google/uuid"

type Variable struct {
	Value     any
	Id        string
	Callbacks []func(any)
}

func Create(val any) Variable {
	//agent := GetAgent()
	id := uuid.New().String()
	//agent.Subscribe(id)
	return Variable{val, id, []func(any){}}
}

func (currentVar *Variable) Update(val any) {
	currentVar.Value = val
	for _, fn := range currentVar.Callbacks {
		fn(val)
	}
}

func (currentVar *Variable) OnUpdate(fn func(any)) {
	currentVar.Callbacks = append(currentVar.Callbacks, fn)
}
