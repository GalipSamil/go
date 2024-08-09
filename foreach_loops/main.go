package main

import "fmt"

func main() {
	//var numbers = []int{1, 2, 3, 4}

	/*for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}*/

	/*for _, value := range numbers { // hata alıyorduk index i kullanmayıp yazdığımız için _ kullanarak hatayı çözdük.
		fmt.Println(value)
	}*/

	/*var language = "Golang"

	for _, character := range language {
		fmt.Printf("Character %c", character)
	}*/

	var names = map[string]int{
		"mustafa": 10,
		"murat":   20,
		"galip":   30,
	}

	for key, value := range names {
		fmt.Printf("Key: %s , Value: %d\n", key, value)
	}

}
