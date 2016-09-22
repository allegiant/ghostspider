package core

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) SetName(name string) {
	p.Name = name
}

func (p Person) Out() {
	fmt.Println(p.Name)
}
