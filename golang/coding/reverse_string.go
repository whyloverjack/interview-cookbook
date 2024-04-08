package main

func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	str := "你好，世界，我要反转字符串"
	reversed := reverseString(str)
	println(reversed)
}
