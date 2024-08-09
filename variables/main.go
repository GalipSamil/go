package main

import (
	"fmt"
)

func main() {
	/*var productName string
	var quantity int
	var discount float32
	var IsInStock bool

	productName = "Golang Dersleri"
	quantity = 11
	discount = 0.35
	IsInStock = true

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(IsInStock, reflect.TypeOf(IsInStock))*/

	/*var productName string = "Goland Dersleri"
	fmt.Println(productName, reflect.TypeOf(productName))

	var productName1 = "Goland Dersleri 1" // string diy ebelirtmemize gerek yok
	fmt.Println(productName, reflect.TypeOf(productName1))

	productName2 := "Goland Dersleri2" // := ifadesi ile değişken tanımlayabiliyoruz. var a gerek kalmıyor.
	fmt.Println(productName, reflect.TypeOf(productName2))*/

	/*quantity := 10 // değişkenin ne olmasını istdiğimiz durumada bu şekilde tanımlama yapmak mantıksız çünkü mesela int64 ile tanımlamak istiyoruz ama böyle direk int olarak tanımlıyor.
	fmt.Println(quantity, reflect.TypeOf(quantity))

	var quantity1 int64 = 12 // := li tanımlama bu int64 ü kullanmak mümkün değil eğer int64 ,int16,int32 kullanacaksak var int... ile tanımlama yapmamız lazım.
	fmt.Println(quantity, reflect.TypeOf(quantity1))*/
	var productName string
	var quantity int
	var discount float32
	var IsInStock bool

	productName = "Golang Dersleri"
	quantity = 11
	discount = 0.35
	IsInStock = true

	//fmt.Println("Product name:", productName, "Quantity", quantity, "Discount", discount, "IsInStock", IsInStock)

	fmt.Printf("Product name: %s,Quantity: %d,Discount:%f,IsInStock: %t", productName, quantity, discount, IsInStock) // pirntf bu yüzden kullanıyoruz.go main

}
