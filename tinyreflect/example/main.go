package main

import (
	"fmt"
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
	rt := TypeOf(s)
	rf := rt.Field(0)
	fmt.Println(rf.Name)
}
