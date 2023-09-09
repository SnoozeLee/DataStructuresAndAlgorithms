package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 最少货币数
 * @param arr int整型一维数组 the array
 * @param aim int整型 the target
 * @return int整型
 */
func minMoney(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	aimMap := make(map[int]bool)
	count := 0
	return giveChange(arr, []int{aim}, count, &aimMap)
}

func giveChange(faceValues []int, aimList []int, count int, aimMap *map[int]bool) int {
	if len(aimList) == 0 {
		return -1
	}
	newAimList := []int{}
	for i := 0; i < len(aimList); i++ {
		if aimList[i] == 0 {
			return count
		} else {
			for j := 0; j < len(faceValues); j++ {
				newAim := aimList[i] - faceValues[j]
				if newAim >= 0 && !(*aimMap)[newAim] {
					(*aimMap)[newAim] = true
					newAimList = append(newAimList, newAim)
				}
			}
		}
	}
	return giveChange(faceValues, newAimList, count+1, aimMap)
}
