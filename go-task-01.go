package main

import (
	"fmt"
	"sort"
	"strconv"
)

//只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
//找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
//例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

// go 对数器怎么编写 ？
// 1.循环数组，将数组的值作为key put 到map 中, value 为int 类型，记录map的次数
// 2.遍历map , get key , 比较value 的值，
// go 有没有 表达式使用， 例如Java的函数编程，直接调用lammd 表达式快速写代码

// 只出现一次的数字
func getOnlyOnceNum(nums []int) int {
	// 边届处理
	// for
	// map 塞值
	// for map if 比较 return
	var mapping map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if mapping[nums[i]] == 0 {
			mapping[nums[i]] = 1
		} else {
			a := mapping[nums[i]]
			mapping[nums[i]] = a + 1
		}
		if mapping[nums[i]] > 1 {
			return mapping[nums[i]]
		}
	}
	// 报错
	fmt.Println("数组中没有唯一的元素")
	return 0
}

/*
* 回文数
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// 转字符串
	s := strconv.Itoa(x)
	// 12321
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

/*
* 回文数
 */
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return original == reversed
}

/*
* 最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//["flower","flow","flight"]
	fmt.Println(strs)
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i] // f , l
		// 第一个字符串不用匹配直接和后面的匹配
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够或字符不匹配
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始加一
	for i := n - 1; i >= 0; i-- {
		// 当前位加一
		digits[i]++
		// 检查是否需要进位
		if digits[i] < 10 {
			return digits
		}
		// 需要进位，当前位置0
		digits[i] = 0
	}
	// 如果所有位都进位了，需要在最前面加一个1
	return append([]int{1}, digits...)
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int, k int) int {
	if len(nums) <= k {
		return len(nums)
	}

	i := k
	for j := k; j < len(nums); j++ {
		if nums[j] != nums[i-k] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按照区间起始点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// 检查是否有重叠
		if current[0] <= last[1] {
			// 合并区间，取最大的结束点
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

// 两数之和
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // 值到索引的映射

	for i, num := range nums {
		complement := target - num
		if idx, ok := numMap[complement]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil // 题目保证有解，这行不会执行
}
