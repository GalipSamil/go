package main

import "fmt"

func main() {

	//a = 10
	/*var a = 10
	var b int
	var p *int

	b = a

	p = &a

	/*	fmt.Println(a)
		fmt.Println(p) // adres değeri  oxc00000
		fmt.Println(*p)*/ // adres değerine gittik ve oradaki değeri aldıl 10

	//*p = 20

	//fmt.Println(a, b)

	var a int = 10
	fmt.Println(a)

	//add12(a)
	add12pointer(&a)
	fmt.Println(a)

}

func add12(x int) { // pass by value
	x = x + 12
}

func add12pointer(x *int) { // pass by refeerence
	*x = *x + 12
}
