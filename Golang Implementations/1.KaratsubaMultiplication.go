package main

import(
	"fmt"
	"strconv"
	"math"
	"math/big"
)



func main() {
	a, b := "9999999", "99999999"
	fmt.Println("-----Karatsuba Multiplication Algorithm------")
	product := karatsubaMul(a, b)
	fmt.Println("Product =", product)
	a = "3141592653589793238462643383279502884197169399375105820974944592"
	b = "2718281828459045235360287471352662497757247093699959574966967627"
	fmt.Println("-----Karatsuba Multiplication Algorithm for Big Integers------")
	bigProduct := karatsubaMulBigInt(a, b)
	fmt.Println("bigProduct =", bigProduct)
}

// Karatsuba Multiplication Algorithm for Normal Integers that fits int value in GO
func karatsubaMul(x string, y string) int{
	if x == "0" || y == "0"{
		return 0
	}
	digitsX := len(x)
	digitsY := len(y)
	if digitsX == 1 || digitsY == 1{
		a, _ := strconv.Atoi(x)
		c, _ := strconv.Atoi(y)
		return a*c
	}
	var zeros int
	if digitsX >= digitsY{
		if !((digitsX != 0) && (digitsX&(digitsX-1))==0){
			zeros = int(math.Pow(2, float64(int(math.Log2(float64(digitsX))+1))))
		} else {
			zeros = digitsX
		}
	} else { // digitsY > digitsX
		if !((digitsY != 0) && (digitsY&(digitsY-1))==0){
			zeros = int(math.Pow(2, float64(int(math.Log2(float64(digitsY))+1))))
		} else {
			zeros = digitsY
		}
	}
	for i:=0; i<zeros-digitsX; i++{
		x = "0" + x
	}
	for i:=0; i<zeros-digitsY; i++{
		y = "0" + y
	}
	a, b := x[:len(x)/2], x[len(x)/2:]
	c, d := y[:len(y)/2], y[len(y)/2:]
	step1 := karatsubaMul(a, c)
	step2 := karatsubaMul(b, d)
	t1, _ := strconv.Atoi(a)
	t2, _ := strconv.Atoi(b)
	t3, _ := strconv.Atoi(c)
	t4, _ := strconv.Atoi(d)
	step3 := karatsubaMul(strconv.Itoa(t1+t2), strconv.Itoa(t3+t4))
	step4 := step3 - step1 - step2
	rslt := int(math.Pow(10, float64(len(x))))*step1 + int(math.Pow(10, float64(len(x)/2)))*step4 + step2
	return rslt
}


// Dependency for below Algorithms
// Useful for big integers like for ex. 123456789123456789123456789
func bigInt(s string) *big.Int{
	val, _ := big.NewInt(0).SetString(s, 10)
	return val
}

// Karatsuba Multiplication Algorithm for Big Integers
func karatsubaMulBigInt(x string, y string) *big.Int{
	if x == "0" || y == "0"{
		return bigInt("0")
	}
	digitsX := len(x)
	digitsY := len(y)
	if digitsX == 1 || digitsY == 1{
		a := bigInt(x)
		c := bigInt(y)
		return big.NewInt(0).Mul(a, c)
	}
	var zeros int
	if digitsX >= digitsY{
		if !((digitsX != 0) && (digitsX&(digitsX-1))==0){
			zeros = int(math.Pow(2, float64(int(math.Log2(float64(digitsX))+1))))
		} else {
			zeros = digitsX
		}
	} else { // digitsY > digitsX
		if !((digitsY != 0) && (digitsY&(digitsY-1))==0){
			zeros = int(math.Pow(2, float64(int(math.Log2(float64(digitsY))+1))))
		} else {
			zeros = digitsY
		}
	}
	for i:=0; i<zeros-digitsX; i++{
		x = "0" + x
	}
	for i:=0; i<zeros-digitsY; i++{
		y = "0" + y
	}
	a, b := x[:len(x)/2], x[len(x)/2:]
	c, d := y[:len(y)/2], y[len(y)/2:]
	step1 := karatsubaMulBigInt(a, c)
	step2 := karatsubaMulBigInt(b, d)
	t1 := big.NewInt(0).Add(bigInt(a), bigInt(b)).String()
	t2 := big.NewInt(0).Add(bigInt(c), bigInt(d)).String()
	step3 := karatsubaMulBigInt(t1, t2)
	step4 := big.NewInt(0).Sub(step3, big.NewInt(0).Add(step1, step2))
	term1 := big.NewInt(0).Mul(big.NewInt(0).Exp(bigInt("10"), big.NewInt(int64(len(x))), nil), step1)
	term2 := big.NewInt(0).Mul(big.NewInt(0).Exp(bigInt("10"), big.NewInt(int64(len(x)/2)), nil), step4)
	term3 := big.NewInt(0).Add(term1, term2)
	rslt := big.NewInt(0).Add(term3, step2)
	return rslt
}