// https://blog.golang.org/laws-of-reflection
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func Example_interface() {
	var r io.Reader
	var w io.Writer
	var e interface{}
	f, err := os.Open("/tmp/test.txt")
	if err != nil {
		log.Fatal("open file failed.")
	}
	r = f
	w = r.(io.Writer)
	e = w
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", e)
	// Output:
	// *os.File
	// *os.File
	// *os.File
	// *os.File
}

func Example_TypeOf() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	// Output:
	// type: float64
}

func Example_ValueOf() {
	var x float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(x).String())
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v) // by default the fmt package digs into a reflect.Value to show the concrete value inside
	fmt.Println("value:", v.Float())
	// Output:
	// value: <float64 Value>
	// type: float64
	// kind: float64
	// value: 3.4
	// value: 3.4
}

func Example_Getter() {
	// To keep the API simple, the "getter" and "setter" methods of Value operate on the largest type that can hold the value.
	// That is, the Int method of Value returns an int64 and the SetInt value takes an int64.
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Printf("type of Uint(): %T\n", v.Uint())
	fmt.Println("value:", uint8(v.Uint()))
	// Output:
	// type: uint8
	// kind: uint8
	// type of Uint(): uint64
	// value: 120
}

func Example_Kind() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	// The `Kind` of a reflection object describes the underlying type, not the static type
	// In other words, the `Kind` cannot discriminate an int from a MyInt even though the `Type` can.
	fmt.Println("kind:", v.Kind())
	// Output:
	// type: main.MyInt
	// kind: int
}

func Example_Interface() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("value:", v.Interface().(float64))
	fmt.Println("value:", v.Interface())
	// This is odd.
	// see (Why not fmt.Println(v)? Because v is a reflect.Value; we want the concrete value it holds.)
	fmt.Println("value:", v)
	// Output:
	// value: 3.4
	// value: 3.4
	// value: 3.4
}

func Example_settability() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	e := p.Elem()
	fmt.Println("type of e:", e.Type())
	fmt.Println("settability of e:", e.CanSet())
	e.SetFloat(7.1)
	fmt.Println("value of e:", e.Interface())
	fmt.Println("value of x:", x)
	// Output:
	// settability of v: false
	// type of p: *float64
	// settability of p: false
	// type of e: float64
	// settability of e: true
	// value of e: 7.1
	// value of x: 7.1
}

func Example_struct() {
	type T struct {
		A int
		B string
	}
	t := T{
		23,
		"mellon",
	}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(31)
	s.Field(1).SetString("collie")
	fmt.Println("t is now:", t)
	// Output:
	// 0: A int = 23
	// 1: B string = mellon
	// t is now: {31 collie}
}
