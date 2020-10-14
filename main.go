package main

import "fmt"

func main() {
	fmt.Println("Done")
}

func Loop1(times int) {
	for i := 0; i < 10; i++ {
		fmt.Sprintf("%d", i)
	}
}

func Loop2(times int) {
	for i := 0; i < 10; i += 2 {
		fmt.Sprintf("%d", i)
		fmt.Sprintf("%d", i)
	}
}

func Loop5(times int) {
	for i := 0; i < 10; i += 5 {
		fmt.Sprintf("%d", i)
		fmt.Sprintf("%d", i)
		fmt.Sprintf("%d", i)
		fmt.Sprintf("%d", i)
		fmt.Sprintf("%d", i)
	}
}
