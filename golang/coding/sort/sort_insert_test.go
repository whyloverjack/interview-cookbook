package sort

import (
	"fmt"
	"testing"
)

/*
*
插入排序
将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
*/
func TestSortInsert(t *testing.T) {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	//start := time.Now().UnixNano()
	fmt.Print("升序排序:")
	insertSort(values)
	fmt.Println(values)

	fmt.Print("逆序排序:")
	insertReSort(values)
	//end := time.Now().UnixNano()
	fmt.Println("最终排序:")
	//usetime := end - start
	fmt.Println(values)
	//fmt.Println("耗时:",usetime)
}

// 升序排序
func insertSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ { //从第二个开始，在有序队列中寻找合适位置插入
		temp := arr[i] // 记录要插入的值
		//fmt.Println("插入值:",temp)
		j := i                         //从已经排序的序列最右边的开始比较，找到比其小的数
		for j > 0 && temp < arr[j-1] { // 在左侧有序队列从右向左一次比较，找到合适的位置
			arr[j] = arr[j-1]
			j--
		}
		// 存在比其小的数，插入
		if j != i {
			arr[j] = temp
		}
		fmt.Println(arr)
	}
}

// 逆序排序
func insertReSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ { //从第二个开始，在有序队列中寻找合适位置插入
		temp := arr[i] // 记录要插入的值
		//fmt.Println("插入值:",temp)
		j := i                         //从已经排序的序列最右边的开始比较，找到比其小的数
		for j > 0 && temp > arr[j-1] { // 在左侧有序队列从右向左一次比较，找到合适的位置
			arr[j] = arr[j-1]
			j--
			//fmt.Println(arr)
		}
		// 找到合适位置，插入
		if j != i {
			arr[j] = temp
		}
		fmt.Println(arr)
	}
}
