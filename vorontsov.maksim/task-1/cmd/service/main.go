package main

import (
	"errors"
	"fmt"
)

func readFloat() (int, error) {
	var a int

	_, err := fmt.Scan(&a)
	if err != nil {
		return 0, err
	}

	return a, nil
}

func readOperand() (string, error) {
	var operand string

	_, err := fmt.Scan(&operand)
	if err != nil {
		return "", errors.New("Invalid operation")
	}

	switch operand {
	case "+", "-", "*", "/":
		return operand, nil
	default:
		return "", errors.New("Invalid operation")
	}
}

func add(a, b int) (int, error) {
	return a + b, nil
}

func subtract(a, b int) (int, error) {
	return a - b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

func divide(a, b int) (int, error) {
	if b == 0.0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func calculate(a, b int, operand string) (int, error) {
	var err error

	switch operand {
	case "+":
		return add(a, b)
	case "-":
		return subtract(a, b)
	case "*":
		return multiply(a, b)
	case "/":
		return divide(a, b)
	default:
		err = errors.New("Invalid operation")
	}

	return 0, err
}

func main() {
	var a, b int
	var operation string
	var err error

	a, err = readFloat()
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	b, err = readFloat()
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	operation, err = readOperand()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result int
	result, err = calculate(a, b, operation)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
