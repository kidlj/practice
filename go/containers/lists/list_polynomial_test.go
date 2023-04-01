package lists

import "fmt"

func Example_polynomial_Add() {
	p1 := NewPolynomial()
	p1.Insert(&PolynomialNode{Exponent: 100, Coefficient: 100})
	p1.Insert(&PolynomialNode{Exponent: 1, Coefficient: 3})
	p1.Insert(&PolynomialNode{Exponent: 10, Coefficient: 2})
	p1.Insert(&PolynomialNode{Exponent: 10, Coefficient: 3})
	fmt.Println(p1.String())

	p2 := NewPolynomial()
	p2.Insert(&PolynomialNode{Exponent: 100, Coefficient: 100})
	p2.Insert(&PolynomialNode{Exponent: 10, Coefficient: 3})
	p2.Insert(&PolynomialNode{Exponent: 50, Coefficient: 50})
	p2.Insert(&PolynomialNode{Exponent: 5, Coefficient: 5})
	p2.Insert(&PolynomialNode{Exponent: 1, Coefficient: 4})
	fmt.Println(p2.String())

	result := p1.Add(p2)
	fmt.Println(result.String())

	// Output:
	// 100 * x^100 + 5 * x^10 + 3 * x^1
	// 100 * x^100 + 50 * x^50 + 3 * x^10 + 5 * x^5 + 4 * x^1
	// 200 * x^100 + 50 * x^50 + 8 * x^10 + 5 * x^5 + 7 * x^1
}

func Example_polynomial_Multiply() {
	p1 := NewPolynomial()
	p1.Insert(&PolynomialNode{Exponent: 100, Coefficient: 100})
	p1.Insert(&PolynomialNode{Exponent: 10, Coefficient: 2})
	fmt.Println(p1.String())

	p2 := NewPolynomial()
	p2.Insert(&PolynomialNode{Exponent: 100, Coefficient: 100})
	p2.Insert(&PolynomialNode{Exponent: 10, Coefficient: 5})
	fmt.Println(p2.String())

	result := p1.Multiply(p2)
	fmt.Println(result.String())
	// Output:
	// 100 * x^100 + 2 * x^10
	// 100 * x^100 + 5 * x^10
	// 10000 * x^200 + 700 * x^110 + 10 * x^20
}
