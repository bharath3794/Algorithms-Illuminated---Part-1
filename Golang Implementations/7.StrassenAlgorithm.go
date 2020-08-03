package main

import (
	"fmt"
	"math"
)

type array2d struct{
	a [][]int
}

func main() {
	mat1 := array2d{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}}
	mat2 := array2d{[][]int{[]int{1, 2, 2}, []int{0, 1, 0}, []int{0, 0, 1}}}
	fmt.Println("-----Strassen's Matrix Multiplication using Recursion-----")
	rsltMat := strassenMatrixMul(&mat1, &mat2)
	fmt.Println("rsltMat =", *rsltMat)	
}


/*
**************Divide and Conquer Approach****************
*/

/*
* Strassen Algorithm for Matrix Multiplication (Divide and Conquer Approach)
* Applicable for For matrices that are even or uneven 
  [ 𝑚∗𝑛  matrix with either  𝑚  or  𝑛  as a power of 2 or both  𝑚  and  𝑛  are not powers of 2]
*/
func strassenMatrixMul(v1 *array2d, v2 *array2d) *array2d{
	if len(v1.a[0]) != len(v2.a){
 		panic("Two matrices can only be multiplied if and only if columns of 1st matrix = rows of 2nd")
 	}
 	finalRows := len(v1.a)
 	finalCols := len(v2.a[0])
	matrixTransformPower2(v1, v2)
	if len(v1.a) == 1{
		return &array2d{[][]int{[]int{v1.a[0][0]*v2.a[0][0]}}}
	}
	a, b, c, d := splitMatrix(v1)
	e, f, g, h := splitMatrix(v2)
	p1 := strassenMatrixMul(a, subMatrix(f, h))
	p2 := strassenMatrixMul(addMatrix(a, b), h)
	p3 := strassenMatrixMul(addMatrix(c, d), e)
	p4 := strassenMatrixMul(d, subMatrix(g, e))
	p5 := strassenMatrixMul(addMatrix(a, d), addMatrix(e, h))
	p6 := strassenMatrixMul(subMatrix(b, d), addMatrix(g, h))
	p7 := strassenMatrixMul(subMatrix(a, c), addMatrix(e, f))
	term1 := subMatrix(addMatrix(addMatrix(p4, p5), p6), p2)
	term2 := addMatrix(p1, p2)
	term3 := addMatrix(p3, p4)
	term4 := subMatrix(addMatrix(p1, p5), addMatrix(p3, p7))
	rsltMat := combineMatrix(term1, term2, term3, term4)
	removeZeros(rsltMat, finalRows, finalCols)
	return rsltMat
}


/*
*********Dependencies for strassenMatrixMul() & recMatrixMul()************
*/
/*
* Make any two matrices that are not powers of 2 to powers of 2
* [ 𝑚∗𝑛  matrix with either  𝑚  or  𝑛  as a power of 2 or both  𝑚  and  𝑛  are not powers of 2] -> 
  -> [ 𝑛∗𝑛  matrix with  𝑛  as power of 2]
 */
 func matrixTransformPower2(v1 *array2d, v2 *array2d){
 	v1rows := len(v1.a)
 	v1cols := len(v1.a[0])
 	v2rows := len(v2.a)
 	v2cols := len(v2.a[0])
 // This if condition is to stop executing next statements if the given matrices already are
 // powers of two and satisfies all the required conditions to apply divide and conquer approach
 // with recursive calls
 	if v1rows == v1cols && v1rows == v2rows && v1rows == v2cols &&
 		powerOfTwo(v1rows) && powerOfTwo(v1cols) && 
 		powerOfTwo(v2rows) && powerOfTwo(v2cols) {
 			return
 		}
 	var powerTwo int
 	var rows int
 	var cols int
 	var num int
 	if v1rows < v2cols{
 		rows = v2rows
 		cols = v2cols
 	} else {
 		rows = v1rows
 		cols = v1cols
 	}
	if rows < cols{
		num = cols
	} else { // rows >= cols
		num = rows
	}
	if ((num != 0) && (num&(num-1))==0){ // Is num a power of 2? If it is then
		powerTwo = num
	} else { // If not a power of 2
		powerTwo = int(math.Pow(2, float64(int(math.Log2(float64(num))+1))))
	}
	for i:=0; i<powerTwo-v2rows; i++{ 
		temp := make([]int, v2cols)
		v2.a = append(v2.a, temp)
	}
	for i:=0; i<len(v2.a); i++{ // i rows of v2 after modification
		temp := make([]int, powerTwo-v2cols) // powerTwo-v2cols of powerTwo are appended to each row
		v2.a[i] = append(v2.a[i], temp...)
	}
	for i:=0; i<powerTwo-v1rows; i++{
		temp := make([]int, v1cols)
		v1.a = append(v1.a, temp)
	}
	for i:=0; i<len(v1.a); i++{ // i rows of v2 after modification
		temp := make([]int, powerTwo-v1cols) // powerTwo-v2cols of powerTwo are appended to each row
		v1.a[i] = append(v1.a[i], temp...)
	}
}

