package main

import (
	"fmt"
)

func Chapter2() {

	fmt.Println("Chapter 2")

	/* COMPOSTE TYPES */
	array()
	slice()
	stringRunesBytes()
	maps()
	structs()
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

	var s1 []int
	fmt.Println(s1 == nil) // true

	var s2 = []int{}       // usefull when converting slice in JSON using 'encoding/json'
	fmt.Println(s2 == nil) // false

	fmt.Println(s1, s2)

	var s3 = []int{1, 2, 3}
	fmt.Println(s3)

	s4 := make([]int, 5)

	for k := range s4 {
		s4[k] = k
	}
	for k, v := range s4 {
		println(k, v)
	}
	slicingSlice()
}

func slicingSlice() {
	a0 := []int{1, 2, 3, 4, 5, 6}

	a1 := a0[:2]
	a2 := a0[1:]
	a3 := a0[1:4]
	fmt.Printf("original : %v\n", a0)
	fmt.Printf("x[:2] : %v\n", a1)
	fmt.Printf("x[1:] : %v\n", a2)
	fmt.Printf("x[1:4] : %v\n", a3)

	/*
		- When you take slice from slice, you are not making
			copy of the data. Instead, we have now two variable that are sharing memory
	*/
	s0 := []int{1, 2, 3, 4, 5}
	s1 := s0[:3] // [1, 2, 3]
	s2 := s0[1:] // [2, 3, 4, 5]

	s1[1] = 15      // 2 --> 15
	s2[3] = 13      // 5 --> 13
	s0[0] = 20      // 1 --> 20
	fmt.Println(s0) // [20 15 3 4 13]

	/*
		- It get more confusing while using append function
		- while slicing slice into subslice,
			subslice's capacity = original's capaicty - offset of subslice within the original slice { subSlice[offset:] }
	*/
	x0 := []int{1, 2, 3, 4, 5}
	x1 := x0[:2]
	x2 := x0[1:]
	fmt.Println(cap(x0), cap(x1), cap(x2))
	x1 = append(x1, 30)

	fmt.Printf("x0 : %v\nx1 : %v\nx2 : %v\n", x0, x1, x2)

	// Full slice Expression
	/*
		- subSlice[offset:length:?]
		- here ? -->
			Last position in parent slice's capacity that's available for the subslice,
			subtract the starting offset from this number to get the subslice's capacity
	*/
	y0 := []int{1, 2, 3, 4, 5}
	y1 := y0[:2:2]
	y2 := y0[3:4:5]
	fmt.Println(cap(y0), cap(y1), cap(y2)) // 5 2 5-3
	y1 = append(y1, 30)
	y2 = append(y2, 60, 02)

	fmt.Printf("y0 : %v\ny1 : %v\ny2 : %v\n", y0, y1, y2)

	/*
		CONVERTING Arrays to Slices
		- taking slice from array have same memory sharing properties as above
	*/
	z0 := [4]int{5, 6, 7, 8}
	z1 := z0[:2]
	z2 := z0[2:]
	z0[0] = 10
	fmt.Printf("z0 : %v\nz1 : %v\nz2 : %v\n", z0, z1, z2)

	/*
		Built-in Function [COPY]
		- create slice that's independent of the original

	*/

	b0 := []int{1, 2, 3, 4}
	b1 := make([]int, 4)
	num := copy(b1, b0)
	fmt.Println("Total element copied : ", num)
	fmt.Printf("b0 : %v\nb1: %v\n", b0, b1)

	/*
		- take two parameter copy(des, src) and return number of copied element
		- which is limited by length
	*/
	b2 := make([]int, 2)
	num = copy(b2, b0[2:])
	fmt.Println("Total element copied : ", num) // 2
	fmt.Printf("b0 : %v\nb2: %v\n", b0, b2)

	// can be used on array
	b3 := [4]int{5, 4, 6, 7}
	fmt.Printf("b3 : %v\n", b3)
	copy(b2, b3[:])
	copy(b3[:], b0)
	fmt.Printf("b2 : %v\nb3 : %v\n", b2, b3)
}

func stringRunesBytes() {
	/*
		- rune is alias for int32
		- byte is alias for unit8

		In GO
		- String is madeup of sequence of bytes
	*/
	var s string = "Hello 😊" // as it takes four bytes to represent the smiling face emoji in UTF-8
	fmt.Println(len(s))      // len(s) = 10 not 7

	/*
		- Single rune and byte can be converted to string
		- Only Byte and Rune are allow for type conversion to string
	*/
	var a0 rune = 'x'
	var a1 string = string(a0)

	var a2 byte = 'y'
	var a3 string = string(a2)

	var a4 int = 65
	var a5 string = string(byte(a4))

	fmt.Println(a1, a3, a5) // x y A

	/*
		- String can be converted into slice of bytes and rune
		- Most data in go is read nad written as sequence of bytes
		- So, Most of time we will use Bytes as type converstion for string
	*/
	var s0 string = "Hello 😊"
	var s1 []byte = []byte(s0)
	var s2 []rune = []rune(s0)
	fmt.Println(s1) // string converted into UTF-s Bytes [72 101 108 108 111 32 240 159 152 138]
	fmt.Println(s2) // string converted into rune [72 101 108 108 111 32 128522]

}

