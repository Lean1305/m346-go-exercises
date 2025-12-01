package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) []float64 {
	D := b*b - 4*a*c

	if D > 0 {
		return []float64{
			(-b + math.Sqrt(D)) / (2 * a),
			(-b - math.Sqrt(D)) / (2 * a),
		}
	}

	if D == 0 {
		return []float64{
			-b / (2 * a),
		}
	}

	return []float64{} // keine Lösung
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // zwei Lösungen
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // eine Lösung
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // keine Lösung
}
