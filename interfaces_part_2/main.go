package main

import "fmt"

type XEncoder struct {
}

func (xEncoder *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi")
}

func (xEncoder *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi ")
}

func main() {
	var xEncoder *XEncoder
	xEncoder = &XEncoder{}
	xEncoder.Encode("234234")
	XEncoder.Decode("432423")

}
