package main

import(
	"fmt"
	"strconv"
	"strings"
)


func main() {
	a, b := "9999999", "99999999"
	fmt.Println("------Multiplication Algorithm based on 3rd Grade------")
	cusProduct  := multiply(a, b)
	fmt.Println("cusProduct =",cusProduct)
}



/*
Start of the Multiplication Algorithm based on 3rd Grade;
Check for the ***End***
*/
func multiply(s1 string, s2 string) string{
	prodSlice := []string{}
	for i:=len(s2)-1; i>=0; i--{
		prod := myMul(s1, s2[i:i+1])
		prodSlice = append(prodSlice, prod)
	}
	for i:=0; i<len(prodSlice); i++{
		idx := len(prodSlice)-i-1
		for j:=0; j<len(s2)-i-1; j++{
			prodSlice[idx] += "0"
		}
	}
	for i:=0; i<len(prodSlice)-1; i++{
		zeros := len(prodSlice[len(prodSlice)-1])-len(prodSlice[i])
		for j:=0; j<zeros; j++{
			prodSlice[i] = "0"+prodSlice[i]
		}
	}
	temp := prodSlice[0]
	for i:=1; i<len(prodSlice); i++{
		temp = myAdd(temp, prodSlice[i])
	}
	return temp
}


func myMul(s1 string, s2 string) string {
	if len(s2) != 1{
		panic("The second element passed to the function must be of single digit")
	}
	if s2 == "0"{
		return "0"
	} else if s2 == "1"{
		return s1
	}
	prodIntSlice := []string{"0"}
	carryIntSlice := []string{}
	v2, _ := strconv.Atoi(s2)
	for i:= 0; i<len(s1); i++{
		v1, _ := strconv.Atoi(s1[i:i+1])
		carry, prod := bitMul(v1, v2)
		prodIntSlice = append(prodIntSlice, strconv.Itoa(prod))
		carryIntSlice = append(carryIntSlice, strconv.Itoa(carry))
	}
	carryIntSlice = append(carryIntSlice, "0")
	totalStr := myAdd(strings.Join(prodIntSlice, ""), strings.Join(carryIntSlice, ""))
	return totalStr
}

func bitMul(i1 int, i2 int) (int, int){
	if i1>9 || i2>9{
		return -1, -1
	}
	if i1 == 0 || i2 == 0{
		return 0, 0
	}
	p := i1*i2
	carry := p/10
	prod := p%10
	if carry != 0{
		carry = carry%10
	}
	return carry, prod
}

func myAdd(s1 string, s2 string) string{
	if len(s1) != len(s2){
		panic("Strings are of unequal length. Must be of equal length")
	}
	sumIntSlice := []int{0}
	carryIntSlice := []int{}
	for i:= 0; i<len(s1); i++{
		v1, _ := strconv.Atoi(s1[i:i+1])
		v2, _ := strconv.Atoi(s2[i:i+1])
		carry, sum := bitAdd(v1, v2)
		sumIntSlice = append(sumIntSlice, sum)
		carryIntSlice = append(carryIntSlice, carry)
	}
	carryIntSlice = append(carryIntSlice, 0)
	totalSlice := []string{}
	for i:=0; i<len(carryIntSlice); i++{
		totalSlice = append(totalSlice, strconv.Itoa(sumIntSlice[i]+carryIntSlice[i]))
	}
	totalStr := strings.Join(totalSlice, "")
	if string(totalStr[0]) == "0"{
		totalStr = totalStr[1:]
	}
	return totalStr
}

func bitAdd(i1 int, i2 int) (int, int){
	if i1>9 || i2>9{
		return -1, -1
	}
	if i1 == 0 && i2 == 0{
		return 0, 0
	}
	s := i1 + i2
	carry := s/10
	sum := s%10
	if carry != 0{
		carry = carry%10
	}
	return carry, sum
}
