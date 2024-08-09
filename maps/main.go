package main

import "fmt"

func main() {
	/*var names map[string]int
	names = make(map[string]int, 0)

	names["Galip"] = 1
	names["Şamil"] = 2
	names["Duger"] = 3

	//fmt.Println(names)

	// fmt.Println((names["Şamil"]))

	fmt.Println(names["Duygu"]) // melike yok 0 döndü*/
	names := map[string]int{
		"Mustafa": 1,
		"Murat":   2,
		"Yağci":   3,
	}

	delete(names, "Murat")
	fmt.Println(names)
}
