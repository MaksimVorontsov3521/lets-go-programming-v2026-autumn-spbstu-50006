package main

import (
	"fmt"
	"io"
	"os"
)

const (
	initMinT       = 15
	initMaxT       = 30
	signG          = ">="
	signL          = "<="
	impossibleTemp = -1
)

type TempRange struct {
	minTemp int
	maxTemp int
}

func NewTempRange() *TempRange {
	return &TempRange{
		minTemp: initMinT,
		maxTemp: initMaxT,
	}
}

func (tempRange *TempRange) Update(sign string, temp int) error {
	switch sign {
	case signG:
		if temp > tempRange.minTemp {
			tempRange.minTemp = temp
		}
	case signL:
		if temp < tempRange.maxTemp {
			tempRange.maxTemp = temp
		}
	default:
		return fmt.Errorf("invalid sign: %s", sign)
	}

	return nil
}

func (tempRange *TempRange) GetOptimal() int {
	if tempRange.minTemp > tempRange.maxTemp {
		return impossibleTemp
	}

	return tempRange.minTemp
}

func processDepartment(numEmployees int, out io.Writer) error {
	tempRange := NewTempRange()

	for i := 0; i < numEmployees; i++ {
		var (
			sign string
			temp int
		)

		_, err := fmt.Scan(&sign, &temp)
		if err != nil {
			return fmt.Errorf("invalid value: %w", err)
		}

		if err := tempRange.Update(sign, temp); err != nil {
			return err
		}

		fmt.Fprintln(out, tempRange.GetOptimal())
	}

	return nil
}

func main() {
	var numDepartments int

	_, err := fmt.Scan(&numDepartments)
	if err != nil || numDepartments <= 0 {
		fmt.Println("Invalid number of departments")
		return
	}

	for i := 0; i < numDepartments; i++ {
		var numEmployees int

		_, err = fmt.Scan(&numEmployees)
		if err != nil || numEmployees < 0 {
			fmt.Println("Invalid number of employees")
			return
		}

		if err := processDepartment(numEmployees, os.Stdout); err != nil {
			fmt.Println(err)
			return
		}
	}
}
