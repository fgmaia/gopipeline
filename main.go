package main

import "fmt"

func main() {
	num := []int{2, 3, 4, 7, 1}
	//stage1
	dataChannel := sliceToChannel(num)
	//stage2
	finalChannel := square(dataChannel)
	//stage3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
