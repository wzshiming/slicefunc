package slicefunc

import (
	"testing"
)

func Test_all(t *testing.T) {
	nf1 := JoinBy(func(fa []func()) {
		fa[3]()
		fa[2]()
		fa[1]()
		fa[0]()
		fa[0]()
		fa[1]()
		fa[2]()
		fa[3]()
	}, func(int) {}, func() {
		t.Log(1)
	}, func() {
		t.Log(2)
	}, func() {
		t.Log(3)
	}, func(i int) {
		t.Log(i)

	})
	Call(nf1, Injs(100))

	nf2 := Join(nf1, nf1, func() {
		t.Log("end")
	})

	Call(nf2, Injs(100))
}
