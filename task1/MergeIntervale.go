package task1

import (
	"sort"
)

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)

	// 添加第一个区间
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := merged[len(merged)-1]

		// 检查是否有重叠
		if current[0] <= lastMerged[1] {
			// 有重叠，合并（取最大的结束点）
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// 没有重叠，添加新区间
			merged = append(merged, current)
		}
	}

	return merged
}
