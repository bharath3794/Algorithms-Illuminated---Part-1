package main

import (
	"fmt"
	"sort"
	"math"
)

type array struct{
	a []int
}


type point2d struct{
	x float64
	y float64
}

type points struct{
	p []point2d
}

func main() {
	arr := array{[]int{2, -5, 6, 7, 1, -8}}
	fmt.Println("-----Closest Pair 1D using Brute Force-----")
	smallDist, ele := closestPair1DBrute(&arr)
	fmt.Println("Smallest Distance =", smallDist, " of elements", ele)
	fmt.Println("-----Closest Pair 1D using sort built-in method-----")
	smallDist, ele = closestPair1DSort(&arr)
	fmt.Println("Smallest Distance =", smallDist, " of elements", ele)


	pointsSlice := points{[]point2d{point2d{2, 3}, point2d{12, 30}, point2d{40, 50}, point2d{5, 1}, 
	       point2d{12, 10}, point2d{3, 4}, point2d{6, 2}, point2d{41, 51}}}
	fmt.Println("-----Closest Pairs 2D using Brute Force-----")
	smallDist1, points1 := closestPair2DBrute(pointsSlice)
	fmt.Println("Smallest Distance =", smallDist1, " of elements", points1)
	fmt.Println("-----Closest Pairs 2D using Recursion-----")
	smallDist1, points1 = recClosestPair2DFinal(pointsSlice)
	fmt.Println("Smallest Distance =", smallDist1, " of elements", points1)

}


/*
****************START******************
* Finding 2D Closest Pairs using Recursion
* Algorithm Complexity = 𝑂(𝑛𝑙𝑜𝑔(𝑛))
*/
func recClosestPair2DFinal(v points) (float64, [][]point2d) {
	sortedX, sortedY := recClosestPair2DSort(&v)
	return recClosestPair2D(sortedX, sortedY)

}

// This function is dependency for recClosestPair2DFinal() function
// This function recursively calls the left half and right half
func recClosestPair2D(sortedX, sortedY points) (float64, [][]point2d) {
	if len(sortedX.p) <= 3 && len(sortedY.p) <= 3{
		return closestPair2DBrute(sortedX)
	}
	mid := len(sortedX.p)/2
	leftX := points{sortedX.p[:mid]}
	rightX := points{sortedX.p[mid:]}
	median := leftX.p[mid-1]
	leftY, rightY := points{[]point2d{}}, points{[]point2d{}}
	for _, v := range sortedY.p{
		if v.x <= median.x{
			leftY.p = append(leftY.p, v)
		} else {
			rightY.p = append(rightY.p, v)
		}
	}
	leftDist, leftClosestPair := recClosestPair2D(leftX, leftY)
	rightDist, rightClosestPair := recClosestPair2D(rightX, rightY)
	splitDist, splitClosestPair := recClosestPair2DSplit(sortedX, sortedY, math.Min(leftDist, rightDist))
	var smallDist float64
	var closestPair [][]point2d
	if leftDist < rightDist && leftDist < splitDist{
		smallDist, closestPair = leftDist, leftClosestPair
	} else if rightDist < leftDist && rightDist < splitDist{
		smallDist, closestPair = rightDist, rightClosestPair
	} else if splitDist < leftDist && splitDist < rightDist{
		smallDist, closestPair = splitDist, splitClosestPair
	} else if leftDist == rightDist{
		leftClosestPair = append(leftClosestPair, rightClosestPair...)
		smallDist, closestPair = leftDist, leftClosestPair 
	} else if rightDist == splitDist{
		rightClosestPair = append(rightClosestPair, splitClosestPair...)
		smallDist, closestPair = rightDist, rightClosestPair 
	} else if splitDist == leftDist{
		splitClosestPair = append(splitClosestPair, leftClosestPair...)
		smallDist, closestPair = splitDist, splitClosestPair 
	} 
	return smallDist, closestPair
}

