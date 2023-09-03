package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	jumpBefore1 := 2 // 往前一级，也就是2级台阶，有2种方案
	jumpBefore2 := 1 // 往前两级，也就是1级台阶，有1中方案
	var jumpNow int
	for i := 3; i <= number; i++ {
		jumpNow = jumpBefore2 + jumpBefore1
		jumpBefore2 = jumpBefore1
		jumpBefore1 = jumpNow
	}
	return jumpNow
}
