package main

import (
	"fmt"
	"github.com/ilmaruk/patterns/decorator"
)

func main() {
	fmt.Println("=== DecoratedComponent Pattern ===")

	var component decorator.ConcreteComponent = decorator.ConcreteComponent{}
	test(component)

	fmt.Println("===")

	var decoratedComponent decorator.DecoratedComponent = decorator.NewDecoratedComponent(component)
	test(decoratedComponent)
}

func test(c decorator.Component) {
	c.Operation()
}
