// Golang program illustrates how
// to implement an interface
package main

import "fmt"

// Creating an interface
type tank interface {

    // Methods
    Tarea() float64
    Volume() float64
}

type myvalue1 struct {
	radius float64
	height float64
}

type myvalue2 struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m myvalue1) Tarea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m myvalue1) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

func (m myvalue2) Tarea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m myvalue2) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

// Main Method
func main() {

	// Accessing elements of
	// the tank interface
	var t tank
	t = myvalue1{10, 14}
	t = myvalue2{10, 14}
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())
}
