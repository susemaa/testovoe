package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, s int) int {
	var leftPointer = 0
	var rightPointer = len(arr) - 1
	for leftPointer <= rightPointer {
		var midPointer = int((leftPointer + rightPointer) / 2)
		var midValue = arr[midPointer]

		if midValue == s {
			return midPointer
		} else if midValue < s {
			leftPointer = midPointer + 1
		} else {
			rightPointer = midPointer - 1
		}
	}

	return -1
}

func range_count(ages []int, ranges [][2]int) []int {
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	})
	var count []int
	var x, y, k int
	for i := range ranges {
		k = 0
		for {
			x = binarySearch(ages, ranges[i][0]+k)
			if x != -1 {
				break
			}
			k++
		}
		k = 0
		if ranges[i][1] > ages[len(ages)-1] {
			y = len(ages) - 1
			count = append(count, y-x+1)
		} else {
			for {
				y = binarySearch(ages, ranges[i][1]+k)
				if y != -1 {
					break
				}
				k++
			}
			count = append(count, y-x)
		}
	}
	return count
}

func ezway(ages []int, ranges [][2]int) []int {
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	})
	var count []int
	var x int
	for i := range ranges {
		x = 0
		for j := range ages {
			if ages[j] >= ranges[i][0] {
				x++
			}
			if ages[j] > ranges[i][1] {
				x--
				break
			}
		}
		count = append(count, x)
	}
	return count
}

func main() {
	a := []int{32, 33, 18, 22, 27, 25}
	b := [][2]int{
		{20, 30},
		{30, 40},
		{10, 30},
		{20, 100},
	}
	ans1 := ezway(a, b)
	ans2 := range_count(a, b)
	fmt.Println(ans1, ans2)

}
