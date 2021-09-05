package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("Luas lingkaran\t\t : %.2f \n", area)
	fmt.Printf("keliling lingkaran\t : %.2f\n", circumference)
}

// 	var names = []string{"John", "Wick"}
// 	printMessage("hallo", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)

// 	rand.Seed(time.Now().Unix())

// 	var randomValue int

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number : ", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number : ", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number : ", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = rand.Int()%(max-min+1) + min
// 	return value

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference

}
