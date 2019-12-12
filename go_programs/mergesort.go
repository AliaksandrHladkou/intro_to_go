package main

import "fmt"
import "math/rand"
import "time"


func main() {
    
    rand.Seed(time.Now().UTC().UnixNano())
    array := randomArray(100000)
    
    
    result := make(chan []int)	//creating new channel 
    start := time.Now()
    go mergeSort(array, result)	//invoke the function using channel

    r := <- result 			//recieve from the channel and assign to r

	//elapsed := time.Since(start)
    //fmt.Println("Array with first value of ", r[0], " and time: ", elapsed)   

    t := time.Now()
    elapsed := t.Sub(start)

    fmt.Println(r[0], " Time ", elapsed)

    //for _,v := range r {
    //    fmt.Println(v)
    //}
    close(result)	//closing channel
}

func randomArray(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}

func mergeSort(data []int, r chan []int) {
    if len(data) == 1 {
        r <- data
        return
    }

    leftChan := make(chan []int)
    rightChan := make(chan []int)
    middle := len(data)/2

    go mergeSort(data[:middle], leftChan)
    go mergeSort(data[middle:], rightChan)

    left := <-leftChan
    right := <-rightChan

    close(leftChan)
    close(rightChan)
    r <- merge(left, right)
    return
}

func merge(left []int, right []int) (result []int) {
    result = make([]int, len(left) + len(right))
    lidx, ridx := 0, 0

    for i:=0;i<cap(result);i++ {
        switch {
            case lidx >= len(left):
                result[i] = right[ridx]
                ridx++
            case ridx >= len(right):
                result[i] = left[lidx]
                lidx++
            case left[lidx] < right[ridx]:
                result[i] = left[lidx]
                lidx++
            default:
                result[i] = right[ridx]
                ridx++
        }
    }

    return
}