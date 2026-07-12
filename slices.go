package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 5)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Assign values to the slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	x := append(s, "d")
	x = append(s, "e", "f")
	fmt.Println("append(s):", s)
	fmt.Println("append(x):", x)
	x[3] = "d"
	fmt.Println("append(x):", x)
	fmt.Println("x len:", len(x))
	// Since we appended enough elements to exceed the initial capacity,
	// the new capacity now shows that the underlying array size was doubled.
	fmt.Println("x cap:", cap(x))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// Subslice
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Subslice from beginning to specified end idx (non-inclusive)
	l = s[:5]

	// Subslice that goes to end of slice
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declared:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Create 2D slice. The inner slices will be empty
	twoD := make([][]int, 3)
	fmt.Println("twoD:", twoD)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
	fmt.Println("2d len:", len(twoD))
	fmt.Println("2d cap:", cap(twoD))
}
