package observer

import "fmt"

type Observer interface {
	Id() string
	Update(msg string)
}

type ColouredObserver struct {
	colour string
}

func NewColouredObserver(colour string) ColouredObserver {
	return ColouredObserver{
		colour: colour,
	}
}

func (o ColouredObserver) Id() string {
	return o.colour
}

func (o ColouredObserver) Update(msg string) {
	fmt.Printf("[%s] %s\n", o.colour, msg)
}
