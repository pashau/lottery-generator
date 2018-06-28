package lotto

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func contains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func uniquerandnumbers(amount int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	numberSlice := make([]int, amount)
	for i := 0; i <= amount-1; i++ {
		var random int = rand.Intn(max)
		if contains(numberSlice, random) {
			i--
			continue
		}
		numberSlice[i] = random
	}
	sort.Ints(numberSlice)
	return numberSlice
}

func inputInt(inputTxt string) int {
	var x int
	fmt.Print(inputTxt)
	// blank identifier?
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	return x
}

/*
func main() {
	var fields int = inputInt("Amount of Fields? ")
	for i := 1; i <= fields; i++ {
		print(i, "#: ")
		fmt.Print(uniquerandnumbers(5, 50))
		fmt.Println(uniquerandnumbers(2, 10))
	}
}
*/
