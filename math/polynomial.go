package main

import (
	"fmt"
	"strconv"
)

// each term of the polynomail is represented as a struct with the
// corresponding coefficient and exponent
type term struct {
	coefficient float64
	exponent    int
}

// Polynomial is a slice of individual terms for the representation of
// a polynomial.
type Polynomial []term

// Write a function that takes a polynomial of the form:
// 		f(x) = a_1x^(b_1)+ a_2x^(b_2) + ... + a_nx^(b_n)
// eg.: f(x) = 5x^4 - 2x^3 + x
// Then it compute the derivative of this polynomial using the constant & power rule
func main() {
	// f(x) = 5x^3 + 11x^5 - 1/2x^-1
	fX := Polynomial{
		{5, 3},
		{11, 5},
		{float64(-1) / 2, -1}}
	fmt.Printf("f(x)=%s\n", fX.print())
	for i := 0; i < 10; i++ {
		fPrime := derive(fX)
		fmt.Printf("f'(x)=%s\n", fPrime.print())
		fX = fPrime
	}
}

func derive(poly Polynomial) Polynomial {
	deriv := make([]term, 0)
	for i := 0; i < len(poly); i++ {
		t := poly[i]
		newExp := t.exponent - 1
		newCoeff := float64(newExp) * t.coefficient
		if newCoeff == 0 {
			continue
		}
		if newExp == 0 {
			newCoeff = 1
		}
		deriv = append(deriv, term{newCoeff, newExp})
	}
	return deriv
}

func (p Polynomial) print() string {
	out := ""
	for _, v := range p {
		sign := "+"
		if v.coefficient < 0 {
			sign = ""
		}
		out += fmt.Sprintf("%s%sx^(%d)",
			sign,
			strconv.FormatFloat(v.coefficient, 'f', -1, 64),
			v.exponent)
	}
	return out
}
