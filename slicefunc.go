package slicefunc

import (
	"reflect"

	"github.com/codegangsta/inject"
)

func Join(f func([]func()), t interface{}, fs ...interface{}) interface{} {
	if len(fs) == 0 {
		return nil
	}
	if len(fs) == 1 {
		return fs[0]
	}
	typ := reflect.ValueOf(t)
	if typ.Kind() != reflect.Func {
		return nil
	}
	r := reflect.MakeFunc(typ.Type(), func(args []reflect.Value) (results []reflect.Value) {
		inj := inject.New()
		for _, v := range args {
			inj.Map(v.Interface())
		}
		funcs := []func(){}
		for _, v := range fs {
			tt := v
			funcs = append(funcs, func() {
				Call(tt, inj)
			})
		}
		f(funcs)
		return
	})
	return r.Interface()
}

func Call(f interface{}, inj inject.Injector) {
	vs := reflect.ValueOf(f)
	vt := vs.Type()
	args0 := []reflect.Value{}
	for i := 0; i != vt.NumIn(); i++ {
		arg := inj.Get(vt.In(i))
		if !arg.IsValid() {
			arg = reflect.New(vt.In(i)).Elem()
		}
		args0 = append(args0, arg)
	}
	vs.Call(args0)
}

func Injs(args ...interface{}) inject.Injector {
	inj := inject.New()
	for _, v := range args {
		inj.Map(v)
	}
	return inj
}
