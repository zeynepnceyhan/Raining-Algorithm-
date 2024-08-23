package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 5, 3, 3, 2, 2, 1, 1}
	fmt.Println("Volume of water: ", raining(arr))
}

func raining(arr []int) int {
	n := len(arr)

	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// Calculate leftMax array
	leftMax[0] = arr[0]
	for i := 1; i < n; i++ { // Loop should run until n-1
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(arr[i])))
	}

	// Calculate rightMax array
	rightMax[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(arr[i])))
	}

	// Calculate the total water trapped
	totalWater := 0
	for i := 0; i < n; i++ {
		levelOfWater := int(math.Min(float64(leftMax[i]), float64(rightMax[i])))
		if levelOfWater > arr[i] {
			totalWater += levelOfWater - arr[i]
		}
	}
	return totalWater
}
