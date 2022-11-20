package main

import (
	"fmt"
	"strconv"
	_ "strconv"
)

func main() {
	fmt.Println("Enter a list")
	var input string
	fmt.Scanln(&input)
	var num int
	fmt.Scanln(&num)
	slice := convertStringToSlice(input)
	answer := countValue(slice, num)
	fmt.Println(answer)

}

func convertStringToSlice(input string) (slice []int) {
	slice = []int{}
	for i := 0; i < len(input); i++ {
		value := string(input[i])
		if value != "," {
			value1, _ := strconv.Atoi(value)
			slice = append(slice, value1)
		}
	}
	return
}
func countValue(slice []int, value int) string {
	var ret string
	for i1, j := range slice {
		for i2, h := range slice {
			if j+h == value {
				ret = strconv.Itoa((i2)) + "," + strconv.Itoa(i1)
			}
		}
	}
	return ret
}
