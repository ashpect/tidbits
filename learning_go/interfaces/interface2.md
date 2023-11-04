## The empty interface 
The empty interface (and also type assertions) <br>
THIS IS NICE!!

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

    //Type assertion
	s := i.(string)
	fmt.Println(s)

    // s will get value and ok will be a bool true
	s, ok := i.(string)
	fmt.Println(s, ok)

    //for testing type assertion, f will get 0 and ok will be a bool false as i holds string
	f, ok := i.(float64)
	fmt.Println(f, ok)

    // direct assignment will result in panic
	f = i.(float64) // panic
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```
Output :

```
(<nil>, <nil>)
(42, int)
(hello, string)
hello
```
An empty interface may hold values of any type.  For example, fmt.Print takes any number of arguments of type interface{}.

## Type Switches
Self-explanatory

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```