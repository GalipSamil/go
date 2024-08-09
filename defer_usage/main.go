package main

import "fmt"

func main() {
	defer fmt.Println("Hello") // erteleme yaptı
	fmt.Println("Selam")
	fmt.Println("Galip")
	defer fmt.Println("james")   // LIFO
	defer fmt.Println("Cameron") //LIFO  last In First OUT

	var conditions = true
	defer fmt.Println("first Book")
	if conditions {
		fmt.Println("Zeynep")
	}
	defer fmt.Println("Second. Defer") // main func bittiği için çalıştırmadı bunu

	if conditions {
		panic("An error accurred ") // uygulama panic atsa da defer çalışıyor pek anlamadım ama öyleymiş

	}

	defer fmt.Println("Selam Kizlar")
	fmt.Println("Selam galip")

}
