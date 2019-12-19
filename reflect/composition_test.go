package reflect

import "reflect"

import "fmt"

type Circle struct {
	Point
	Radius int
	Interface
}

type Interface interface {
	Hello() string
}

type Hello struct{}

func (h Hello) Hello() string {
	return "hello"
}

type Point struct {
	X int
	Y int
}

func (p Point) Distance() int {
	return 0
}

func Example_Argument() {
	c := Circle{
		Point: Point{
			X: 1,
			Y: 2,
		},
		Radius:    3,
		Interface: Hello{},
	}
	fmt.Println(typ(Circle.Distance))
	fmt.Println(typ(Point.Distance))
	fmt.Println(c.Hello())
	fmt.Println(typ(c.Hello))

	// Output:
	// 1
	// Circle
	// 1
	// Point
	// hello
	// 0
	// None
}

func typ(t interface{}) string {
	rt := reflect.TypeOf(t)
	fmt.Println(rt.NumIn())
	if rt.NumIn() > 0 {
		return rt.In(0).Name()
	}
	return "None"
}
