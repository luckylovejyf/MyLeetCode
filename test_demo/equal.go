package test_demo

import "fmt"

type A struct {
	ID   int
	Name string
}

type B struct {
	ID   int
	Name string
}

type C struct {
	Name string
	ID   int
}

type D struct {
	ID1  int
	Name string
}

func TypeEqual() {
	a := struct {
		ID   int
		Name string
	}{2, "1"}
	b := struct {
		ID   int
		Name string
	}{1, "1"}
	/*c := struct {
		Name string
		ID   int
	}{"1", 1}
	d := struct {
		ID1  int
		Name string
	}{1, "1"}*/

	fmt.Println(a == b) // 可比较
	// fmt.Println(a == c) // Invalid operation: a == c (mismatched types struct {...} and struct {...})
	// fmt.Println(a == d) // Invalid operation: a == d (mismatched types struct {...} and struct {...})
}

func NilEqual() {
	var ptr *int64 = nil
	var cha chan int64 = nil
	// var fun func() = nil
	var inter interface{} = nil
	// var ma map[string]string = nil
	// var slice []int64 = nil
	// fmt.Println(ptr == cha)
	// fmt.Println(ptr == fun)
	fmt.Println(ptr == inter)
	// fmt.Println(ptr == ma)
	// fmt.Println(ptr == slice)

	// fmt.Println(cha == fun)
	fmt.Println(cha == inter)
	// fmt.Println(cha == ma)
	// fmt.Println(cha == slice)

	// fmt.Println(fun == inter)
	// fmt.Println(fun == ma)
	// fmt.Println(fun == slice)

	// fmt.Println(inter == ma)
	// fmt.Println(inter == slice)

	// fmt.Println(ma == slice)
}
