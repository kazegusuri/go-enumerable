package enumerable

import (
	"fmt"
	"reflect"
)

func Map(in, fn interface{}) interface{} {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()
	ret := reflect.MakeSlice(inv.Type(), 0, len)
	for i := 0; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v})
		ret = reflect.Append(ret, value[0])
	}
	return ret.Interface()
}

func Map3(in, out, fn interface{}) error {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()
	outType := reflect.Indirect(reflect.ValueOf(out)).Type()
	outv := reflect.MakeSlice(outType, 0, len)
	for i := 0; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v})
		outv = reflect.Append(outv, value[0])
	}
	reflect.ValueOf(out).Elem().Set(outv)
	return nil
}

func Reduce(in, out, fn interface{}) error {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()
	if len == 0 {
		return fmt.Errorf("empty slice")
	}

	result := inv.Index(0)
	for i := 1; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v, result})
		result = value[0]
	}
	reflect.ValueOf(out).Elem().Set(result)
	return nil
}

func Filter(in, out, fn interface{}) error {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()
	outType := reflect.Indirect(reflect.ValueOf(out)).Type()
	outv := reflect.MakeSlice(outType, 0, len)
	for i := 0; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v})
		if value[0].Bool() {
			outv = reflect.Append(outv, v)
		}
	}
	reflect.ValueOf(out).Elem().Set(outv)
	return nil
}

func Min(in, out, fn interface{}) error {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()
	if len == 0 {
		return fmt.Errorf("empty slice")
	}

	min := inv.Index(0)
	for i := 1; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v, min})
		if value[0].Int() < 0 {
			min = v
		}
	}
	reflect.ValueOf(out).Elem().Set(min)
	return nil
}

func Max(in, out, fn interface{}) error {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()
	if len == 0 {
		return fmt.Errorf("empty slice")
	}

	max := inv.Index(0)
	for i := 1; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v, max})
		if value[0].Int() > 0 {
			max = v
		}
	}
	reflect.ValueOf(out).Elem().Set(max)
	return nil
}

func Any(in, fn interface{}) (bool, error) {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()

	for i := 0; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v})
		if value[0].Bool() {
			return true, nil
		}
	}

	return false, nil
}

func All(in, fn interface{}) (bool, error) {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()

	for i := 0; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v})
		if !value[0].Bool() {
			return false, nil
		}
	}

	return true, nil
}

func Each(in, fn interface{}) error {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()

	for i := 0; i < len; i++ {
		v := inv.Index(i)
		fnv.Call([]reflect.Value{v})
	}

	return nil
}

func Find(in, out, fn interface{}) error {
	inv := reflect.ValueOf(in)
	fnv := reflect.ValueOf(fn)
	len := inv.Len()

	for i := 0; i < len; i++ {
		v := inv.Index(i)
		value := fnv.Call([]reflect.Value{v})
		if value[0].Bool() {
			reflect.ValueOf(out).Elem().Set(v)
			return nil
		}
	}
	return fmt.Errorf("not found")
}
