package main // -> If package is named main then it will tell the compiler to complie it into a exec program

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myBill := createBill()
	promptOptions(myBill)

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new Bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new Bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created a new Bill with name: ", name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a -add item, s - save bill, t - add tip): ", reader)
	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		priceString, _ := getInput("Item price: ", reader)

		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, price)

		fmt.Println("Item addded - ", name, price)
		promptOptions(b)
	case "t":
		tipString, _ := getInput("Enter tip amount($): ", reader)
		tip, err := strconv.ParseFloat(tipString, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(tip)

		fmt.Println("Tip addded - ", tip)
		promptOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("you saved the bill", b.name)
	default:
		fmt.Println("That was not a valid option... ")
		promptOptions(b)
	}
}

func getInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}