package slice_utils

import (
	"math/rand"
	"reflect"
	"time"
)

// slice 隨機排序
func RandShuffle(slice interface{}) []interface{} {
	anyTypeSlice := createAnyTypeSlice(slice)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(anyTypeSlice), func(i, j int) {
		anyTypeSlice[i], anyTypeSlice[j] = anyTypeSlice[j], anyTypeSlice[i]
	})

	return anyTypeSlice
}

func createAnyTypeSlice(slice interface{}) []interface{} {
	val, ok := isSlice(slice)

	if !ok {
		return nil
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out
}

//Determine whether it is slcie data
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
