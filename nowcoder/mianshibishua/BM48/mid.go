package main

import "sort"

// 先来个速写版本

var arr []int
var arrIsOdd = false
var length = 0

func Insert(num int) {
	arr = append(arr, num)
	length++
	arrIsOdd = !arrIsOdd
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})
}

func GetMedian() float64 {
	if arrIsOdd {
		return float64(arr[length/2])
	}
	return (float64(arr[length/2-1]) + float64(arr[length/2])) / 2
}
