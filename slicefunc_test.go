package slicefunc

import (
	"fmt"
	"testing"
)

func Test_all(t *testing.T) {

	nf2 := Join(A1, A2, A{}, A{})
	f, _ := nf2.(func(i int) string)
	f(100)
	//Call(nf2, Injs(100))
}

type A struct {
}

func (A) Call(s string) error {
	fmt.Printf(s)
	return fmt.Errorf("end")
}

func A1(i int) string {
	return fmt.Sprint(i + 10)
}

func A2(i int, s string) (int, string) {
	return i + 50, fmt.Sprint(i) + " " + s
}
