package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 6, 3, 4, 8, 7, 6, 0}
	n := InversePairs(arr)
	fmt.Println(arr, "->", n)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
const m = 1000000007

type Res struct {
	v int
}

func (r *Res) add(n int) {
	r.v = (r.v + n%m) % m
}

func InversePairs(nums []int) int {
	res := &Res{v: 0}
	//	利用归并排序的方法做
	mergeCalc(nums, 0, len(nums)-1, res)
	return res.v
}

func mergeCalc(nums []int, L, R int, res *Res) {
	if L == R {
		return
	}
	mid := (R-L)/2 + L
	mergeCalc(nums, L, mid, res)
	mergeCalc(nums, mid+1, R, res)

	i := L
	j := mid + 1
	assistArr := make([]int, R-L+1)
	assistIdx := 0 //
	for i <= mid && j <= R {
		if nums[i] <= nums[j] {
			assistArr[assistIdx] = nums[i]
			assistIdx++
			i++
		} else {
			//	产生逆序对
			res.add(mid - i + 1)
			assistArr[assistIdx] = nums[j]
			assistIdx++
			j++
		}
	}
	for i <= mid {
		assistArr[assistIdx] = nums[i]
		assistIdx++
		i++
	}
	//for j <= R {
	//	assistArr[assistIdx] = nums[j]
	//	assistIdx++
	//	j++
	//}
	// 这一部分的拷贝因为是排好序了，所以可以不用做，下方拷贝assistArr右边界缩小即可
	i = L
	for assistIdx = 0; assistIdx < j-L; assistIdx++ {
		nums[i] = assistArr[assistIdx]
		i++
	}

	return
}