func maps() {
	/*
		MAPS

		- can grow automatically as you add key-value pairs to them
		- if knew what size is required can create it with map
		- Maps are not comparable
			- can check if they are equal to nil
			- cannot use operatot like == or !=

		- Key for map can be of any type except slice or map
	*/

	/*
		- This is declartion of map will give nil map with length of zero
		- attempting to write to nil map variable causes a panic
	*/
	var nilMap map[string]int
	fmt.Println("Nil map length : ", len(nilMap))

	/*
		- we are using empty map literal
		- Will give non nil map with length of zero
		- but we can read and write into this map
	*/
	totalWins := map[string]int{}
	totalWins["Nishu"] = 2
	totalWins["Viku"] = 3
	totalWins["NoOne"]++
	fmt.Println(totalWins)

	/*
		- nonempty map literal
	*/

	teams := map[string][]string{
		"Oracs":   {"Fred", "Ralph", "BIjou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	/*
		- If you know how many key-value pairs are required in map
		- Created with this method length will still be zero
		- they can grow past the initially specified
	*/
	ages := make(map[int]string, 10)
	fmt.Println(len(ages)) // 0

	/*
		Comma Ok Idiom
	*/
	m := map[string]int{
		"nishu": 5,
		"viku":  6,
	}

	v, ok := m["nishu"] // v -> value of key, ok -> bool tell if exist
	fmt.Println(v, ok)

	v, ok = m["viku"]
	fmt.Println(v, ok)

	v, ok = m["someone"]
	fmt.Println(v, ok)

	/*
		Deleteing from maps
		- Using built-in function 'delete'
	*/
	n := map[string]int{
		"nishu": 5,
		"viku":  6,
	}
	fmt.Println(n)
	delete(n, "nishu")
	fmt.Println(n)

	/*
		Using Maps as Sets

		- Use the key of map for element of set (as here int is type use as key)
		- Use bool for value (for check if element is not double )
	*/

	intSet1 := map[int]bool{}                         // after for loop length = 8
	vals1 := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10} // total length = 11
	for _, v := range vals1 {
		intSet1[v] = true
	}
	fmt.Println(len(vals1), len(intSet1)) // 11 8
	fmt.Println(intSet1[5])
	fmt.Println(intSet1[500])
	if intSet1[100] {
		fmt.Println("100 is in the set ")
	}

	/*
		- Some People prefer to use struct{} for value when map is being is
			used to implement a set
		- As empty struct{} use zero bytes, where as bool use 1 byte
	*/
	intSet2 := map[int]struct{}{}                     // after for loop length = 8
	vals2 := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10} // total length = 11
	for _, v := range vals2 {
		intSet2[v] = struct{}{}
	}
	if _, ok := intSet2[5]; ok {
		fmt.Println("5 is in the set")
	}
	fmt.Println(intSet2)

	/*
		- Unless you have very large data set effect on memory usage is very low.
	*/

}

func structs() {
	/*
		STRUCTS

		- when you have realted data than you want to group together, use struct{}
	*/

	type person struct {
		name string
		age  int
		pet  string
	}

	// Defining Variable

	var fred person // All value inside struct get to their zero value like int -> 0 etc
	fmt.Println(fred)

	bob := person{}
	fmt.Println(bob)

	/*
		- declaring with comma-separated list
		- values should be assigned in correct order
	*/
	viku := person{
		"Vikrant",
		40,
		"Dog",
	}
	fmt.Println(viku)

	/*
		- Declaring with map literal style
		- No need to declare it in specific order
		- Useful as if your initial struct have new value add to it
			code will not get compromised
	*/
	nishu := person{
		age:  30,
		name: "Nishant",
	}
	fmt.Println(nishu)

	// to access field of struct
	nishu.name = "Nishant Singh"
	fmt.Println(nishu)

	/*
		Anonymous Structs

		- these struct are associated with single instance
		- These Struct have two common use cases:
			- Translating external data into Struct and visa versa (Marshaling and UnMarshaling)
			- Writing tests
	*/

	// 1
	var car struct {
		company string
		model   string
		year    uint
	}

	car.company = "Maruti"
	car.model = "Auto K10"
	car.year = 2015
	fmt.Println(car)

	// 2
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)

	/*
		Comparing and Converting Structs

		- Structs that are composed of comparable types are comparable
		- Structs that are composed of Slices, Maps, function and channel are not comparable
	*/

	type firstPerson struct {
		name string
		age  int
	}
	myfirstPerson := firstPerson{
		name: "Nishu",
		age:  20,
	}
	fmt.Println(myfirstPerson)

	/*
		- As you can see we can have type conversion between myfirstPerson and mysecondPerson
		- But we can not compare these two instances as they are different types
	*/
	type secondPerson struct {
		name string
		age  int
	}
	mysecondPerson := secondPerson(myfirstPerson)
	fmt.Println(mysecondPerson)

	/*
		- We cannot have type conversion from myfirstPerson to mythirdPerson
		- As order of field is not same
	*/
	type thirdPerson struct {
		age  int
		name string
	}

	/*
		- we cannot have type conversion from myfirstPerson to myfourthPerson
		- As field names don't match
	*/
	type fourthPerson struct {
		firstName string
		age       int
	}

	/*
		- we cannot have type conversion from myfirstPerson to myfifthPerson
		- As additional field is in this struct
	*/
	type fifthPerson struct {
		age         int
		name        string
		luckyNumber int
	}

	/*
		- if you have two struct and one of them is Anonymous Struct then you
			can compare other one with Anonymous struct without type conversion
			if fields of both structs have same names, order, and types.
	*/

	f := firstPerson{
		name: "Nishu",
		age:  20,
	}

	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g)

}
