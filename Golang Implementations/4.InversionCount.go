package main

import (
	"fmt"
)

type array struct{
	a []int
}

func main() {
	arr := array{[]int{1, 8, 9, 10, 2, 4, 7, 5, 6, 3}}
	fmt.Println("-----Counting Inversions using Brute Force-----")
	inversions := inversionsBruteForce(&arr)
	fmt.Println("No.of Inversions =", inversions)
	fmt.Println("-----Counting Inversions using Recursion (Divide & Conquer)-----")
	sortedArr, inversions := recInversionCount(&arr)
	fmt.Println("sortedArr =", sortedArr, "and No.of Inversions =", inversions)
}


/* 
* Divide and Conquer Approach
* Algorithm Complexity = O(nlog(n))
*/
func recInversionCount(v *array) (*array, int){
	if len(v.a) == 1{
		return v, 0
	}
	mid := (len(v.a))/2
	leftHalf, leftInv := recInversionCount(&array{v.a[:mid]})
	rightHalf, rightInv := recInversionCount(&array{v.a[mid:]})
	sortedArr := []int{}
	splitInv := 0
	i, j := 0, 0
	for i < len(leftHalf.a) && j < len(rightHalf.a){
		if leftHalf.a[i] <= rightHalf.a[j]{
			sortedArr = append(sortedArr, leftHalf.a[i])
			i++
		} else{
			sortedArr = append(sortedArr, rightHalf.a[j])
			splitInv += len(leftHalf.a)-i
			j++
		}
	}
	if i<len(leftHalf.a) {
		sortedArr = append(sortedArr, leftHalf.a[i:]...)
	} 
	for j < len(rightHalf.a) {
		sortedArr = append(sortedArr, rightHalf.a[j])
		j++
	}
	return &array{sortedArr}, leftInv+rightInv+splitInv
}

/* 
* Brute Force Approach
* Algorithm Complexity = O(n^2)
*/
func inversionsBruteForce(v *array) int {
	var cnt int
	for i:=0; i<len(v.a)-1; i++{
		for j:=i; j<len(v.a); j++{
			if v.a[i] > v.a[j]{
				cnt += 1
			}
		}
	}
	return cnt
}

