package main

import (
	"fmt"
)

const (
	initMinT = 15
	initMaxT = 30
)

func main() {
	var (
		n int
		k int
	)

	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid number of departments")

		return
	}

	for range n {
		_, err = fmt.Scan(&k)
		if err != nil || k < 0 {
			fmt.Println("Invalid number of employees")

			return
		}

		if !processDepartment(k) {
			return
		}
	}
}

func processDepartment(k int) bool {
	minTemp := initMinT
	maxTemp := initMaxT

	for j := range k {
		var (
			sign string
			temp int
		)

		_, err := fmt.Scan(&sign, &temp)
		if err != nil {
			fmt.Println("Invalid value")

			return false
		}

		switch sign {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		default:
			fmt.Println("Invalid character")

			return false
		}

		result := getOptimalTemp(minTemp, maxTemp, sign, j == 0)
		fmt.Println(result)
	}

	return true
}

func getOptimalTemp(minTemp, maxTemp int, sign string, isFirst bool) int {
	if minTemp > maxTemp {
		return -1
	}

	if isFirst {
		if sign == ">=" {
			return minTemp
		}

		return maxTemp
	}

	return minTemp
}
