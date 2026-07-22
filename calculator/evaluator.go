package calculator

import (
	"strconv"
	"unicode"
	//"github.com/MariamA-Helal/CLI-Calculator/calculator/operations.go"
)

func tokenize(input string) []string {
	var tokens []string
	var currentNum string

	for _, char := range input {

		if unicode.IsSpace(char) {
			continue
		}

		if unicode.IsDigit(char) || char == '.' {
			currentNum += string(char)
		} else {
			if currentNum != "" {
				tokens = append(tokens, currentNum)
				currentNum = ""
			}
			tokens = append(tokens, string(char))
		}

	}

	if currentNum != "" {
		tokens = append(tokens, currentNum)
	}

	return tokens
}

// Recursive Descent Parsing Algorith

type parser struct {
	tokens []string
	pos    int
}

func (p *parser) getNumber() float64 {
	val, _ := strconv.ParseFloat(p.tokens[p.pos], 64)
	p.pos++
	return val
}

func (p *parser) addSub() float64 {
	a := p.mulDiv()

	for p.pos < len(p.tokens) && (p.tokens[p.pos] == "+" || p.tokens[p.pos] == "-") {
		operator := p.tokens[p.pos]
		p.pos++

		b := p.mulDiv()

		switch operator {
		case "+":
			a = Add(a, b)
		case "-":
			a = Sub(a, b)
		}
	}
	return a
}

func (p *parser) mulDiv() float64 {
	a := p.pow()

	for p.pos < len(p.tokens) && (p.tokens[p.pos] == "*" || p.tokens[p.pos] == "/") {
		operator := p.tokens[p.pos]
		p.pos++

		b := p.pow()

		switch operator {
		case "*":
			a = Multi(a, b)
		case "/":
			decAns, _, _ := Div(a, b)
			a = decAns
		}
	}
	return a
}

func (p *parser) pow() float64 {
	a := p.getNumber()

	if p.pos < len(p.tokens) && p.tokens[p.pos] == "^" {
		p.pos++
		b := p.getNumber()

		return Power(a, b)
	}

	return a

}

func Evaluate(input string) float64 {
	tokens := tokenize(input)

	p := parser{tokens: tokens, pos: 0}

	return p.addSub()

}
