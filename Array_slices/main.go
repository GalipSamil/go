package main

import "fmt"

func main() {
	/*var names [3]string

	names[0] = "galip"
	names[1] = "HASAN"
	names[2] = "mEHMETY"*/

	/*var names = [4]string{"Mustafa", "Murat", "Galip"} // yada buraya eklersek aynı hatayı alırız.
	names[3] = "Serhat"                                // hata verdi çünkü 3 elemanlı biz 4. eleman eklemeye çalışıyoruz.
	fmt.Println(names[0:2])*/

	var names = []string{"Mustafa", "Murat", "Coşkun"}

	names = append(names, "Serhat")
	fmt.Println(names[2:4])

}
