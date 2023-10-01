package main

import "fmt"

func Chapter2() {
	/* COMPOSTE TYPES */
	array()
	slice()
}

func array() {

	// ARRAY (Rarely Used)
	/*
		- Go consider size of array to be part of it type
		- Can't use type conversion on array of diff size
	*/
	var arr [3]int // declaration
	var y = [4]int{1, 2, 3, 4}
	var z = [12]int{1, 5: 4, 6, 10: 24, 32} //  with  5: 23 --> will put 24 in 5 index of array

	fmt.Println("ARRAYS: ")
	fmt.Println(arr, y, z)

	var a = [...]int{1, 12, 23}
	var b = [3]int{1, 2, 3}
	fmt.Println(a == b) // false as content of array are not same

	// GO only have one-dimensinal array but we can simulate multidimension
	var multiArr [2][3]int
	fmt.Println(multiArr)          // [[0,0,0], [0,0,0]]
	fmt.Printf("%T ", multiArr[0]) // have type of [3]int
	fmt.Println(len(multiArr))     // length of 2
}

func slice() {

	// SLICES (Used Mostly)
	var slice = []int{1, 2, 3} // declartion
	fmt.Println("SLICES: ")
	fmt.Println(slice)

	var s = []int{1, 5: 4, 6, 10: 24, 32}
	fmt.Println(s)

	var multiSlice [][]int
	fmt.Println(multiSlice) // []

	var sl []int
	fmt.Println(sl == nil) // true
	/*
		We can't use == with two slice only thig we can compare is it with 'nil'
	*/

	// Build-in functions
	// --> append (used to grow slice)
	/*
		- GO is call by value language
		- by calling append function and
			passing slice it make copy of
			it and a add value in it
		- so we have to assign returned variable
	*/
	var mySliceToAppend []int
	mySliceToAppend = append(mySliceToAppend, 10) // append(`slice of any type`, `value of that slice`)

	mySliceToAppend = append(mySliceToAppend, 12, 24, 354)

	y := []int{1, 43, 23}
	mySliceToAppend = append(mySliceToAppend, y...) // to append elements in different slice in another slice use `...` operator
	fmt.Println(mySliceToAppend, y)

	// --> capacity
	/*
		- Every Slice have capacity, which is reserved consecutive memory location
		- Every time size of array greater than capacity of array,
			append function use GO runtime to allocate new slice with
			greater capacity
		--> GO Runtime
			- It provide services like memory allocation, garbage collection
				concurrency support, networking and implementations of
				built-in function and types
			- GO runtime is compiled into ervery Go binary,
				make distribution easier and aviods issues like
				compatibility
	 F*/

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	// --> make
	/*
		- To initialize size and capacity that we want use make() func
		-
	*/
	useMakeLen := make([]int, 5) // make(Type, size) --> initialize slice with length of 5 [0,0,0,0,0] with capacity of 5
	fmt.Println(useMakeLen, len(useMakeLen), cap(useMakeLen))
	useMakeLen = append(useMakeLen, 10)
	fmt.Println(useMakeLen, len(useMakeLen), cap(useMakeLen)) // [0,0,0,0,0,10]

	useMakeLenCap := make([]int, 5, 10) // size = 5, capacity = 10
	fmt.Println(useMakeLenCap, len(useMakeLenCap), cap(useMakeLenCap))

	useMake := make([]int, 0, 10) // size = 0, capacity = 10
	useMake = append(useMake, 4, 5, 6, 7)
	fmt.Println(useMake, len(useMake), cap(useMake)) // size = 4, capacity = 10

	var s2 []int
	fmt.Println(s2 == nil) // true

	var s1 = []int{}       // usefull when converting slice in JSON using 'encoding/json'
	fmt.Println(s1 == nil) // false

	fmt.Println(s1, s2)

}
