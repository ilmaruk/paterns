package main

import "fmt"
import "github.com/ilmaruk/patterns/observer"

func main() {
	fmt.Println("=== Observer Pattern ===")

	var subject observer.Subject = observer.NewConcreteSubject()
	subject.AddObserver(observer.NewColouredObserver("red"))
	subject.AddObserver(observer.NewColouredObserver("green"))
	subject.AddObserver(observer.NewColouredObserver("blue"))
	subject.RemoveObserver("green")
	subject.Notify("Hello, world!")
}
