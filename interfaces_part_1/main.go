package main

import "fmt"

type Electronic struct {
	desi int
}

type Book struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}

func main() {
	/*book1 := &Book{desi: 10}
	book2 := &Book{desi: 20}
	fmt.Println(book1.ShippingCost())
	fmt.Println(book2.ShippingCost())*/

	books := []Book{{desi: 10}, {desi: 20}}
	fmt.Println(calculateTotalShippingCost(books))
}

func calculateTotalShippingCost(books []Book) int {
	total := 0

	for _, book := range books {
		total += book.ShippingCost()
	}
	return total
}
