package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println("constant:", s)
    
    // Has no type until given one, like explicit conversion or by context
    const n = 500_000_000
    const f = 50.0 / n
    fmt.Printf("n type: %T\n", n)

    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
    // If math.Sin expects float, why is n still an int?
    fmt.Printf("n type: %T\n", n)
}
