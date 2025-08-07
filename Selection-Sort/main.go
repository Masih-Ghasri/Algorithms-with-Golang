package main

import "fmt"

// findSmallest finds the index of the smallest element in the array
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

// selectionSort sorts an array from smallest to largest
func selectionSort(arr []int) []int {
	newArr := make([]int, 0, len(arr))
	copiedArr := make([]int, len(arr))
	copy(copiedArr, arr) // Create a copy to avoid modifying the original

	for len(copiedArr) > 0 {
		smallestIndex := findSmallest(copiedArr)
		newArr = append(newArr, copiedArr[smallestIndex])
		// Remove the smallest element by creating a new slice
		copiedArr = append(copiedArr[:smallestIndex], copiedArr[smallestIndex+1:]...)
	}
	return newArr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	sorted := selectionSort(arr)
	fmt.Println(sorted) // Output: [2 3 5 6 10]
}
