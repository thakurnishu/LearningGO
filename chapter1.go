package main

import "fmt"

// var (
// 	x    int
// 	y        = 20
// 	z    int = 123
// 	d, e     = 10, "hi"
// 	f, g string
// )

// // unUsed variable
// const thisConstant int = 134

func Chapter1() {
	// DATA TYPES

	var flag bool // no value assigned, set to false
	var isTrue = true

	var integer int // no value assigned, set to 0 (int64)
	var myint int64 = 234

	var thisIsString string // no value assigned, set to ""
	var myString string = "Nishant Singh"

	// Declaring
	var x int = 10
	var y = 10
	z := 10

	// Every declared local variable must be read

	fmt.Println(flag, isTrue, integer, myint, thisIsString, myString, x, y, z)
}
