package main

import "fmt"

// Classic recursion example: factorial!
func fact(n int) int {
    // Stopping criteria
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main(){
    fmt.Println(fact(7))
}
