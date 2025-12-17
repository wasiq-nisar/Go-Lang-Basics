package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}

// Reciever Function
func (b *bill) format() string {		
	// Now this func will only be used for an objecy of bill type because of (b bill) added before func name
	formattedString := "Bill Breakdown: \n"

	var total float64 = 0
	for key, value := range b.items {
		formattedString += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}
	formattedString += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip) 

	// -25 will make the value 25 characters long and for "cake" it will use 4 and then fill remaining 21 with spaces on the right side, 25 without minus will add spaces on the left
	formattedString += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total)
	return formattedString
}

func (b *bill) updateTip(tip float64) {
	// if we take a pointer to a struct go will automatically dereference it without using *b
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b* bill) saveBill() {
	data := []byte(b.format())

	error := os.WriteFile("bills/"+b.name+".txt", data, 0644)	// 0644 -> Permission for the file
	if error != nil {
		panic(error)
	}
	fmt.Println("Bill was saved to file")
}