package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var table [5][6]string
	var cube [4][3][3]string
	s := make([]int, 5)
	r := make([]int, 2, 5)
	t := []int{1, 2, 3, 4, 5}
	nums := [5]int{1, 2, 3, 4, 5} // an array of int (5 elements)
	var c [3]string
	c[0] = "iOS"
	c[1] = "Android"
	c[2] = "Windows"
	nums2 := [5]int{1, 2, 3, 4, 5}
	nums3 := nums2

	fmt.Println(nums)

	nums1 := [...]int{0, 2, 4, 6}
	fmt.Println(len(nums1))

	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] =
				strconv.Itoa(row) + "," + strconv.Itoa(column)
		}
	}
	fmt.Println(table)

	for row := 0; row < 4; row++ {
		for column := 0; column < 3; column++ {
			for depth := 0; depth < 3; depth++ {
				cube[row][column][depth] =
					strconv.Itoa(row) +
						strconv.Itoa(column) +
						strconv.Itoa(depth)
			}
		}
	}
	fmt.Println(cube)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println(r)
	fmt.Println(len(r))
	fmt.Println(cap(r))

	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))
	t = append(t, 6, 7, 8)
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))
	t = append(t, 9, 10)
	fmt.Println(len(t))
	fmt.Println(cap(t))
	u := t
	fmt.Println(u)
	fmt.Println(t)
	u[9] = 100
	fmt.Println(u)
	fmt.Println(t)
	t = append(t, 11)
	fmt.Println(u)
	fmt.Println(t)
	fmt.Println(cap(u))
	fmt.Println(cap(t))

	fmt.Println(c[0:2])
	fmt.Println(c[:])
	fmt.Println(c[1:])
	fmt.Println(c[:1])

	g := []int{1, 2, 3, 4, 5}
	g = g[2:4]
	fmt.Println(g)
	fmt.Println(cap(g))
	fmt.Println(len(g))

	fmt.Println(nums2)
	fmt.Println(nums3)
	nums3[0] = 0
	fmt.Println(nums2)
	fmt.Println(nums3)

	t = []int{1, 2, 3, 4, 5}
	v := make([]int, len(t))
	copy(v, t)
	fmt.Println(t)
	fmt.Println(v)

	w := make([]int, 2, 5)
	copy(w, t)
	fmt.Println(t)
	fmt.Println(w)
	fmt.Println(cap(w))
	fmt.Println(len(w))

	t = []int{1, 2, 3, 4, 5}
	t, err := insert(t, 2, 9)
	if err == nil {
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}
	t, error := delete(t, 2)
	if error == nil {
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}
}

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New(
			"Index cannot be less than 0")
	}
	if index >= len(orig) {
		return append(orig, value), nil
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}

func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index >= len(orig) {
		return nil, errors.New("Index out of range")
	}
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}
