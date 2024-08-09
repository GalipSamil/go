package main

import "fmt"

func main() {
	/*
	   age := 18
	   if age >= 18 {
	       fmt.Println("You can vote!")
	   } else {
	       fmt.Println("You cannot vote!")
	   }
	*/

	a := 40
	b := 20
	c := 30

	if a > b && a > c {
		fmt.Println("a is the biggest number")
	} else if b > a && b > c {
		fmt.Println("Greatest variable is b!")
	} else {
		fmt.Println("Greatest variable is c!")
	}
}
