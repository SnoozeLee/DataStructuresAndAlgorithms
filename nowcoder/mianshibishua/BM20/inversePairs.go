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
	res := Res{v: 0}
	//	利用插入排序的方法做：
	if len(nums) <= 1 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		t := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if t < nums[j] {
				res.add(1)
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = t
	}
	return res.v
}

//5/6 组用例通过
//您的程序未能在规定时间内运行结束，请检查是否循环有错或算法复杂度过大。
//您可以用fmt.Println在函数中打印信息分析问题
//基本证明思路没错，现在换过复杂度第一点的排序算法就可以了
