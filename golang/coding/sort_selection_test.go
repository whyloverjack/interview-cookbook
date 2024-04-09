package main

import (
	"fmt"
	"testing"
)

/**
选择排序
选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是：第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，然后再从剩余的未排序元素中寻找到最小（大）元素，然后放到已排序的序列的末尾。以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法。
*/
// 升序排序  写法1
func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0 // 最大值的索引，开始假设，第一个数最大
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ { // 每一次遍历，都会有一个数值在它最终的位置上
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
		fmt.Println(nums) // 方便看到每一次排序后的结果
	}
}

// 升序排序 写法2
func selectSort2(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		min := i
		//寻找最小的一个数，保存索引值
		for j := i + 1; j < length; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
		fmt.Println(nums) // 方便看到每一次排序后的结果
	}
}

// 降序排序 写法1
func selectReSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := i
		//寻找最大的一个数，保存索引值
		for j := i + 1; j < length; j++ { // 每一次遍历，都会有一个数值在它最终的位置上
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		fmt.Println(nums) // 方便看到每一次排序后的结果
	}
}

func TestSortSelection(t *testing.T) {
	nums := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	selectSort(nums)
	fmt.Println(nums)
}
