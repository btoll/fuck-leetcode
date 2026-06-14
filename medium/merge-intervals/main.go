package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	res = append(res, append([]int(nil), intervals[0]...))
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := res[len(res)-1]
		// The current interval is contained in the last,
		// i.e., [[1 4] [2 3]].
		// This will include cases where both intervals
		// contain equal values.
		if cur[0] >= last[0] && cur[1] <= last[1] {
			continue
		}
		// If overlaps and extends the end, merge by extending last.
		if cur[0] <= last[1] && cur[1] > last[1] {
			last[1] = cur[1]
			res[len(res)-1] = last
			continue
		}
		// No overlap.
		res = append(res, append([]int(nil), cur...))
	}
	return res
}

//func sortIntervals(intervals [][]int) {
//	for i := range intervals {
//		for j := 1; j < len(intervals)-i; j++ {
//			if intervals[j-1][0] > intervals[j][0] {
//				intervals[j-1], intervals[j] = intervals[j], intervals[j-1]
//			}
//		}
//	}
//}

func main() {
	tests := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{4, 7}, {1, 4}},
		{{1, 4}, {2, 3}},
		{{1, 4}, {1, 4}},
		{{1, 4}, {0, 4}},
	}
	for _, test := range tests {
		fmt.Println(merge(test))
	}
}
