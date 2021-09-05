package main

import (
	"fmt"
	"strings"
)

func main() {
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)

}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ",")

	fmt.Printf("Hello, my name is : %s\n", name)
	fmt.Printf("my hobbies are : %s\n", hobbiesAsString)
}
