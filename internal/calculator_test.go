package internal

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		expression string
		want       float64
	}{
		{"3 + 2 * 2", 7},
		{"3 + 2 * 2 - 1", 6},
		{"10 / 2", 5},
		{"2.5 * 4", 10},
		{"100 / 25", 4},
		{"5 + 2 * (3 - 1)", 9},
		{"(5 + 2) * 3", 21},
		{"50 / (5 * 2)", 5},
	}

	for _, tt := range tests {
		c := NewCalculator(tt.expression)
		got := c.Calculate()
		if got != tt.want {
			t.Errorf("Calculate(%q) = %v, want %v", tt.expression, got, tt.want)
		}
	}
}

func TestTokenize(t *testing.T) {
	expression := "3 + 4 * 2 / ( 1 - 5 )"
	c := NewCalculator(expression)
	want := []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")"}
	if len(c.Tokens) != len(want) {
		t.Fatalf("Expected %d tokens, got %d", len(want), len(c.Tokens))
	}
	for i, token := range want {
		if c.Tokens[i] != token {
			t.Errorf("Expected token %d to be %q, got %q", i, token, c.Tokens[i])
		}
	}
}
