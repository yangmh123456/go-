package main

import "fmt"

func findpn(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	var x int
	fmt.Scanf("%d", &x)
	if findpn(x) {
		fmt.Printf("%d 是素数。/n", x)
	} else {
		fmt.Printf("%d 不是素数。/n", x)
	}

}
