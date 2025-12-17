package main

// Everything declared in the same package can be used in any file having the smae package

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("hello", n)
}

func showScore() {
	fmt.Println("You scored this many points:", score)
}