/*
Add two given matrices
*/
func addMatrix(m1 *array2d, m2 *array2d) *array2d{
	if len(m1.a) != len(m2.a) || len(m1.a[0]) != len(m2.a[0]){
		panic("Two matrices of different rows*columns can't be added")
	}
	sumMat := array2d{[][]int{}}
	for i:=0; i<len(m1.a); i++{
		temp := []int{}
		for j:=0; j<len(m1.a[0]); j++{
			temp = append(temp, m1.a[i][j]+m2.a[i][j])
		}
		sumMat.a = append(sumMat.a, temp)
	}
	return &sumMat
}


/*
Subtract two given matrices
*/
func subMatrix(m1 *array2d, m2 *array2d) *array2d{
	if len(m1.a) != len(m2.a) || len(m1.a[0]) != len(m2.a[0]){
		panic("Two matrices of different rows*columns can't be added")
	}
	sumMat := array2d{[][]int{}}
	for i:=0; i<len(m1.a); i++{
		temp := []int{}
		for j:=0; j<len(m1.a[0]); j++{
			temp = append(temp, m1.a[i][j]-m2.a[i][j])
		}
		sumMat.a = append(sumMat.a, temp)
	}
	return &sumMat
}

/*
* Split the given matrix evenly into four parts. 
* To work correctly the matrix rows and columns should be equal and are powers of 2
*/
func splitMatrix(v *array2d) (*array2d, *array2d, *array2d, *array2d){
	s1 := array2d{[][]int{}}
	s2 := array2d{[][]int{}}
	s3 := array2d{[][]int{}}
	s4 := array2d{[][]int{}}
	n := len(v.a)
	for i:=0; i<n/2; i++{
		s1.a = append(s1.a, v.a[:n/2][i][:n/2])
		s2.a = append(s2.a, v.a[:n/2][i][n/2:])
		s3.a = append(s3.a, v.a[n/2:][i][:n/2])
		s4.a = append(s4.a, v.a[n/2:][i][n/2:])
	}
	return &s1, &s2, &s3, &s4
}


/*
* This do exactly opposite to splitMatrix() function
* Given four parts of the matrix, this matrix combines them into single matrix
*/
func combineMatrix(v1 *array2d, v2 *array2d, v3 *array2d, v4 *array2d) *array2d {
	newMat := array2d{[][]int{}}
	for i:=0; i<len(v1.a); i++{
		newMat.a = append(newMat.a, append([]int{}, v1.a[i]...))
		newMat.a[i] = append(newMat.a[i], v2.a[i]...)
	}
	j := 0
	for i:=len(v1.a); i<len(v3.a)+len(v1.a); i++{
		newMat.a = append(newMat.a, append([]int{}, v3.a[j]...))
		newMat.a[i] = append(newMat.a[i], v4.a[j]...)
		j++
	}
	return &newMat
}


/*
* Remove unnecessary zeros from the final result and 
  get only those possible rows and columns
* This function do the same i.e. given an array, required no.of rows and 
  required no.of columns it in-place modifies the elements to required no.of rows and columns
*/
func removeZeros(v *array2d, reqRows int, reqCols int){
	v.a = v.a[:reqRows]
	for i:=0; i<len(v.a); i++{
		v.a[i] = v.a[i][:reqCols]
	}
}


func powerOfTwo(num int) bool{
	return ((num != 0) && (num&(num-1))==0)
}