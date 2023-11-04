```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
    Scale()
}

func main() {

    //Initialize a variable of type Abser
	var a Abser


	f := MyFloat(-math.Sqrt2)
	a = f  // a MyFloat implements Abser
    fmt.Println(a.Abs())


	v := Vertex{3, 4}
	a = &v // a *Vertex implements Abser
    fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v
    
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Scale() {
    f = f * 10
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale() {
	v.X = v.X * 10
    v.Y = v.Y * 10
}

```

With `a = &v // a *Vertex implements Abser`, it is not important for function `Abs()` to be defined with a pointer receiver. The interface implementation is implicit. The interface is implemented by the pointer receiver, so the interface is implemented by the value that the pointer points to.

But while `a = v`, it is important for function `Abs()` to be defined without a pointer receiver. 

func (v *Vertex) Abs() and func (v *Vertex) Scale() works with `a = &v` but not with `a = v` <br>
and func (f MyFloat) Abs() works with `a = f` as well as with `a = &f`.

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
    Scale()
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	a = &v // a *Vertex implements Abser
    // Basically the pointer implements the Abser interface by the below methods.
	a.Scale()
    fmt.Println(a.Abs()) 
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale() {
	v.X = v.X * 10
    v.Y = v.Y * 10
}
``` 
will result in 
1.4142135623730951
50<br>
while changing `func (v *Vertex) Scale()` to `func (v *ertex) Scale()` will result in <br>
1.4142135623730951
5
<br><br>
Calling the Scale method on a, which is a method of Vertex with a value receiver. This is allowed because Go can automatically take the address of the value (v) and call the method with a pointer receiver (func (v *Vertex) Scale()) but not vice-versa.