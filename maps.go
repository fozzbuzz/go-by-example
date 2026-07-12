package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 2
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Access is via copy, not reference
	v1 = 55
	fmt.Println("v1:", v1)
	fmt.Println("map v1:", m["k1"])

	// For non-existing keys, the zeroed value is returned (why...)
	// Apparently useful for traversal to detect cycles, for example
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Check for existence like this
	_, ok := m["k3"]
	fmt.Println("v3 exists:", ok)

	// Check for existence with condition
	if _, ok := m["k3"]; ok {
		fmt.Println("k3 exists")
	} else {
		fmt.Println("k3 doesn't exist")
	}

	// Built-ins
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))
	delete(m, "k2")
	fmt.Println("map post-delete:", m)
	clear(m)
	fmt.Println("map post-clear:", m)

	// In-line init
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	// Iterator over kv pairs
	for k, v := range n2 {
		fmt.Println(k, v)
	}

	// Shallow copy if values are reference types
	n3 := map[string][]int{"foo": {1, 2, 3}}
	n3Clone := maps.Clone(n3)
	n3Clone["foo"][0] = 99
	n3Match := true
	for i := range 3 {
		if n3["foo"][i] != n3Clone["foo"][i] {
			n3Match = false
			break
		}
	}
	fmt.Println("n3 == n3Clone:", n3Match)

	// Changing value of key in the clone doesn't change the OG value
	n3Clone["foo"] = []int{98, 99, 100}
	fmt.Println("n3:", n3)
	fmt.Println("n3Clone:", n3Clone)

	// Copy keys to dst from src
	m3 := map[string][]int{
		"one": {1, 2, 3},
		"two": {4, 5, 6},
	}
	m4 := map[string][]int{
		"one": {7, 8, 9},
	}

	maps.Copy(m4, m3)
	fmt.Println("m4 is:", m4)

	m4["one"][0] = 100
	fmt.Println("m3 is:", m3)
	fmt.Println("m4 is:", m4)

	// Custom entry deletion
	maps.DeleteFunc(n3Clone, func(k string, v []int) bool {
		return k == "foo"
	})
	fmt.Println("n3Clone should be empty:", n3Clone)

	// Insert from sequence
	m1 := map[int]string{4: "four"}
	s1 := []string{"zero", "one", "two", "three"}
	maps.Insert(m1, slices.All(s1))
	fmt.Println("m1:", m1)

	// Get keys
	keys := slices.Sorted(maps.Keys(m1))
	fmt.Println("keys:", keys)

	// Get values
	values := slices.Sorted(maps.Values(m1))
	fmt.Println("values:", values)
}
