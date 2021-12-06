package decorator

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponent struct{}

func (c ConcreteComponent) Operation() {
	fmt.Println("I am the concrete component.")
}
