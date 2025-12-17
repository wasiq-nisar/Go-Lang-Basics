package main

import (
	"fmt"
	"strings"
)
var score = 99.5

// main func is the entry point. Can only have one main
func basics() {
	fmt.Println("Hello World")

	// Strings

	// If we decalre a variable and not use it GO consider it a error
	var nameOne string = "mario"
	var nameTwo = "luigi"	// Automatically infers to be string
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"	// Short Hand to initialize a variable and can't be used outside the function
	fmt.Println(nameFour)

	// Int
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits and Memory
	var numOne int8 = 24 // 8-Bit number can not be more then 127
	var numTwo int8 = -128
	var numThree uint = 255
	fmt.Println(numOne, numTwo, numThree)

	// Float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1231234.12351235123512352
	scoreThree := 123.41234		// By default considered float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	age := 35
	name := "shaun"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted string), %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v = value in default format
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q = quotes
	fmt.Printf("age is of type %T \n", age)                    // %T is the type
	fmt.Printf("you scored %f points! \n", 225.55)             // %f = float format
	fmt.Printf("you scored %0.1f points! \n", 225.55)          // %0.2f = float with 2 decimal points

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

	// Arrays
	
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
	
	// Functions
	fn1, sn1 := getInitials("tifa lockheart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("baret")
	fmt.Println(fn2, sn2)

	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}

	// Maps
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[647583927] = "yoshi"
	fmt.Println(phonebook)

	// Pointers
	name2 := "tifa"

	// updateName(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name2)

	// * gets the value at the specified memory address(Derefferencing)
	m := &name2
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(m) // using pointer as arg, can pass &name directly instead of "m" as well
	fmt.Println(name)

	// Structs
	// myBill := newBill("mario's Bill")
	// myBill.updateTip(10)
	// myBill.addItem("onion soup", 4.50)
	// myBill.addItem("pickle soup", 9.50)
	// fmt.Println(myBill.format())
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string 
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""
}

// func updateName(n string) {
// 	n = "wedge"
// }

func updateName(n *string) {
	*n = "wedge"
}