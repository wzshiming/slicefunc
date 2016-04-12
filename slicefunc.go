package slicefunc

import (
	"reflect"

	"github.com/codegangsta/inject"
)

// 默认采用第一个的参数
func Join(fs ...interface{}) interface{} {
	return JoinBy(fs[0], fs...)
}

// 连接函数 返回error 结束 下一次传递
func JoinBy(t interface{}, fs ...interface{}) interface{} {

	val := reflect.ValueOf(t)
	if val.Kind() != reflect.Func {
		return nil
	}

	for k, v := range fs {
		c := toCaller(v)
		if c == nil {
			return nil
		}
		fs[k] = c
	}

	typ := val.Type()
	r := reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
		inj := inject.New()
		for _, v := range args {
			inj.Set(v.Type(), v)
		}
		inj = sliceFunc(inj, fs)
		for i := 0; i != typ.NumOut(); i++ {
			results = append(results, inj.Get(typ.Out(i)))
		}
		return
	})
	return r.Interface()
}

func sliceFunc(inj inject.Injector, fs []interface{}) inject.Injector {
	for _, v := range fs {
		data, err := Call(v, inj)
		if err != nil {
			inj.Map(err)
			return inj
		}
		for _, v := range data {
			inj.Set(v.Type(), v)
			_, ok := v.Interface().(error)
			if ok {
				return inj
			}
		}
	}
	return inj
}

func toCaller(f interface{}) interface{} {
	val := reflect.ValueOf(f)
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Func {
		return f
	}
	cal := val.MethodByName("Call")
	if !cal.IsValid() {
		return nil
	}
	return cal.Interface()
}

func CallArgs(f interface{}, args ...interface{}) ([]reflect.Value, error) {
	return Call(f, Injs(args...))
}

func Call(f interface{}, inj inject.Injector) ([]reflect.Value, error) {
	return inj.Invoke(f)
}

func Injs(args ...interface{}) inject.Injector {
	inj := inject.New()
	for _, v := range args {
		inj.Map(v)
	}
	return inj
}
