package main

import "fmt"
//import "sort"
import "math/rand"

func main() {
	fmt.Println("RandArr ", randomArray(10))

	fmt.Println("MergeSort: ", mergeSort(randomArray(10)))
}

func randomArray(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}


func mergeSort(slice []int) {
	if len(slice) < 2 {
		return slice
	}

	middle := len(slice) / 2
	merge(mergeSort(slice[:middle]), mergeSort(slice[middle:]))
}

func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}