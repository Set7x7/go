package main

import (
	"fmt"
)

func main() {

	fmt.Println(fibonacci(0, 2))
	/*
		nextOdd := makeOddGenerator()
		fmt.Println(nextOdd()) // 1
		fmt.Println(nextOdd()) // 3
		fmt.Println(nextOdd()) // 5

		x := []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}
		for _, v := range x {
			fmt.Printf(strconv.Itoa(v) + "/2= ")
			fmt.Println(half((v)))
		}
		fmt.Println(greatest(x...))
	*/
}
func fibonacci(x int, y int, z int) func() int {
	return func() int {
		z, x, y = x, y, x+y
		return z
	}
// }

/*
The Fibonacci sequence is defined as:
fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Write a recursive function which can find fib(n).
*/
func fibonacci(startNum int, addCount int) int {
	currNum := startNum
	nextNum := currNum + 1
	result := 0
	for i := 0; i <= addCount; i++ {
		result = func(currNum int, nextNum int) int {
			return currNum + nextNum
		}
		currNum = nextNum
		nextNum = result
	}
	fmt.Println(result)
	return result
}

func avarage(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += float64(v)
	}
	return total / float64(len(xs))
}

func half(x int) (result int, isEven bool) {
	result = x / 2
	if x%2 == 0 {
		isEven = true
	}
	return result, isEven
}

func greatest(numbers ...int) int {
	theBest := 0
	for _, n := range numbers {
		if theBest < n {
			theBest = n
		}
	}
	return theBest
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
