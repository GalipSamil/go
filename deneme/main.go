package main

import "fmt"

func main() {
	fruits := map[string]int{
		"elma":  5,
		"armut": 10,
	}

	for key, value := range fruits {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