// This function is dependency for recClosestPair2DFinal() function
// Sorts the given points{} type into two slices, one based on X co-ordinate and other based on Y co-ord
func recClosestPair2DSort(v *points) (points, points){
	temp := make([]point2d, len(v.p))
	copy(temp, v.p)
	sortX := points{temp}
	sort.Slice(sortX.p, func (i, j int) bool {return sortX.p[i].x < sortX.p[j].x})
	temp = make([]point2d, len(v.p))
	copy(temp, v.p)
	sortY := points{temp}
	sort.Slice(sortY.p, func (i, j int) bool {return sortY.p[i].y < sortY.p[j].y})
	return sortX, sortY
}

// This function is dependency for recClosestPair2D() function
// It computes the closest pair across two splits i.e. of left half and right half
func recClosestPair2DSplit(sortedX, sortedY points, delta float64) (float64, [][]point2d){
	median := sortedX.p[len(sortedX.p)/2 - 1]
	Sy := []point2d{}
	for _, v := range sortedY.p{
		if median.x-delta <= v.x &&  v.x <= median.x+delta{
			Sy = append(Sy, v)
		}
	}
	smallDist := delta
	var pairs [][]point2d
	for i:=0; i<len(Sy)-1; i++{
		for j:=i+1; j<int(math.Min(float64(8), float64(len(Sy)-i))); j++{
			curDist := dist(Sy[i], Sy[j])
			if curDist < smallDist{
				smallDist = curDist
				pairs = [][]point2d{[]point2d{Sy[i], Sy[j]}}
			} else if curDist == smallDist{
				pairs = append(pairs, []point2d{Sy[i], Sy[j]})
			}
		}
	}
	return smallDist, pairs

}

/* 
* Finding Closest Pairs 2D using Brute Force technique
* Algorithm Complexity = 𝑂(𝑛^2)
*/
func closestPair2DBrute(v points) (float64, [][]point2d){
	smallDist := math.Inf(1)
	var ele [][]point2d
	for i:=0; i<len(v.p)-1; i++{
		for j:=i+1; j<len(v.p); j++{
			curDist := dist(v.p[i], v.p[j])
			if curDist < smallDist{
				smallDist = curDist
				ele = [][]point2d{[]point2d{v.p[i], v.p[j]}}
			} else if curDist == smallDist{
				ele = append(ele, []point2d{v.p[i], v.p[j]})
			}
		}
	}
	return smallDist, ele
}

// This function finds distance between two given points
func dist(a, b point2d) float64{
	return math.Sqrt(math.Pow((a.x-b.x), 2) + math.Pow((a.y-b.y), 2))
}

/*
*****************END****************************
* End of Finding 2D Closest Pairs using Recursion
*/


/*
* Closest Pair in 1D using Brute Force
* Algorithm Complexity = 𝑂(𝑛^2)
*/
func closestPair1DBrute(v *array) (float64, [][]int) {
	smallDist := math.Inf(1)
	var ele [][]int
	for i:=0; i<len(v.a)-1; i++{
		for j:=i+1; j<len(v.a); j++{
			curDist := abs(v.a[i]-v.a[j])
			if curDist < smallDist{
				smallDist = curDist
				ele = [][]int{[]int{v.a[i], v.a[j]}}
			}else if curDist == smallDist{
				ele = append(ele, []int{v.a[i], v.a[j]})
			}
		}
	}
	return smallDist, ele
}



/*
* Closest Pair in 1D using sort() method
* Algorithm Complexity = 𝑂(𝑛𝑙𝑜𝑔(𝑛))
*/
func closestPair1DSort(v *array) (float64, [][]int) {
	sort.Ints(v.a)
	smallDist := math.Inf(1)
	var ele [][]int
	for i:=0; i<len(v.a)-1; i++{
		curDist := abs(v.a[i]-v.a[i+1])
		if curDist < smallDist{
			smallDist = curDist
			ele = [][]int{[]int{v.a[i], v.a[i+1]}}
		}else if curDist == smallDist{
			ele = append(ele, []int{v.a[i], v.a[i+1]})
		}
	}
	return smallDist, ele
}


// This function returns absolute value of the passed value
func abs(x int) float64{
	if x<0{
		return -float64(x)
	}
	return float64(x)
}