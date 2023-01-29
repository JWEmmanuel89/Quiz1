// CMPS2242 - Quiz 1

package main

import (
	"fmt"
	"math"
)

// #1. Create struct type Triangle
type Triangle struct {
	Base   float64
	Height float64
}

// #2. Create method type Area
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// #3. Create method type Perimeter
func (t Triangle) Perimeter() float64 {
	sideC := math.Sqrt((t.Base * t.Base) + (t.Height * t.Height))
	return t.Base + t.Height + sideC
}

func main() {
	// #4. Create variable with value for type Triangle
	tri1 := Triangle{
		Base:   2,
		Height: 5,
	}

	//#5. Call and test Area and Perimeter of type Triangle
	fmt.Println(tri1.Area())
	fmt.Println(tri1.Perimeter())
}
