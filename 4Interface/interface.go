package main

import (
	"fmt"

	"daker.wang/Azen/Go-execise/4Interface/mock"
	"daker.wang/Azen/Go-execise/4Interface/real"
)

// Retriever 拉到数据
type Retriever interface {
	Get(url string) string
}

func main() {

	inspectTest()
	typeAssertionTest()
	safeTypeAssertionTest()
}

func safeTypeAssertionTest() {
	var r Retriever
	r = &real.Retriever{}

	realR, ok := r.(*real.Retriever)
	if ok {
		fmt.Println(realR.Contents)	
	}
}

func typeAssertionTest() {
	var r Retriever
	r = &real.Retriever{}

	realR := r.(*real.Retriever)
	fmt.Println(realR.Contents)
}

func inspectTest() {
	var r Retriever

	r = &mock.Retriever{Contents: "这是假数据，嗯嗯！"}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	r = &real.Retriever{}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("\nmock Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("\nreal Contents", v.Contents)
	}
}
