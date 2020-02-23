package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

type dimension struct {
	Height int
	Width  int
}

type Restaurant struct {
	NumberOfCustomers *int `json:",omitempty"`
}

type Dog struct {
	Breed    string `json:",omitempty"`
	Weightkh int
	Size     *dimension `json:",omitempty"`
}

func sum(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v {
		ref := reflect.ValueOf(val)
		switch ref.Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32:
			res += float64(ref.Int())
		case reflect.Uint8:
			res += float64(ref.Uint())
		case reflect.String:
			f, err := strconv.ParseFloat(ref.String(), 64)
			if err != nil {
				panic(err)
			}
			res += f
		default:
			fmt.Printf("Unsupported type %T. Ignoring.\n", val)
		}
	}
	return res
}

type MyInt int64

func main() {
	b := bytes.NewBuffer([]byte("hello"))
	if isStringer(b) {
		fmt.Printf("%T is stringer \n", b)
	}
	i := 123
	if isStringer(i) {
		fmt.Printf("%T is a stringer \n", i)
	}
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}

func implements(concrete interface{}, target interface{}) bool {
	iface := reflect.TypeOf(target).Elem()

	v := reflect.ValueOf(concrete)
	t := v.Type()

	if t.Implements(iface) {

	}
}
