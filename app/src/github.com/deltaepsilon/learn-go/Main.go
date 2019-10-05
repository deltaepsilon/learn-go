package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	i int    = 42
	j string = strconv.Itoa(i + 10)
)

type Animal struct {
	Name   string
	Origin string
}

func main() {

	fmt.Printf("%v, %T \n", i, i)

	// j = strconv.Itoa(i)

	fmt.Printf("%v, %T \n", j, j)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field)

	if true {
		fmt.Println("Hey all")
	}

	switchTest := 5
	switch {
	case switchTest < 10:
		fmt.Println("we are 10")
		fallthrough
	case switchTest < 11:
		fmt.Println("we are 11")
		fmt.Println("we are also 11")
		fallthrough
	case switchTest < 12:
		fmt.Println("we are 12")
	case switchTest > 12:
		fmt.Printf("We are %v", switchTest)
	}

	// l := 7
	// for {
	// 	fmt.Printf("NOOOO %v\n", l)
	// 	l++

	// 	if l > 10 {
	// 		break
	// 	}
	// }

Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i*j > 6 {
				break Loop
			}

			fmt.Println(i, j)

		}
	}

}
