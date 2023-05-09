package main

import (
	"fmt"
)

func main() {
	num := 5
	condition := num%2 == 1
	if condition {
		fmt.Println("Number is odd")
	}

	if raining() || snowing() {
		fmt.Println("Stay in doors")
	}

	if !raining() || snowing() {
		fmt.Println("Let's ski!")
	}

	if !raining() && !snowing() {
		fmt.Println("Lets go outdoors")
	}

	if raining() && snowing() {
		fmt.Println("It's going tobe really cold!")
	}

	v, err := doSomething()
	if err {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println(v)
	}

	if w, er := doSomething(); !er {
		fmt.Println(w)
	} else {
		fmt.Println("Error occured: ", er)
	}

	parity := checkParity(num)
	fmt.Println(parity)

	dayOfTheWeek := ""
	switch num {
	case 1:
		dayOfTheWeek = "Monday"
		fmt.Print("Monday blues...")
	case 2:
		dayOfTheWeek = "Tuesday"
	case 3:
		dayOfTheWeek = "Wednesday"
	case 4:
		dayOfTheWeek = "Thursday"
	case 5:
		dayOfTheWeek = "Friday"
		fmt.Println("TGIF!!!")
	case 6:
		dayOfTheWeek = "Saturday"
	case 7:
		dayOfTheWeek = "Sunday"
	default:
		dayOfTheWeek = "--error--"
	}
	fmt.Println(dayOfTheWeek)
	grade := "C"
	switch grade {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "C":
		fallthrough
	case "D":
		fmt.Println("Passed")
	case "F":
		fmt.Print("Failed")
	default:
		fmt.Println("Absent")
	}
	anotherGrade := "D"
	switch anotherGrade {
	case "A", "B", "C", "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Undefined")
	}
	score := 79
	rimweGrade := ""
	switch {
	case score < 50:
		rimweGrade = "F"
	case score < 60:
		rimweGrade = "D"
	case score < 70:
		rimweGrade = "C"
	case score < 80:
		rimweGrade = "B"
	default:
		rimweGrade = "A"
	}
	fmt.Println(rimweGrade)
}

func raining() bool {
	fmt.Println("Check if it is raining now...")
	return true
}

func snowing() bool {
	fmt.Println("Check if it is snowing now...")
	return true
}

func doSomething() (int, bool) {
	return 5, false
}

func checkParity(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}
