package test_demo

import (
	"fmt"
)

func SliceCap() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 2)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 3)
	fmt.Println(s, len(s), cap(s)) // slice翻倍扩容
}

func SliceParamEdit() {
	var s []int
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

	edit(s)
	fmt.Println(s, len(s), cap(s)) // 修改slice会影响外层
}

func edit(s []int) {
	s[0] = 2
	fmt.Println(s, len(s), cap(s))
}

func SliceParamAppend() {
	s := make([]int, 0, 2)
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

	appendS(s)
	fmt.Println(s, len(s), cap(s)) // 只要append，即使不扩容，也会修改底层数组，不会影响外层
}

func appendS(s []int) {
	s = append(s, 2)
	fmt.Println(s, len(s), cap(s))
}
