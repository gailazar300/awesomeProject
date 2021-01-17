package main

import "fmt"

func main() {
	n := 1
	ans, ok := fibonacci(n)
	if !ok {
		fmt.Printf("small number")
	} else {
		fmt.Printf("n is %v ans is %v\n", n, ans)
	}

}
func fibonacci(n int) (ans int, error bool) {
	var arr []int = make([]int, int(n+1))
	if n == 0 || n == 1 {
		return n, false
	}
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n], true
}
