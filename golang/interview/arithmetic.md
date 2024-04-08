# 计算斐波那契数列的第n项
斐波那契数列是一个递归定义的数列，其中每个数字是前两个数字的和。例如，前几个数字是0、1、1、2、3、5、8、13、21等等。求第n项的值可以使用递归来解决。
# 遍历二叉树，判断是否相同
```go
type TreeNode struct {
	Val   interface{} // 使用接口类型以便存储各种可能的数据类型
	Left  *TreeNode
	Right *TreeNode
}
```
实现代码: [../coding/tree_compare.go](../coding/tree_compare.go)
```shell
go run ../coding/arch/tree_compare.go
```

# 字符串反转
实现一个函数，输入一个字符串，将其反转并返回。
```go
func reverseString(s string) string{
	
}
func main() {
	str := "你好，世界，我要反转字符串"
	reversed := reverseString(str)
	println(reversed)
}
```
实现代码: [../coding/reverse_string.go](../coding/reverse_string.go)
```shell
go run ../coding/reverse_string.go
```