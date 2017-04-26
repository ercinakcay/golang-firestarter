package main

import "fmt"

// Basic struct
type square struct {
	side float64
}

// Implementation of interface method.
func (z square) area() float64 {
	return z.side * z.side
}

// basic interface.
type shape interface {
	area() float64
}

// Implemented method running.
func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	// Initialization of square
	s := square{10}
	info(s)
}
