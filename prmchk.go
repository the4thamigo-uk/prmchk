package prmchk

import (
	"reflect"
)

// Check recursively scans the struct type for cases where a field promotion is overridden by a field of the same name
func Check(val interface{}) bool {
	return CheckType(reflect.TypeOf(val))
}

// CheckType recursively scans the struct type for cases where a field promotion is overridden by a field of the same name
func CheckType(t reflect.Type) bool {
	return check(t)
}

func check(t reflect.Type, t2s ...reflect.Type) bool {
	t2s = append(t2s, t)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type
		if ft.Kind() == reflect.Struct {
			if f.Anonymous && fieldConflicts(ft, t2s...) {
				return true
			}
			if check(ft, t2s...) {
				return true
			}
		}
	}
	return false
}

// fieldConflict checks whether there are any fields in t1 that have the same name as any fields in
// any of the types t2s. All types passed in should be of type reflect.Struct
func fieldConflicts(t1 reflect.Type, t2s ...reflect.Type) bool {
	for _, t2 := range t2s {
		if fieldConflict(t1, t2) {
			return true
		}
	}
	return false
}

// fieldConflict checks whether there are any fields in t1 that have the same name as any fields in t2
func fieldConflict(t1, t2 reflect.Type) bool {
	fieldNames := func(t reflect.Type) (fs []string) {
		for i := 0; i < t.NumField(); i++ {
			fs = append(fs, t.Field(i).Name)
		}
		return
	}
	fn1 := fieldNames(t1)
	fn2 := fieldNames(t2)
	return intersect(fn1, fn2)
}

func intersect(fn1, fn2 []string) bool {
	for _, n1 := range fn1 {
		for _, n2 := range fn2 {
			if n1 == n2 {
				return true
			}
		}
	}
	return false
}
