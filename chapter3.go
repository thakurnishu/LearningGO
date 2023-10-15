package main

import (
	"fmt"
	"math/rand"
)

func Chapter3() {
	fmt.Println("Chapter 3")
	blocks()
	ifelse()
	forLoops()
}

func blocks() {

	/*
		Blocks

		- Each place where declaration occurs is called block.
		-
	*/

	/*
		Shadowing Variables
	*/

	// here x := 10 is shadowed by x := 5 so you cannot access x := 10 in this block
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10

	y := 10
	if y > 5 {
		y, z := y, 30
		fmt.Println(y, z)
	}
	fmt.Println(y)

	/*
		Don't shadow package import
	*/
	// fmt := "oops"
	// fmt.Println(fmt)

	/*
		Universe Block

		- Universe Block which contains all other block
		- GO is small language with only 25 keywords, and built-in functions(make or close),
			types(like int anf string), constant(like true and false), and nil are not in this List
		- GO considers these predeclared identifers and defines them in unviverse block
		- As this names are delcared in universe block means they can be shadowed in other scope.
	*/
	fmt.Println(true)
	// true := 10
	// fmt.Println(true)  // 10
}

func ifelse() {

	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low[n]")
	} else if n > 5 {
		fmt.Println("That's too big[n]:", n)
	} else {
		fmt.Println("That's a good number[n]:", n)
	}

	/*
		Scoping variable to an if statement
	*/
	if m := rand.Intn(10); m == 0 {
		fmt.Println("That's too low[m]")
	} else if n > 5 {
		fmt.Println("That's too big[m]:", n)
	} else {
		fmt.Println("That's a good number[m]:", n)
	}
}

func forLoops() {
	/*
		In Golang their four ways to work with loops with 'for'
	*/

	// A complete, C-style 'for'
	for i := 0; i < 3; i++ {
		fmt.Println("hello", i)
	}

	// A Condition-only 'for'
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// An infinite 'for'
	j := 10
	for {
		fmt.Println("helloo")
		if j == 10 {
			break
		}
	}

	/*
		- for ignoring any of below (key, value) can use ' _ '
	*/

	// for-range statement
	evnValus := []int{2, 4, 6, 8, 10, 12, 14}
	for i, v := range evnValus {
		fmt.Printf("Index: %v => Value: %v\n", i, v)
	}

	// Itreating over maps
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i+1)
		for k, v := range m {
			fmt.Printf("Key: %v => value: %v\n", k, v)
		}
	}

	// Itreating over strings
	s := []string{"hello", "apple_π!"}
	for _, sample := range s {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

}
