## Interface values
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
(value, type)

An interface value holds a value of a specific underlying concrete type.
Calling a method on an interface value executes the method of the same name on its underlying type.

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## Interface values with nil underlying values

I is interface and T is a struct type
```
var i I
var t *T
i = t
//assigns a null value to interface i , try printing %v and %T of i
```
and 
```
var i I
i = &T{"Hello"}
```
are same
```
// %V and %T of first case
(<nil>, *main.T)
<nil>

// %V and %T of second case
(&{hello}, *main.T)
hello
```

## Nil interface values
Above was assigning a nil value to interface, now we will assign a nil interface to a variable.
```go 
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	//i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```
Output : 
```
(<nil>, <nil>)
```
A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error 
Why ? : because above had a type and the value was nil, here the type is nil, so which method to call?
