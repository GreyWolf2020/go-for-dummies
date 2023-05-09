package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := 6
	a := []int{1, 2, 3, 4, 5}
	var cond func(int) bool
	var i func() int

	displayTime("Mon 2006-01-02 15:04:05", "Current Date and Time")

	swap(&x, &y)
	fmt.Println(x, y)

	s := addNum(5, 6)
	fmt.Println(s)

	o, e := countOddEven("12345678")
	fmt.Println(o, e)

	fmt.Println(addNums(1, 2, 3, 4, 5, 6))
	fmt.Println(addNums(5, 6, 7, 8, 9))

	i = doSomething
	fmt.Println(i())

	i = func() (j int) {
		j = 20
		return
	}
	fmt.Println(i())

	gen := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}

	cond = func(val int) bool {
		return val%2 == 0
	}
	evens := filter(a, cond)
	fmt.Println(evens)

	cond = func(val int) bool {
		return val > 3
	}
	gr8erThan3 := filter(a, cond)
	fmt.Println(gr8erThan3)

}

func displayTime(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func addNum(num1, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func countOddEven(s string) (int, int) {
	odds, evens := 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

func addNums(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}

func doSomething() int {
	return 5
}

func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}

func filter(arr []int, cond func(int) bool) (result []int) {
	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return
}
