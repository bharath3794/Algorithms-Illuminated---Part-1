package main

import (
	"fmt"
)

type array2d struct{
	a [][]int
}

func main() {
	mat1 := array2d{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}}
	mat2 := array2d{[][]int{[]int{1, 2, 2}, []int{0, 1, 0}, []int{0, 0, 1}}}
	fmt.Println("-----Matrix Multiplication Brute Force Method 1-----")
	product := matrixMulBrute1(&mat1, &mat2)
	fmt.Println("Product =", product)
	fmt.Println("-----Matrix Multiplication Brute Force Method 2-----")
	product = matrixMulBrute2(&mat1, &mat2)
	fmt.Println("Product =", product)
}


/*
**************Brute Force Methods****************
*/

/* 
* Brute Force Approach Method 1
* Algorithm Complexity = O(n^3)
*/
func matrixMulBrute1(v1 *array2d, v2 *array2d) *array2d {
	mulArr := array2d{[][]int{}}
	for i:=0; i<len(v1.a); i++{
		tempArr := []int{}
		for k:=0; k<len(v1.a); k++{
			rslt := 0
			for j:=0; j<len(v1.a[0]); j++{
				rslt += v1.a[i][j] * v2.a[j][k]
			}
			tempArr = append(tempArr, rslt)
		}
		mulArr.a = append(mulArr.a, tempArr)
	}
	return &mulArr
}


/* 
* Brute Force Approach Method 2
* Algorithm Complexity = O(n^3)
*/
func matrixMulBrute2(v1 *array2d, v2 *array2d) *array2d {
	temp := [][]int{}
	for j:=0; j<len(v2.a[0]); j++{ // j columns
		t1 := []int{}
		for i:=0; i<len(v2.a); i++{ // i rows
			t1 = append(t1, v2.a[i][j])
		}
		temp = append(temp, t1)
	}
	v2.a = temp
	prod := [][]int{}
	for i:=0; i<len(v1.a); i++{
		t1 := []int{}
		for j:=0; j<len(v2.a); j++{
			t1 = append(t1, matrixMulBrute2Depend(v1.a[i], v2.a[j]))
		}
		prod = append(prod, t1)
	}
	return &array2d{prod}
}

/* 
* Brute Force Approach Method 2 Dependency
*/
func matrixMulBrute2Depend(a []int, b []int) int{
	rslt := 0
	for i:=0; i<len(a); i++{
		rslt += a[i] * b[i]
	}
	return rslt
}