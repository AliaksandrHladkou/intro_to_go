//Author: Aliaksandr Hladkou
//Date: 10/05/2019
//Generating array with random numbers

package main

import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    array := randomArray(10)
    fmt.Println("Generated numbers are: ", array)
}

func randomArray(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}