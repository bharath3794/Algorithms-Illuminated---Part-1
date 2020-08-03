package main

import (
	"fmt"
	"math/rand"
	"time"
)

type array struct{
	a []int
}



func main() {
	// arr := array{[]int{2, -5, 6, 7, 1, -8}}
	arr := array{[]int{11, 6, 10, 2, 15, 8, 1, 7, 14, 3, 9, 12, 4, 5, 13}}
	fmt.Println("-----RSelect Algorithm to find ith Smallest Element in an Array (i.e. ith Order Statistic)-----")
	targetEle := rSelect(&arr, 0, len(arr.a)-1, 7)
	fmt.Println("targetEle =", targetEle)
}


// This function finds the ith smallest element of the passed array
// by taking pivot as random index
func rSelect(v *array, p, r int, order int) int{
	target := order-1
	if p >= r{
		return v.a[r]
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
	var targetEle int
	if q == target{
		targetEle = v.a[q]
	} else if target < q {
		targetEle = rSelect(v, p, q-1, order)
	} else { // target > q
		targetEle = rSelect(v, q+1, r, order)
	}
	return targetEle
}