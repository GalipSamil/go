package main

import "fmt"

func findMax(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	nums := []int{1, 3, 45, 534, 5, 45, 5345, 534, 544234}
	fmt.Println("The biggest number in the array:", findMax(nums))

	var arr = []int{1, 43, 54, 54, 65, 620, 540}
	fmt.Println("Value at index 2:", arr[2])

	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println("The biggest number in the second array:", max)

	fruists := map[string]int{
		"karpuz": 6,
		"Kiraz":  7,
	}

	fruists["portakal"] = 8

	fmt.Println("Updated fruits", fruists)

	customer := []string{"galip", "samil", "duger"}

	fmt.Println(customer)

	maps := map[string]int{
		"elma":   7,
		"kayisi": 3,
	}

	fmt.Println(maps)

	var x int = 10
	var p *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Addrtess of x:", p)
	fmt.Println("Value at address p:", *p)

	*p = 20

	fmt.Println("Updated value of x:", x)

	age := 22
	pAge := &age
	fmt.Println("Address of age", pAge) // adress
	fmt.Println("Age(from pointer)", *pAge)
	*pAge = 41
	fmt.Println("Age=", age)                // 41
	fmt.Println("Age(from pointer)", *pAge) // 41

	point := 1000
	pPoint := &point
	*pPoint = *pPoint * 2
	fmt.Println("point=", point) // 500

}
