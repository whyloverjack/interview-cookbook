# 交替打印数字和字母

**问题描述**

使用两个 `goroutine` 交替打印序列，一个 `goroutine` 打印数字， 另外一个 `goroutine` 打印字母， 最终效果如下：
```bash
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
```

[coding/print_number_letter.go](../coding/print_number_letter.go)
```shell
go run ../coding/print_number_letter.go
```

# 分配金币
* 你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。 
* 分配规则如下:  
  * 名字中每包含1个'e'或'E'分1枚金币 
  * 名字中每包含1个'i'或'I'分2枚金币 
  * 名字中每包含1个'o'或'O'分3枚金币 
  * 名字中每包含1个'u'或'U'分4枚金币 
  用go写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币, 补充distributeCoins函数?
```go
func distributeCoins(people []string, coins int) (map[string]int, int) {
	
}
func main() {
	people := []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	coins := 50

	distribution, remainingCoins := distributeCoins(people, coins)

	fmt.Println("Coin Distribution:")
	for person, coins := range distribution {
		fmt.Printf("%s: %d coins\n", person, coins)
	}
	fmt.Printf("Remaining Coins: %d\n", remainingCoins)
}
```
[coding/distribute_coins.go](../coding/distribute_coins.go)
```shell
go run ../coding/distribute_coins.go
```