package test

import (
	"fmt"
	"testing"
)

type A struct {
	a, b, c int
}

type B interface {
}

func TestAssert(t *testing.T) {
	var a []A
	a = append(a, A{1, 2, 3})
	a = append(a, A{1, 2, 3})
	a = append(a, A{1, 2, 3})
	a = append(a, A{1, 2, 3})
	var as B = a
	as = as.([]A)
	fmt.Println(as)

}
