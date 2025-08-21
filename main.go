package main

import (
	"fmt"

	"github.com/nelsonsaake/learn-prototype-design-pattern/1/shapes"
)

// Application

type Application struct {
	Shapes []shapes.Shape
}

func NewApplication() *Application {

	application := &Application{}

	c := shapes.NewCircle()
	c.X = 10
	c.Y = 10
	c.Radius = 20

	application.Shapes = append(application.Shapes, c)

	return application
}

func main() {

	fmt.Println()
}
