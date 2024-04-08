package main

import (
	"fmt"
	"testing"
)

/**
 * slice 是一个引用类型，底层是一个数组，slice 本身只是一个引用，指向底层数组的某个位置，slice 本身不存储任何数据。
 * slice的扩容策略是：当 slice 的容量不足时，会自动扩容，扩容的策略是：如果原 slice 的长度小于 1024，则新 slice 的容量是原 slice 的 2 倍；如果原 slice 的长度大于等于 1024，则新 slice 的容量是原 slice 的 1.25 倍。
 */
func TestSlice(t *testing.T) {
	m1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	m2 := m1[3:7]
	fmt.Println(m2, cap(m2))

	m2 = append(m2, 12)
	fmt.Println(m1)
	fmt.Printf("m1:%d; m2:%d; \n", cap(m1), cap(m2))

	m1 = append(m1, 11)
	fmt.Println(cap(m1))
}

/**
 * new 和 make 都是用来分配内存的。
 * new针对的是值类型，返回的是指向该类型的指针，比如：int, string, array; 而make针对的是引用类型，返回的是该引用类型的实例, 比如：slice, map, channel。
 * new返回变量的地址，make返回变量本身。
 * new分配的空间被清零，make分配空间后会初始化。
 */
func TestNewAndMake(t *testing.T) {

}
