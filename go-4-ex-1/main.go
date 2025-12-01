package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float64, maxPoints float64) float64 {
	return 1.0 + (gotPoints/maxPoints)*5.0
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(17.5, 28.0))
	fmt.Println(computeGrade(12.0, 20.0))
	fmt.Println(computeGrade(8.0, 16.0))
}
