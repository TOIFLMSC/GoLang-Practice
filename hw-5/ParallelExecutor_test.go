package parallelexecuter

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func DegreeofTwo(i int) error { // test function that counts degrees of twos
	a := math.Pow(2, float64(i))
	b := 2.0
	if i == 0 {
		b = 1.0
	} else if i > 0 {
		for j := 1; j < i; j++ {
			b = b * 2
		}
	} else {
		b = 0.5
		for j := 1; j < -(i); j++ {
			b = b * 0.5
		}
	}
	if a != b {
		return errors.New("math(degree): degree of number 2 is wrong") // returning special error with type error.Strings
	} else {
		return nil // if our calculations are right returning zero value for error
	}
}

func DegreeofThree(i int) error { // test function that counts degrees of threes
	a := math.Pow(3, float64(i))
	b := 3.0
	if i == 0 {
		b = 1.0
	} else if i > 0 {
		for j := 1; j < i; j++ {
			b = b * 3
		}
	} else {
		b = 1.0 / 3.0
		for j := 1; j < -(i); j++ {
			b = b * (1.0 / 3.0)
		}
	}
	if a != b {
		return errors.New("math(degree): degree of number 3 is wrong") //returning special error with type error.Strings
	} else {
		return nil // if our calculations are right returning zero value for error
	}
}

func DegreeofFour(i int) error { // test function that counts degrees of fours
	a := math.Pow(4, float64(i))
	b := 4.0
	if i == 0 {
		b = 1.0
	} else if i > 0 {
		for j := 1; j < i; j++ {
			b = b * 4
		}
	} else {
		b = 0.25
		for j := 1; j < -(i); j++ {
			b = b * 0.25
		}
	}
	if a != b {
		return errors.New("math(degree): degree of number 4 is wrong") //returning special error with type error.Strings
	} else {
		return nil // if our calculations are right returning zero value for error
	}
}

func Factorial(i int) error { // test function that counts factorials of entered numbers
	a := 1
	if i == 0 {
		a = 1
		return nil
	} else if i > 0 {
		for j := 1; j <= i; j++ {
			a = a * j
		}
		return nil // if our calculations are right returning zero value for error
	} else {
		return errors.New("math(factorial): factorial cannot be calculated from this number") //returning special error with type error.Strings
	}
}

func TestParallelExecutor(t *testing.T) { // creating testing function
	var a []func(int) error // creating slice of functions
	a = append(a, DegreeofTwo)
	a = append(a, DegreeofThree)
	a = append(a, DegreeofFour)
	a = append(a, Factorial)
	a = append(a, Factorial)
	a1 := ParallelExec(a, 5, 2, 2) //calling function ParallelExec and asigning its value to the variable a1
	for _, i := range a1 {         // in cycle printing errors if they exist (not equal zero value)
		fmt.Println(i)
	}
	var b []func(int) error // creating slice of functions
	b = append(b, DegreeofTwo)
	b = append(b, DegreeofThree)
	b = append(b, DegreeofFour)
	b = append(b, Factorial)
	b = append(b, Factorial)
	b = append(b, DegreeofTwo)
	b = append(b, DegreeofThree)
	b = append(b, DegreeofFour)
	b = append(b, Factorial)
	b = append(b, Factorial)
	b1 := ParallelExec(b, 10, 2, 2) //calling function ParallelExec and asigning its value to the variable b1
	for _, j := range b1 {          // in cycle printing errors if they exist (not equal zero value)
		fmt.Println(j)
	}
	var c []func(int) error // creating slice of functions
	c = append(c, DegreeofTwo)
	c = append(c, DegreeofThree)
	c = append(c, DegreeofFour)
	c = append(c, Factorial)
	c = append(c, Factorial)
	c = append(c, DegreeofTwo)
	c = append(c, DegreeofThree)
	c = append(c, DegreeofFour)
	c = append(c, Factorial)
	c = append(c, Factorial)
	c1 := ParallelExec(c, 10, 2, -2) //calling function ParallelExec and asigning its value to the variable c1
	for _, k := range c1 {           // in cycle printing errors if they exist (not equal zero value)
		fmt.Println(k)
	}
}
