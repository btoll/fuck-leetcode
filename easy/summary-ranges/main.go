package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	rangeStart := nums[0]
	rangeEnd := nums[0]
	res := []string{}
	appendTo := func() {
		if rangeStart < rangeEnd {
			res = append(res, fmt.Sprintf("%d->%d", rangeStart, rangeEnd))
		} else {
			res = append(res, strconv.Itoa(rangeStart))
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == rangeEnd {
			rangeEnd += 1
			continue
		}
		appendTo()
		rangeStart = nums[i]
		rangeEnd = nums[i]
	}

	appendTo()
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	//	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
	// ["0->2","4->5","7"]
}
