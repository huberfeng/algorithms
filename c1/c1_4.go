package c1

import "fmt"

// 从n个数中 取3个数
func Cn3(a []int) {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				fmt.Println(a[i], " ", a[j], " ", a[k])
			}
		}
	}
}

// 从n个数中取3个数，排列
func An3(a []int) {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i==k || i==j || j==k {
					continue
				}
				fmt.Println(a[i], " ", a[j], " ", a[k])
			}
		}
	}
}