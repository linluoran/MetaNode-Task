package main

import (
	"fmt"
	"slices"
	"strconv"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		fmt.Printf("只出现一次的数是: %d\n", nums[0])
		return nums[0]
	}

	countMaps := map[int]int{}
	for _, num := range nums {
		count, ok := countMaps[num]
		if ok {
			if count > 1 {
				continue
			} else {
				countMaps[num] += 1
			}
		} else {
			countMaps[num] = 1
		}
	}

	for num, count := range countMaps {
		if count == 1 {
			fmt.Printf("只出现一次的数是: %d\n", num)
			return num
		}
	}
	return 0
}

func isPalindrome(x int) bool {
	if x < 0 {
		fmt.Println("false")
		return false
	}

	if x < 10 {
		fmt.Println("true")
		return true
	}

	str := strconv.Itoa(x)
	strLen := len(str)
	for i := 0; i < strLen/2; i++ {
		if str[i] != str[strLen-i-1] {
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("true")
	return true
}

func isValid(s string) bool {

	sLen := len(s)

	if sLen%2 != 0 {
		fmt.Println("false")
		return false
	}

	sMap := map[uint8]uint8{
		40:  41,
		123: 125,
		91:  93,
	}

	if sLen == 2 {
		res := s[1] == sMap[s[0]]
		fmt.Println(res)
		return res
	}

	var sSlice []uint8
	for i := 0; i < sLen; i++ {
		if i == sLen-1 {
			sSlice = append(sSlice, s[i])
			break
		}
		if s[i+1] == sMap[s[i]] {
			i += 1
		} else {
			sSlice = append(sSlice, s[i])
		}
	}

	ssLen := len(sSlice)
	if sLen == ssLen {
		fmt.Println("false")
		return false
	}

	if ssLen == 0 {
		fmt.Println("true")
		return true
	}
	return isValid(string(sSlice))
}

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, str := range strs[1:] {
		if len(str) == 0 {
			return ""
		}
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	var resStr []byte
	for i := 0; i < minLen; i++ {
		tmpStr := strs[0][i]
		for _, str := range strs {
			if str[i] != tmpStr {
				return string(resStr)
			}

		}
		resStr = append(resStr, tmpStr)
	}
	return string(resStr)
}

func removeDuplicates(nums []int) int {
	tmp := nums[0]

	numsLen := len(nums)
	if numsLen == 1 {
		return len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] <= tmp {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			tmp = nums[i]
		}
	}
	return len(nums)
}

func plusOne(digits []int) []int {
	digitsLen := len(digits)

	addNum := 0
	for i := digitsLen - 1; i >= 0; i-- {
		if i == digitsLen-1 {
			if digits[i] == 9 {
				addNum += 1
				digits[i] = 0
			} else {
				digits[i] += 1
			}
		} else {
			if addNum != 0 {
				if digits[i] == 9 {
					addNum += 1
					digits[i] = 0
				} else {
					digits[i] += 1
					return digits
				}
			}
		}

		if i == 0 {
			if addNum != 0 {
				digits = append([]int{1}, digits[:]...)
			}
			return digits
		}
	}
	return digits
}

func merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
	start, end := intervals[0][0], intervals[0][1]
	resItem := [][]int{{start, end}}
	for _, Item := range intervals[1:] {
		tmpStart, tmpEnd := Item[0], Item[1]
		mergeFlag := false
		if tmpStart <= end {
			mergeFlag = true
			end = max(end, tmpEnd)
		}

		if !mergeFlag {
			start = tmpStart
			end = tmpEnd
			resItem = append(resItem, []int{start, end})
		} else {
			resItem[len(resItem)-1][1] = end
		}

	}
	return resItem
}

func twoSum(nums []int, target int) []int {
	countMap := make(map[int]int)
	for numIndex, num := range nums {
		tmpRes := target - num
		if index, ok := countMap[tmpRes]; ok {
			return []int{index, numIndex}
		} else {
			countMap[num] = numIndex
		}

	}
	return nil
}

func main() {
	goBase1 := []int{4, 1, 2, 1, 2}
	singleNumber(goBase1)

	goBase2 := 1231321
	isPalindrome(goBase2)

	goBase3 := "([)]"
	isValid(goBase3)

	goBase4 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(goBase4))

	goBase5 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(goBase5))

	goBase6 := []int{9, 9}
	fmt.Println(plusOne(goBase6))

	goBase7 := [][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}
	fmt.Println(merge(goBase7))

	goBase8 := []int{2, 7, 11, 15}
	goBaseTarget8 := 9
	fmt.Println(twoSum(goBase8, goBaseTarget8))
}
