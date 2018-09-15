package main

import (
	"fmt"
)

// Apple sample struct
type Apple struct {
	Quantity int
}

// Incr increment counter
func (a *Apple) Incr(n int) {
	a.Quantity += n
}

// Decr decrement counter
func (a *Apple) Decr(n int) {
	a.Quantity -= n
}

func (a *Apple) String() string {
	return fmt.Sprintf("%v", a.Quantity)
}

func main() {
	apple := &Apple{}
	apple.Incr(10)
	apple.Decr(5)
	fmt.Println(apple)

}
