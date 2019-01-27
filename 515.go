package main

import (
	"fmt"
)

func main() {
	vars := []int{1, 4, 6, 78, 8, 4, 2, 1, 99}
	res, err := max(vars...)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Printf("Max: %d\n", res)

	res, err = min(vars...)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Printf("Min: %d\n", res)

	var s = []string{"yes", "hello", "good", "day"}
	res1 := join(s...)
	fmt.Println(res1)

}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("at least one argument required")
	}
	var max int
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("at least one argument required")
	}

	var min int
	for k, v := range vals {
		if k == 0 {
			min = v
		}
		if v < min {
			min = v
		}
	}
	return min, nil
}

func join(vals ...string) string {
	var result string
	if len(vals) == 0 {
		return result
	}

	for _, v := range vals {
		result += v
	}
	return result

}
