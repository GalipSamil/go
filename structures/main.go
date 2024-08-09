package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "Mustafa Murat Cosar", age: 30}

	customer1.ToString()
	changeName(&customer1)
	customer1.ToString()

	//var customer2 = Customer{id: 2, name: "Galip Samil Duger", age: 23, address: Address{city: "Ankara", district: "Kecioren"}}

	//customer1.age = 31

	//fmt.Println(customer1)
	//fmt.Println(customer2)

}

func changeName(customer *Customer) { // pointer ile yapmamız lazım sa by refrence
	customer.name = "Galip Samil"
}

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city     string
	district string
}

func (customer *Customer) ToString() {
	fmt.Println("Name: %s, Age: %d\n", customer.name, customer.age)
}
