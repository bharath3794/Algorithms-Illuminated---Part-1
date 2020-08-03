package main

import(
	"fmt"
	"math/rand"
	"time"
)

type array struct{
	a []int
}


func main() {
	fmt.Println("--------Quick Sort--------")
	arr := array{[]int{3, 4, -5, 1, 6, 8, -2, -8}}
	fmt.Println("Before Sorting, arr =", arr)
	quickSort(&arr, 0, len(arr.a)-1)
	fmt.Println("After Sorting, arr =", arr)
}


// Quick Sort
func quickSort(v *array, p int, r int){
	if p >= r{
		return
	}
	rand.Seed(time.Now().UnixNano())
	// q is the pivot element
	q := rand.Intn(r-p) + p
	v.a[p], v.a[q] = v.a[q], v.a[p]
	i, j := p+1, p+1
	for j <= r{
		if v.a[j] <= v.a[p]{
			v.a[i], v.a[j] = v.a[j], v.a[i]
			i++
		}
		j++
	}
	v.a[p], v.a[i-1] = v.a[i-1], v.a[p]
	q = i-1
	quickSort(v, p, q-1)
	quickSort(v, q+1, r)
}