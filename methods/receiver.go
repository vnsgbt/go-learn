package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func Scale(v *Vertex, f float64)  {
	v.X *= f
	v.Y *= f
}
func (v *Vertex) ScaleFunc(f float64)  {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	f := MyFloat(-7)

	fmt.Println(v.Abs())
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v)

	Scale(&v, 10)
	fmt.Println(v)

	z := Vertex{3,4}
	z.ScaleFunc(2)
	Scale(&z,10)

	p := &Vertex{4,3}
	p.ScaleFunc(3)
	Scale(p,8)

	fmt.Println(z, p)
}
