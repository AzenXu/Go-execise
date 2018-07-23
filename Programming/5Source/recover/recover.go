package main

import (
	"github.com/pkg/errors"
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred：", err)
		} else {
			panic(r)
		}
	}()

	panic(errors.New("来，给爷报个错！"))
}