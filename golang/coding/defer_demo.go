package main

func calc(index string, a, b int) int {
	ret := a + b
	println(index, a, b, ret)
	return ret
}

/**
 * defer 语句的执行顺序是后进先出
 * index: 1最后执行，但是index:1的第三个参数是一个函数，所以会先执行calc("10", a, b)，也就是输出10 1 2 3
 * index: 2倒数第二个执行，但是index:2的第三个参数是一个函数，所以会先执行calc("20", a, b)，也就是输出20 0 2 2
 * index: 2在执行完calc("20", a, b)之后，实际执行的是calc("2", a, 2)，此时a=0, b=2, 输出2 0 2 2
 * index: 1在执行完calc("10", a, b)之后，实际执行的是calc("1", a, 3)，此时a=1, b=3, 输出1 1 3 4
 * 输出结果：
 * 10 1 2 3
 * 20 0 2 2
 * 2 0 2 2
 * 1 1 3 4
 */
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
