package main

import (
	"fmt"
	"strings"
)

func distributeCoins(people []string, coins int) (map[string]int, int) {
	distribution := make(map[string]int)
	for _, name := range people {
		var countE, countI, countO, countU int
		for _, char := range strings.ToLower(name) {
			switch char {
			case 'e':
				countE++
			case 'i':
				countI++
			case 'o':
				countO++
			case 'u':
				countU++
			}
		}
		distribution[name] = countE + 2*countI + 3*countO + 4*countU
	}

	var remainingCoins int
	for _, amount := range distribution {
		if coins < amount {
			fmt.Printf("Insufficient coins to distribute to all users.\n")
			return distribution, 0
		}
		coins -= amount
		remainingCoins = coins
	}

	return distribution, remainingCoins
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
