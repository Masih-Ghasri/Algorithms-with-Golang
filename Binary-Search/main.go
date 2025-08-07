package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == target {
			return mid
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	var input string
	var searchNumber int
	var numbers []int

	fmt.Printf("Enter search number: like this 1,2,3,4,5,6,7,8 ")
	fmt.Scan(&input)

	strNumbers := strings.Split(input, ",")
	for _, str := range strNumbers {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			fmt.Println("Can not convert string in to int", str)
			return
		}
		numbers = append(numbers, num)
	}

	fmt.Println("Enter target number: ")
	fmt.Scanln(&searchNumber)

	sort.Ints(numbers)

	index := BinarySearch(numbers, searchNumber)

	if index == -1 {
		fmt.Printf("Can not find number in arr %d \n", searchNumber)
	} else {
		fmt.Printf("in arr %d number %d finded\n", searchNumber, index)
		fmt.Println("sort of arr:", numbers)
	}

}
