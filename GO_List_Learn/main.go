package main

import "fmt"

func main() {
	// L := new(LNode)
	var L List
	InitList(&L)
	CreatList(&L, 2)
	Pt(L)
	// b := Book{
	// 	isbn: "000",
	// 	name: "xyz",
	// 	price: 9.99,
	// }
	// insert(&L,b,3)
	// delete(&L, "000")
	Pt(L)
	fmt.Println(getindex(L, "002"))

}
