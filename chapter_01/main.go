package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)

const publisher = "Wiley"

func main() {
	var num1 = 5 // type is inferred
	var num2 int
	var num3 float32        // floating point variable
	var name string         // string variable
	var raining bool        // boolean variable
	var rates float32 = 4.5 // declare and then initialise
	firstName, lastName, age := "Wei-Meng", "Lee", 25
	address := "The White House\n1600 Pennsylvania Avenue NW\nWashington, DC 20500"
	quotation := `"Anyone who has never made
a mistake has never tried
anything new."
--Albert Einstein`
	str1 := "你好,世界"
	str2 := "こんにちは世界"
	str3 := "\u4f60\u597d\uff0c\u4e16\u754c" // 你好,世界
	start := time.Now()
	var input string

	fmt.Println(num1)    // 5
	fmt.Println(num2)    // 0
	fmt.Println(num3)    // 0
	fmt.Print(name)      // "" (empty string)
	fmt.Println(raining) // false
	fmt.Println(rates)
	fmt.Println("The author of the book that am coding to is called", firstName, lastName, "and he is aged", age)
	fmt.Println("The publisher of his book is", publisher)
	fmt.Println(address)
	fmt.Println(quotation)
	fmt.Println(utf8.RuneCountInString(str1))  // 5
	fmt.Println(utf8.RuneCountInString(str2))  // 7
	fmt.Println(utf8.RuneCountInString(str3))  // 5
	fmt.Println(len(str1))                     // 15 = 5 chars * 3 bytes
	fmt.Println(len(str2))                     // 21 = 7 chars * 3 bytes
	fmt.Println(len(str3))                     // 15 = 5 chars * 3 bytes
	fmt.Println(reflect.TypeOf(start))         // time.Time
	fmt.Println(reflect.ValueOf(start).Kind()) // struct
	fmt.Println("Please enter your age: ")
	fmt.Scanf("%s", &input)
	age, err := strconv.Atoi(input) // convert string to int, Atoi means ASCII to integer
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}

	b, err := strconv.ParseBool("t")
	fmt.Println(b)        // true
	fmt.Println(err)      // <nill>
	fmt.Printf("%T\n", b) // bool

	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)   // 3.1415
	fmt.Println(err) // <nil>
	fmt.Printf("%T\n", f)

	i, err := strconv.ParseInt("-18.56", 10, 64) // base 10, 4-bit,
	fmt.Println(i)                               // 0
	fmt.Println(err)                             // strconv.ParseInt: parsing "-18.56": invalid syntax
	fmt.Printf("%T\n", i)                        // int64

	u1, err := strconv.ParseUint("18", 10, 64)
	fmt.Println(u1)  // 18
	fmt.Println(err) // <nil>
	fmt.Printf("%T\n", u1)

	u2, err := strconv.ParseUint("-18", 10, 64)
	fmt.Println(u2)  // 0
	fmt.Println(err) // strconv.ParseUint: parsing "-18": invalid syntax
	fmt.Printf("%T\n", u2)

	queue := 5
	name = "John"
	s := fmt.Sprintf("%s, your queue number is: %d", name, queue)
	fmt.Print(s)
}
