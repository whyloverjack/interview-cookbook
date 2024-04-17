package sort

import (
	"fmt"
	"testing"
)

/*
*
快速排序算法
基本思想: 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

快速排序算法通过多次比较和交换来实现排序，其排序流程如下：
1. 首先设定一个分界值，通过该分界值将数组分成左右两部分。
2. 将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。此时，左边部分中各元素都小于或等于分界值，而右边部分中各元素都大于或等于分界值。
3. 然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理。
4. 重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。当左、右两个部分各数据排序完成后，整个数组的排序也就完成了。
*/
func quickSort(values []int, left, right int) {
	temp := values[left] // temp 临时存放基数
	p := left            // p 就是 基数位置的索引
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp { //从右向左找第一个 小于 基数 的值
			j--
		}
		if j >= p {
			values[p] = values[j] // 小于基数的数，赋值给 基数位置
			p = j
		}

		for i <= p && values[i] <= temp { //从左向右 找第一个 大于基数的值
			i++
		}
		if i <= p {
			values[p] = values[i] // 大于基数的值，赋值给基数位置
			p = i
		}
	}
	values[p] = temp // 一趟遍历找到基数的最终位置，基数将整个数组一份为二，左右两数组一次递归
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

// 第2种写法
func quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1 //把第一个数值作为 基数值
	head, tail := 0, len(values)-1
	for head < tail {
		//fmt.Println(values)
		if values[i] > mid { //  比基数值大，交换位置，让较大值放在 右侧
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else { // 比基数值小,交换位置，叫较小值放在左侧
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid // 当 head == tail时，为基数找到它的最终位置，将切片分为左右两切片，递归进行
	quick2Sort(values[:head])
	quick2Sort(values[head+1:])
}

func TestQuickSort(t *testing.T) {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	quickSort(values, 0, len(values)-1)
	fmt.Println(values)

	values2 := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	quick2Sort(values2)
	fmt.Println(values2)
}
