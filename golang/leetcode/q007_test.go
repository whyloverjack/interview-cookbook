package leetcode

import (
	"math"
	"strconv"
	"testing"
)

/*基于字符串反转思路实现*/
func reverseInt(x int) int {
	char := ""
	if x < 0 {
		char = "-"
		x = -x
	}
	str := strconv.Itoa(x)

	left, right := 0, len(str)-1
	runes := []rune(str)
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	result := char + string(runes)
	num, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}
	maxInt := math.MaxInt32
	minInt := math.MinInt32
	if num > maxInt || num < minInt {
		return 0
	}
	return num
}

/*更优*/
func reverse(x int) int {
	rev := 0
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev
}

func TestReverseInt(t *testing.T) {
	println(reverse(123))
	println(reverse(-123))
	println(reverse(120))
	println(reverse(0))
	println(reverse(1534236469))
}
