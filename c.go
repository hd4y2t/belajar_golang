package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	//fungsi closure
// 	var getMinMAx = func(n []int) (int, int) {
// 		var min, max int
// 		for i, e := range n {
// 			if i == 0 {
// 				max, min = e, e
// 			} else if e > max {
// 				max = e
// 			} else if e < min {
// 				min = e
// 			}
// 		}
// 		return min, max
// 	}
// 	var numbers = []int{2, 3, 4, 3, 4, 2, 3, 10, 4, 55, 22, 53}
// 	var min, max = getMinMAx(numbers)
// 	fmt.Printf("data : %v \nmin : %v\nmax : %v\n", numbers, min, max)
// }

// func main() {
// 	//IIFE (immediately-invoked function expression)
// 	var number = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
// 	var newNumbers = func(min int) []int {
// 		var r []int
// 		for _, e := range number {
// 			if e < min {
// 				continue
// 			}
// 			r = append(r, e)
// 		}
// 		return r
// 	}(3)

// 	fmt.Println("nomor asli : ", number)
// 	fmt.Println("filter nomor : ", newNumbers)
// }

//closure sebagai nilai kembalian
// func main() {
// 	var max = 3
// 	var number = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
// 	var howMany, getNumbers = findMax(number, max)
// 	var theNumbers = getNumbers()

// 	fmt.Println("numbers\t :", number)
// 	fmt.Printf("find\t : %d\n\n", max)

// 	fmt.Println("found\t :", howMany)
// 	fmt.Println("value\t:", theNumbers)
// }
// func findMax(number []int, max int) (int, func() []int) {
// 	//closure sebagai nilai kembalian
// 	var res []int
// 	for _, e := range number {
// 		if e <= max {
// 			res = append(res, e)
// 		}
// 	}
// 	return len(res), func() []int {
// 		return res
// 	}
// }

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContains0 = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5

	})
	fmt.Println("data asli\t\t: ", data)
	//data asli

	fmt.Println("filter ada huruf \"o\"\t:", dataContains0)
	//filter ada huruf "o" : jason

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)

}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
