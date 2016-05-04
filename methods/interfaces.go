package main

import (
	"math"
	"fmt"
)

type Abser interface { Abs() float64 }
type MyFloat float64
type Vertex struct { X, Y float64 }

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


type I interface {
	M()
	N()
}
type T struct { S string }
type F float64

func (t *T) M()  { fmt.Println(t.S) }
func (f F) M() { fmt.Println(f) }

func (t *T) N()  {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) N()  { }

func describe(i I)  { fmt.Printf("(%v, %T)", i, i) }

func describeJ(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	var a Abser

	f := MyFloat(-math.Sqrt2)
	v:= Vertex{3,4}

	a = f
	a = &v

//	a = v

	fmt.Println(a.Abs())

	var i I
	i = &T{"hello"}
	i.M()
  describe(i)

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	fmt.Println()
	i.N()

	i = &T{"hello"}
	describe(i)
	i.N()

	var j interface{}
	describeJ(j)

	j = 42
	describeJ(j)

	j = "hello"
	describeJ(j)
}


