package main

import (
	"fmt"
	"sort"
)

type array struct{
	a []int
}



func main() {
	// arr := array{[]int{2, -5, 6, 7, 1, -8}}
	arr := array{[]int{11, 6, 10, 2, 15, 8, 1, 7, 14, 3, 9, 12, 4, 5, 13}}
	fmt.Println("-----DSelect Algorithm to find ith Smallest Element in an Array (i.e. ith Order Statistic)-----")
	targetEle := dSelect(&arr, 0, len(arr.a)-1, 12)
	fmt.Println("targetEle =", targetEle)
}


// This function finds the ith smallest element of the passed array
// by taking pivot as median of Medians returned from medianOfMedians() function
func dSelect(v *array, p, r int, order int) int{
	target := order-1
	if p >= r{
		return v.a[r]
	}
	// pivot is the pivot element which is a value of it and not the index of pivot element
	pivot := mediansOfMedians(array{v.a[p:r+1]})
	var pivotIdx int
	i, j := p, p
	for j <= r{
		if v.a[j] < pivot {
			v.a[i], v.a[j] = v.a[j], v.a[i]
			i++
		} else if v.a[j] == pivot {
			v.a[i], v.a[j] = v.a[j], v.a[i]
			pivotIdx = i
			i++
		}
		j++
	}
	v.a[pivotIdx], v.a[i-1] = v.a[i-1], v.a[pivotIdx]
	pivotIdx = i-1
	var targetEle int
	if pivotIdx == target{
		targetEle = v.a[pivotIdx]
	} else if target < pivotIdx {
		targetEle = dSelect(v, p, pivotIdx-1, order)
	} else { // target > pivotIdx
		targetEle = dSelect(v, pivotIdx+1, r, order)
	}
	return targetEle
}


// This function is dependency for dSelect() function
// It recursively finds the median of medians and the final median is then returned
func mediansOfMedians(v array) int {
	if len(v.a) == 1{
		return v.a[0]
	}
	medians := array{[]int{}}
	var end int
	for i:=0; i<len(v.a); i=i+5{
		temp := []int{}
		if i+5 <=len(v.a){
			end = i+5
		} else {
			end= len(v.a)
		}
		temp = append(temp, v.a[i:end]...)
		sort.Ints(temp)
		medians.a = append(medians.a, temp[len(temp)/2])
	}
	finalMedian := mediansOfMedians(medians)
	return finalMedian
}