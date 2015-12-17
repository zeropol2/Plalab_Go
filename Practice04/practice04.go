package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    first, second := 0, 0
    return func() int {
	result := first + second
        first, second = second, result

        if second == 0 {
            first++
        }
	
        return result
    }
}

func fibonacci2() func() int {
    current, next := 0, 1
    return func() int {
        result := current
        current, next = next, current+next
        return result
    }
}

func main() {
    f := fibonacci()
    // f := fibonacci2()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

