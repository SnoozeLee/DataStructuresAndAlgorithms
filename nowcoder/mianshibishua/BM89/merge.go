package main

import "sort"

type Interval struct {
	Start int
	End   int
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func merge(intervals []*Interval) []*Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start <= intervals[j].Start
	})

	var res []*Interval
	t := intervals[0]
	for i := 1; i < len(intervals); i++ {
		canMerge, newInterval := mergeTwo(t, intervals[i])
		if canMerge {
			t = newInterval
		} else {
			res = append(res, t)
			t = intervals[i]
		}
	}
	return append(res, t)
}

// mergeTwo 是纯函数，不会修改外部变量
func mergeTwo(a, b *Interval) (canMerge bool, newInterval *Interval) {
	//	按左边界排序，确保a的左边界[不大于b的左边界[
	if b.Start < a.Start {
		t := a
		a = b
		b = t
	}

	if a.End >= b.Start {
		//	实时一定可以合并
		if a.End > b.End {
			return true, &Interval{Start: a.Start, End: a.End}
		} else {
			return true, &Interval{Start: a.Start, End: b.End}
		}
	}
	return false, nil
}
