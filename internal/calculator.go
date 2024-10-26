package internal

import (
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
		if unicode.IsDigit(char) || char == '.' {
			currentToken.WriteRune(char)
		} else if unicode.IsSpace(char) {
			if currentToken.Len() > 0 {
				c.Tokens = append(c.Tokens, currentToken.String())
				currentToken.Reset()
			} else {
				if currentToken.Len() > 0 {
					c.Tokens = append(c.Tokens, currentToken.String())
					currentToken.Reset()
				}
				c.Tokens = append(c.Tokens, string(char))
			}
		}
	}

	if currentToken.Len() > 0 {
		c.Tokens = append(c.Tokens, currentToken.String())
	}
}
