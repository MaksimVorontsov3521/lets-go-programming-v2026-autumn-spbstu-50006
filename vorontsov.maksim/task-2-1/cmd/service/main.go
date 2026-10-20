package main

import (
	"fmt"
)

const (
	initMinT = 15
	initMaxT = 30
)

func main() {
	var n, k int
	var err error

	_, err = fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid number of departments")
		return
	}

	for i := 0; i < n; i++ {
		minT := initMinT
		maxT := initMaxT

		_, err = fmt.Scan(&k)
		if err != nil || k < 0 {
			fmt.Println("Invalid number of employees")
			return
		}

		for j := 0; j < k; j++ {
			var sign string
			var temp int

			_, err = fmt.Scan(&sign, &temp)
			if err != nil {
				fmt.Println("Invalid value")
				return
			}

			if sign == ">=" {
				if temp > minT {
					minT = temp
				}
			} else if sign == "<=" {
				if temp < maxT {
					maxT = temp
				}
			} else {
				fmt.Println("Invalid character")
				return
			}

			if minT > maxT {
				fmt.Println(-1)
			} else {
				if j == 0 {
					if sign == ">=" {
						fmt.Println(minT)
					} else {
						fmt.Println(maxT)
					}
				} else {
					fmt.Println(minT)
				}
			}
		}
	}
}
