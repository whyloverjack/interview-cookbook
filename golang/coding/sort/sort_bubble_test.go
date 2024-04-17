package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	bubbleSortAsc(values)
	fmt.Print("升序排序: ")
	fmt.Println(values)

	bubbleSortDesc(values)
	fmt.Print("逆序排序: ")
	fmt.Println(values)
}

/*
冒泡排序的原理是，这里以降序排序为例，对给定的数组进行多次遍历，每次均比较相邻的两个数，如果前一个比后一个大，则交换这两个数。经过第一次遍历之后，最大的数就在最右侧了；第二次遍历之后，第二大的数就在右数第二个位置了；以此类推，每一次循环比较最终都会有一个数排在它最终的位置上。
*/
func bubbleSortAsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubbleSortDesc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
