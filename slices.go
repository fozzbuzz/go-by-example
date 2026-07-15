package main

import (
	"fmt"
	"slices"
	"strings"
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
	fmt.Printf("type:%T\n", twoD)

	// Other slices functions
	fmt.Println("All")
	for i, v := range slices.All(t2) {
		fmt.Println(i, ":", v)
	}

	fmt.Println("Backward")
	for i, v := range slices.Backward(t2) {
		fmt.Println(i, ":", v)
	}

	fmt.Println("Chunk")
	for c := range slices.Chunk(t2, 2) {
		fmt.Println(c)
	}

	fmt.Println("Clip")
	a := [...]int{0, 1, 2, 3, 4, 5}
	// low (incl) : high (excl) : max cap
	b := a[:3:6]
	fmt.Println(b)
	fmt.Println("len:", len(b))
	fmt.Println("cap:", cap(b))
	clip := slices.Clip(b)
	fmt.Println(clip)
	fmt.Println("len:", len(clip))
	fmt.Println("cap:", cap(clip))

	fmt.Println("Clone")
	d := [][2]int{{0, 1}, {2, 3}}
	clone := slices.Clone(d)
	fmt.Println("og:", d)
	fmt.Println("clone:", clone)
	clone[0][0] = 99
	fmt.Println("og:", d)
	fmt.Println("clone:", clone)

	fmt.Println("Compact")
	e := []int{0, 1, 1, 2, 2, 3}
	comp := slices.Compact(e)
	fmt.Println(comp)
	fmt.Println("len:", len(comp))
	fmt.Println("cap:", cap(comp))
	comp = slices.Clip(comp)
	fmt.Println("cap:", cap(comp))

	fmt.Println("Concat")
	s1 := []int{0, 1, 2}
	s2 := []int{3, 4, 5}
	s1_s2 := slices.Concat(s1, s2)
	fmt.Println(s1_s2)
	s1 = append(s1, 6)
	// This is a different slice
	fmt.Println(s1)

	fmt.Println("Contains")
	s3 := []int{1, 3, 5}
	fmt.Println(slices.Contains(s3, 1))
	fmt.Println(slices.Contains(s3, 2))

	fmt.Println("ContainsFunc")
	fmt.Println("has even:", slices.ContainsFunc(s3, func(x int) bool {
		return x%2 == 0
	}))

	fmt.Println("Delete")
	s3 = slices.Delete(s3, 1, 2)
	fmt.Println(s3)
	fmt.Println("cap:", cap(s3))

	fmt.Println("DeleteFunc")
	s4 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s4)
	s4 = slices.DeleteFunc(s4, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("odds:", s4)

	fmt.Println("Equal")
	s5 := []byte{0, 22, 33, 44}
	s5_1 := []byte{0, 22, 33, 44}
	fmt.Println(s5)
	fmt.Println(s5_1)
	fmt.Println("s5 == s5_1:", slices.Equal(s5, s5_1))
	s5_1 = []byte{0, 33, 22, 44}
	fmt.Println(s5)
	fmt.Println(s5_1)
	fmt.Println("s5 == s5_1:", slices.Equal(s5, s5_1))

	fmt.Println("Grow")
	grow := []string{"foo", "bar"}
	fmt.Println("cap:", cap(grow))
	grow = slices.Grow(grow, 9999)
	fmt.Println("grow:", cap(grow))

	fmt.Println("Index")
	idx := []rune{97, 98, 99}
	fmt.Println(idx)
	fmt.Println("idx(98):", slices.Index(idx, 98))
	fmt.Println("idx(44):", slices.Index(idx, 44))

	fmt.Println("IndexFunc")
	idxf := []rune{1, 2, 3, 4}
	fmt.Println(idxf)
	fmt.Println("idx(even):", slices.IndexFunc(idxf, func(x rune) bool {
		return x%2 == 0
	}))

	fmt.Println("Insert")
	insert := []uint16{33, 44, 55}
	fmt.Println(insert)
	fmt.Println("cap:", cap(insert))
	insert = slices.Insert(insert, 0, 10, 11, 12, 13, 14, 15)
	fmt.Println(insert)
	fmt.Println("cap:", cap(insert))

	fmt.Println("IsSorted")
	fmt.Println(insert)
	fmt.Println("sorted:", slices.IsSorted(insert))
	insert = append(insert, 1)
	fmt.Println(insert)
	fmt.Println("sorted:", slices.IsSorted(insert))

	fmt.Println("IsSortedFunc")
	sortedf := []string{"alex", "ALEX", "maya"}
	fmt.Println(sortedf)
	fmt.Println("sortedf:", slices.IsSortedFunc(sortedf, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	}))
}
