package expr

import (
  "fmt"
  "strconv"
  "time"
  "math/rand"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenOperator
	TokenOpenParen
	TokenCloseParen
	TokenEmpty
)

type Token struct {
	Type TokenType
	Value string
}

func tokenize(s string) []Token {
		tokens := []Token{};
		for _, char:= range s {
				switch char {
					case '+', '-', '*', '/':
						tokens = append(tokens, Token{Type: TokenOperator, Value: string(char)})
					case '(':
						tokens = append(tokens, Token{Type: TokenOpenParen, Value: string(char)})
					case ')':
							tokens = append(tokens, Token{Type: TokenCloseParen,Value:string(char)})
						default:
							if len(tokens) > 0 && tokens[len(tokens)-1].Type == TokenNumber {
								tokens[len(tokens)-1].Value += string(char)
							} else {
								tokens = append(tokens, Token{Type: TokenNumber, Value: string(char)})
							}
				}
		}
		return tokens
}

func SolveExpr(s string) int {
		tokens:=tokenize(s);
		precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
		operatorStack := []Token{};
		output := []Token{};
		for _, token := range tokens {
			switch token.Type {
			case TokenNumber:
				output = append(output, token)
			case TokenOperator:
				if len(operatorStack) > 0 && precedence[operatorStack[len(operatorStack)-1].Value] >= precedence[token.Value] {
					output = append(output, operatorStack[len(operatorStack)-1])
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
				operatorStack = append(operatorStack, token)
			case TokenOpenParen:
				operatorStack = append(operatorStack, token)
			case TokenCloseParen:
				for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].Type != TokenOpenParen {
					output = append(output, operatorStack[len(operatorStack)-1])
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
				// Pop the open parenthesis
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
		}

		for len(operatorStack) > 0 {
			output = append(output, operatorStack[len(operatorStack)-1])
			operatorStack = operatorStack[:len(operatorStack)-1]
		}

		stack := []int{};
		for _, token := range output {
			if token.Type == TokenNumber	{
				num,_:= strconv.Atoi(token.Value);
				stack = append(stack,num);
			}else if token.Type == TokenOperator {
				if len(stack) < 2 {
					panic("Invalid expression")
				}
				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				switch token.Value {
				case "+":
					stack = append(stack, a+b)
				case "-":
					stack = append(stack, a-b)
				case "*":
					stack = append(stack, a*b)
				case "/":
					if b == 0 {
						panic("Division by zero")
					}
					stack = append(stack, a/b)
				}
			}
		}
		if len(stack) != 1 {
			panic("Invalid expression")
		}

		return stack[0]

}


type ObjectType int

const (
	Operand ObjectType = iota
	Operator
	LBrac
	RBrac
	None
)

type Object struct {
	Type  ObjectType
	Value string
}

func randomOperator() string {
	operators := []string{"+", "-", "*", "/"}
	return operators[rand.Intn(len(operators)-1)]
}

func randomNumericOperand(maxDigits int) int {
	num := rand.Intn(maxDigits) + 1
	return num
}



func newExprStr(difficulty int, maxDigits int) string {
	if difficulty <= 0 {
		return ""
	}

	i := difficulty % len(Templates())
	objStack := make([]Object, len(Templates()[i]))
	copy(objStack, Templates()[i])

	var numerator, denominator int
	str := ""
	for _, obj := range objStack {
		switch obj.Type {
		case Operator:
			str += randomOperator()
		case Operand:
			if obj.Value == "/" {
				numerator = randomNumericOperand(maxDigits)
				denominator = randomNumericOperand(maxDigits)
				if denominator < numerator {
					denominator = numerator
				}
				str += fmt.Sprintf("%d/%d", numerator, denominator)
			} else {
				str += strconv.Itoa(randomNumericOperand(maxDigits))
			}
		default:
			str += obj.Value
		}
	}

	return str
}


func convertToString(objStack []Object) string {
	var str string
	for _, obj := range objStack {
		str += obj.Value
	}
	return str
}

func NewExpr(difficulty int, maxDigits int) string {
  rand.Seed(time.Now().UnixNano())
	expression := newExprStr(difficulty, maxDigits)
  return expression;
}
