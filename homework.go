package homework01

import (
	"fmt"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement
	allnums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		count := allnums[num]
		if count != 0 {
			allnums[num]++
		} else {
			allnums[num] = 1
		}
	}
	for k, count := range allnums {
		if count == 1 {
			return k
		} else {
			fmt.Printf("数字 %d 出现了: %d\n", k, count)
		}
	}
	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// TODO: implement
	intStr := fmt.Sprintf("%d", x)

	left := 1
	right := len(intStr)
	for left < right {
		if intStr[left-1:left] != intStr[right-1:right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// TODO: implement
	if s == "" {
		return false
	}
	a := []rune{}

	for _, str := range s {
		if str == '(' || str == '{' || str == '[' {
			a = append(a, str)
		} else {
			b := a[len(a)-1]
			if (str == ')' && b != '(') || (str == '}' && b != '{') || (str == ']' && b != '[') {
				return false
			}
			a = a[:len(a)-1]
		}
	}

	return len(a) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	if len(strs) == 0 {
		return ""
	}
	one := strs[0]
	for i := 0; i < len(one); i++ {
		str := one[i]
		for j := 1; j < len(strs); j++ {
			strj := strs[j][i]
			if str != strj || i >= len(strs[j]) {
				return one[:i]
			}
		}
	}
	return one
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] == 9 && i != 0 {
			digits[i] = 0
			digits[i-1]++
		} else if digits[i] == 9 && i == 0 {
			digits[i] = 0
			return append([]int{1}, digits...)
		} else {
			digits[i]++
			return digits
		}
	}
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
	if len(nums) == 0 {
		return 0
	}
	s := 0
	for i := 0; i < len(nums); i++ {
		if nums[s] != nums[i] {
			s++
			nums[s] = nums[i]
		}
	}
	return s + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	if len(intervals) == 0 {
		return [][]int{}
	}
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals)-i; j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	r := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= r[len(r)-1][1] {
			if intervals[i][1] > r[len(r)-1][1] {
				r[len(r)-1][1] = intervals[i][1]
			}
		} else {
			r = append(r, intervals[i])
		}
	}
	return r
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	a := 0
	b := 0

	for i := 0; i < len(nums); i++ {
		d := target - nums[i]
		if d > 0 {
			a = i
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == d {
					b = j
					return []int{a, b}
				}
			}
		}

	}
	return nil
}
