package slicefunc

import (
	"testing"
)

func Test_all(t *testing.T) {
	Call(Join(func(fa []func()) {
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

	}), Injs(100))
}
