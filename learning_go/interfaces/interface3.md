# Stringers
The fmt package (and many others) look for this interface to print values. Stringer is a universal interface. If a type implements Stringer, fmt.Println will call its String method automatically.

#### The stringer interface
```go
type Stringer interface {
    String() string
}
```

Basically,
```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```
gives output 
```
{Arthur Dent 42} {Zaphod Beeblebrox 9001}
```
while 
```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```
gives output
```
Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```
cause fmt and many packages look for this package only and by writing our own string method, we overrode the basic string method written in the package, simple.

# Errors
Similar to string, it is a built in interface : 
```go
type error interface {
    Error() string
}
```
Can be useful in future :
Example of overridding :
```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```
Output :
```
at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
```
