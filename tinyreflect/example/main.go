package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	A int
	B string
}

func main() {
	s := S{
		A: 1,
		B: "abcde",
	}
	x := *(*int)(unsafe.Pointer(&s))
	fmt.Printf("A: %d\n", x)

	y := *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.B)))
	fmt.Printf("B: %s\n", y)

	fmt.Printf("offset A : %d\n", unsafe.Offsetof(s.A))
	fmt.Printf("offset B : %d\n", unsafe.Offsetof(s.B))

	rt := TypeOf(s)
	if false {
		fmt.Printf("%#v\n", rt)
		//fmt.Println(rf.Name)
	}
	//rf := rt.Field(0)

	//{
	//	rf := rt.Field(1)
	//	fmt.Println(rf.Name)
	//}
}
