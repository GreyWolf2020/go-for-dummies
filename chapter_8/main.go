package main

import (
	"fmt"
	"sort"
)

type kv struct {
	key   string
	value int
}

type dob struct {
	dayOfWeek int
	month     int
	year      int
}

type people struct {
	name  string
	email string
	dob   dob
}

type kvPairs []kv

func (p kvPairs) Len() int {
	return len(p)
}

func (p kvPairs) Less(i, j int) bool {
	return p[i].value < p[j].value
}

func (p kvPairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var members map[int]people
	var keys []string
	var memmbersKeys []int
	var weights map[string]int
	var heights map[string]int
	heights = make(map[string]int)
	height := map[string]int{
		"Peter": 170,
		"Joan":  168,
		"Jan":   175,
	}
	members = make(map[int]people)
	members[1] = people{
		name:  "Mary Smith",
		email: "marysmith@example.com",
		dob: dob{
			dayOfWeek: 17,
			month:     3,
			year:      1990,
		},
	}
	members[2] = people{
		name:  "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
			dayOfWeek: 9,
			month:     12,
			year:      1988,
		},
	}
	members[3] = people{
		name:  "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
			dayOfWeek: 1,
			month:     12,
			year:      1988,
		},
	}
	members[4] = people{
		name:  "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
			dayOfWeek: 19,
			month:     8,
			year:      2001,
		},
	}

	heights["Peter"] = 170
	heights["Joan"] = 168
	heights["Jan"] = 175

	fmt.Println(heights["Peter"])
	fmt.Println(heights["Joan"])
	fmt.Println(heights["Jan"])

	fmt.Println(height["Jim"])
	if v, ok := height["Jim"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}
	if _, ok := heights["Joan"]; ok {
		delete(heights, "Joan")
	} else {
		fmt.Println("Key does not exists")
	}

	fmt.Println(len(heights))
	fmt.Println(len(weights))

	for k, v := range height {
		fmt.Println(k, v)
	}

	for k := range height {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Println(k, height[k])
	}

	p := make(kvPairs, len(height))
	i := 0
	for k, v := range height {
		p[i] = kv{k, v}
		i++
	}
	fmt.Println(p)
	sort.Sort(p)
	fmt.Println(p)
	for _, v := range p {
		fmt.Println(v)
	}

	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	}

	for k := range members {
		memmbersKeys = append(memmbersKeys, k)
	}

	sort.Ints(memmbersKeys)

	var sliceMembers []people
	for _, k := range memmbersKeys {
		sliceMembers = append(sliceMembers, members[k])
	}
	fmt.Println(sliceMembers)

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		// compare year
		if sliceMembers[i].dob.year != sliceMembers[j].dob.year {
			return sliceMembers[i].dob.year < sliceMembers[j].dob.year
		}

		// compare month
		if sliceMembers[i].dob.month != sliceMembers[j].dob.month {
			return sliceMembers[i].dob.month < sliceMembers[j].dob.month
		}

		// compare day
		return sliceMembers[i].dob.dayOfWeek < sliceMembers[j].dob.dayOfWeek

	})
	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}
}
