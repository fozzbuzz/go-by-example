package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Print("hello world (no newline)")

	// String concat
	fmt.Println("go" + "lang")
	fmt.Println("super long string that can go on for ages yet we don't want" +
		" to keep it on the same line in code but it's okay to print on a single" +
		" line on the console")

	// Int math
	fmt.Println("1+1 =", 1+1)

	// boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Long-hand initialization, type inferred
	var a = "initial"
	fmt.Println(a)

	// Multi-variable init
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Zero-value init
	var e int
	fmt.Println(e)

	// Shorthand init. Only available inside functions.
	f := "apple"
	fmt.Println(f)
}
