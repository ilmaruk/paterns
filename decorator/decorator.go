package decorator

import "fmt"

type DecoratedComponent struct {
	decorated Component
}

func NewDecoratedComponent(c Component) DecoratedComponent {
	return DecoratedComponent{
		decorated: c,
	}
}

func (d DecoratedComponent) Operation() {
	fmt.Println("Before decorated operator.")
	d.decorated.Operation()
	fmt.Println("After decorated operator.")
}
