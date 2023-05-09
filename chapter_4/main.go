package main

import "fmt"

func main() {
	var OS [3]string
	OS[0] = "iOS"
	OS[1] = "Android"
	OS[2] = "Windows"

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	generateFibonacci()

	for i, v := range OS {
		fmt.Println(i, v)
	}

	for pos, char := range "Hello, world!" {
		fmt.Println(pos, char)
	}

	for pos, char := range "Hello, world!" {
		fmt.Printf("%d %c\n", pos, char)
	}

	for pos, char := range "こんにちは世界" {
		fmt.Printf("%c at byte location %d\n", char, pos)
	}

Outerloop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break Outerloop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("-------------------")
	}

AnotherOuterloop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				continue AnotherOuterloop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("-------------------")
	}

}

func generateFibonacci() {
	max := 100
	a, b := 0, 1
	for b <= max {
		println(b)
		a, b = b, a+b
	}
}
