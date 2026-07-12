package main

import "fmt"

func main() {
	// Zeroed
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	// Declare size & init with values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init:", b)

	// Init with values w/out declaring size, as well as skipping initializing
	// values for idx 1 and 2
	b = [5]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Declare & init 2D array in two steps
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Declare & init 2D array in one go
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)

	twoD[1][0] = 999_999_999_999_999_999
	fmt.Println("2d:", twoD)
}
