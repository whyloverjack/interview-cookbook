package main

import "fmt"

func main() {
	number := make(chan bool)
	letter := make(chan bool, 1)
	done := make(chan struct{})

	go func() {
		for i := 1; i < 11; i += 2 {
			<-letter
			fmt.Print(i)
			fmt.Print(i + 1)

			number <- true
		}
	}()

	go func() {
		char_seq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-number
			fmt.Print(char_seq[i])
			fmt.Print(char_seq[i+1])
			letter <- true
		}
		done <- struct{}{}
	}()

	letter <- true
	<-done
}
