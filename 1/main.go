package main

import (
	"fmt"

	"github.com/nelsonsaake/learn-prototype-design-pattern/1/shapes"
)

// Application

type Application struct {
	Shapes []shapes.Shape
}

func (a *Application) addShape(shape shapes.Shape) {
	a.Shapes = append(a.Shapes, shape)
}

func (a *Application) businessLogic() {

	// Prototype rocks because it lets you produce a copy of
	// an object without knowing anything about it's type
	shapesCopy := []shapes.Shape{}

	for _, shape := range a.Shapes {
		shapesCopy = append(shapesCopy, shape.Clone())
	}
}

func NewApplication() *Application {

	application := &Application{}

	circle := shapes.NewCircle()
	circle.X = 10
	circle.Y = 10
	circle.Radius = 20
	application.addShape(circle)

	anotherCircle := circle.Clone()
	application.addShape(anotherCircle)

	rectangle := shapes.NewRectangle()
	rectangle.Width = 10
	rectangle.Height = 20
	application.addShape(rectangle)

	return application
}

func main() {

	fmt.Println()
}
