package main

import "fmt"

func main() {

	var point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfecet!")
		default:
			fmt.Println("nice")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("try harder")
		}
	}

}
