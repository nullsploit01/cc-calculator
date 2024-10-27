package internal

import (
	"strconv"
	"strings"
	"unicode"
)

type Calculator struct {
	Expression string
	Tokens     []string
}

func NewCalculator(expression string) *Calculator {
	calculator := &Calculator{
		Expression: expression,
	}
	calculator.Tokenize()
	return calculator
}

func (c *Calculator) Tokenize() {
	var currentToken strings.Builder

	for _, char := range c.Expression {
		if unicode.IsDigit(char) || char == '.' { // Collects number characters and decimal points
			currentToken.WriteRune(char)
		} else if unicode.IsSpace(char) { // Handles spaces
			if currentToken.Len() > 0 {
				c.Tokens = append(c.Tokens, currentToken.String())
				currentToken.Reset()
			}
		} else { // Handles operators and parentheses
			if currentToken.Len() > 0 {
				c.Tokens = append(c.Tokens, currentToken.String())
				currentToken.Reset()
			}
			c.Tokens = append(c.Tokens, string(char))
		}
	}

	if currentToken.Len() > 0 {
		c.Tokens = append(c.Tokens, currentToken.String())
	}
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1

	case '*', '/':
		return 2

	default:
		return 0
	}
}

func applyOp(a, b float64, op rune) float64 {

	switch op {
	case '+':
		return a + b

	case '-':
		return a - b

	case '*':
		return a * b

	case '/':
		return a / b

	default:
		return 0
	}
}

func (c *Calculator) Calculate() float64 {
	var ops []rune
	var values []float64

	process := func(p int) {
		for len(ops) > 0 {
			topOp := ops[len(ops)-1]
			if precedence(topOp) < p {
				break
			}

			ops = ops[:len(ops)-1]

			b := values[len(values)-1]
			values = values[:len(values)-1]
			a := values[len(values)-1]
			values[len(values)-1] = applyOp(a, b, topOp)
		}
	}

	for _, token := range c.Tokens {
		if token == "(" {
			ops = append(ops, '(')
		} else if token == ")" {
			process(0)
			ops = ops[:len(ops)-1]
		} else if op, err := strconv.ParseFloat(token, 64); err == nil {
			values = append(values, op)
		} else if op := rune(token[0]); precedence(op) > 0 {
			process(precedence(op))
			ops = append(ops, op)
		}
	}

	process(0)
	return values[0]
}
