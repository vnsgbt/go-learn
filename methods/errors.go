package main

import (
	"time"
	"fmt"
	"image"
	"golang.org/x/tour/pic"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	n := Image{}
	pic.ShowImage(n)

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type Image struct {}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}