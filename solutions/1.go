package solutions

import (
	"fmt"
	"time"
)

type Human struct {
	name string
	age  int
}

func newHuman(name string, age int) Human {
	return Human{
		name: name,
		age:  age,
	}
}

type Action struct {
	Human
	timeAction time.Time
}

func newAction(name string, age int, t time.Time) Action {
	return Action{
		Human{
			name,
			age,
		},
		t,
	}
}

func (h Human) live() {
	fmt.Printf("%s (%d y.o.) lives right now\n", h.name, h.age)
}

func (h Action) walk() {
	fmt.Printf("%s (%d y.o.) walked at %d o'clock\n", h.name, h.age, h.timeAction.Hour())
}

func Example1() {
	human := newHuman("Oleg", 25)
	// Human может жить
	human.live()
	action := newAction(human.name, human.age, time.Now())
	fmt.Println("Again:")
	// Acrion может и жить, и гулять
	action.live()
	action.walk()
}
