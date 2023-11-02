package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func calculate(expression string) int {
	result := 0
	operator := '+'
	num := 0

	for i, char := range expression {
		if char >= '0' && char <= '9' {
			digit, _ := strconv.Atoi(string(char))
			num = num*10 + digit
		}

		if i == len(expression)-1 || (char == '+' || char == '-') {
			switch operator {
			case '+':
				result += num
			case '-':
				result -= num
			}

			operator = char
			num = 0
		}
	}

	return result
}

func generateExpression() string {
	expression := ""
	operatorCount := 10

	for i := 9; i >= 0; i-- {
		expression = strconv.Itoa(i) + expression

		if operatorCount > 0 {
			operator := ""
			switch rand.Intn(3) {
			case 0:
				operator = "+"
			case 1:
				operator = "-"
			}

			expression = operator + expression
			operatorCount--
		}
	}

	return expression
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numSolutions := 5 // Желаемое количество вариантов

	for solutionCount := 0; solutionCount < numSolutions; {
		expression := generateExpression()
		result := calculate(expression)

		if result == 200 {
			fmt.Println("Выражение:", expression)
			fmt.Println("Результат:", result)
			solutionCount++
		}
	}
}
