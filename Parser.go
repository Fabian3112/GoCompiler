package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseProgramm(programm string) Stmt {
	tokens := tokenize(programm)
	rpn := convertToReversePolishNotation(tokens)
	ast := rpnToAst(rpn)
	return ast
}

func tokenize(programm string) []Token {

	programm = strings.ReplaceAll(programm, "\n", " ")
	programm = strings.ReplaceAll(programm, "\t", " ")

	allRecognizbaleTokens := []tokenType{MultToken, PlusToken, AndToken, OrToken, NotToken, EqualToken, LesserToken, GroupOpenToken, GroupCloseToken,
		BlockOpenToken, BlockCloseToken, PrintToken, WhileToken, IfToken, ElseToken, DeclarationToken, SemikononToken}

	for _, token := range allRecognizbaleTokens {
		programm = strings.ReplaceAll(programm, string(token), " "+string(token)+" ")
	}

	programm = strings.ReplaceAll(programm, "true", " true ")
	programm = strings.ReplaceAll(programm, "false", " false ")

	//escape == and := before replacing =
	programm = strings.ReplaceAll(programm, "==", "##")
	programm = strings.ReplaceAll(programm, ":=", ":#")

	programm = strings.ReplaceAll(programm, "=", " = ")
	//undo escaping
	programm = strings.ReplaceAll(programm, "##", "==")
	programm = strings.ReplaceAll(programm, ":#", ":=")

	//remove duplicate spaces
	programm = strings.ReplaceAll(programm, "  ", " ")

	tokensAsString := strings.Split(programm, " ")
	var tokens []Token

	for _, tokenAsString := range tokensAsString {
		if tokenAsString != "" {
			var token Token
			switch tokenAsString {
			case "*":
				token = Token{tokenType: MultToken}
			case "+":
				token = Token{tokenType: PlusToken}
			case "&&":
				token = Token{tokenType: AndToken}
			case "||":
				token = Token{tokenType: OrToken}
			case "!":
				token = Token{tokenType: NotToken}
			case "==":
				token = Token{tokenType: EqualToken}
			case "<":
				token = Token{tokenType: LesserToken}
			case "true":
				token = Token{tokenType: BooleanToken, value: "true"}
			case "false":
				token = Token{tokenType: BooleanToken, value: "false"}

			case "(":
				token = Token{tokenType: GroupOpenToken}
			case ")":
				token = Token{tokenType: GroupCloseToken}
			case "{":
				token = Token{tokenType: BlockOpenToken}
			case "}":
				token = Token{tokenType: BlockCloseToken}

			case "print":
				token = Token{tokenType: PrintToken}
			case "while":
				token = Token{tokenType: WhileToken}
			case "if":
				token = Token{tokenType: IfToken}
			case "else":
				token = Token{tokenType: ElseToken}
			case ":=":
				token = Token{tokenType: DeclarationToken}
			case "=":
				token = Token{tokenType: AssignToken}
			case ";":
				token = Token{tokenType: SemikononToken}

			default:
				_, err := strconv.ParseInt(tokenAsString, 10, 64)
				if err == nil {
					token = Token{tokenType: NumberToken, value: tokenAsString}
				} else {
					token = Token{tokenType: VarToken, value: tokenAsString}
				}
			}

			tokens = append(tokens, token)
		}
	}

	return tokens
}

func convertToReversePolishNotation(tokens []Token) []Token {
	stack := StackToken{}
	var result []Token

	for _, token := range tokens {
		switch token.tokenType {
		case VarToken, NumberToken, BooleanToken:
			result = append(result, token)
		case MultToken, PlusToken, AndToken, OrToken, EqualToken, LesserToken:
			evaluateOperator(token, &stack, &result)
		case NotToken:
			evaluateOperator(token, &stack, &result)
		case GroupOpenToken, BlockOpenToken:
			stack.push(token)
		case GroupCloseToken, BlockCloseToken:
			evaluateGroupClose(token, &stack, &result)
		case PrintToken, WhileToken, IfToken, ElseToken, DeclarationToken, SemikononToken, AssignToken:
			evaluateOperator(token, &stack, &result)
		}
	}
	sortToRPN(&stack, &result)

	return result
}

func evaluateOperator(token Token, stack *StackToken, result *[]Token) {
	for {
		var tokenFromStack Token
		tokenFromStack, err := stack.peak()
		if err != nil {
			break
		}
		if tokenFromStack.tokenType == GroupOpenToken || tokenFromStack.tokenType == BlockOpenToken {
			break
		} else if tokenFromStack.ranking() >= token.ranking() {
			stack.pop()
			*result = append(*result, tokenFromStack)
		} else {
			break
		}
	}
	stack.push(token)
}

func evaluateGroupClose(token Token, stack *StackToken, result *[]Token) error {
	for {
		tokenFromStack, err := stack.pop()
		if err != nil {
			return errors.New("expected (")
		}
		if (tokenFromStack.tokenType == GroupOpenToken && token.tokenType == GroupCloseToken) ||
			(tokenFromStack.tokenType == BlockOpenToken && token.tokenType == BlockCloseToken) {
			break
		} else {
			*result = append(*result, tokenFromStack)
		}

	}
	return nil
}

/*
func evaluateUnaryArgumentOperator(token Token, stack *Stack, result *[]Token){

}


func evaluateInstruction(token Token, stack *Stack, result *[]Token){

}
*/

func sortToRPN(stack *StackToken, result *[]Token) error {
	for {
		tokenFromStack, err := stack.pop()
		if err != nil {
			break
		}
		if tokenFromStack.tokenType == GroupOpenToken {
			return errors.New("expected )")
		}
		*result = append(*result, tokenFromStack)
	}
	return nil
}

type StackToken struct {
	tokens []Token
}

func (stack *StackToken) push(token Token) {
	stack.tokens = append(stack.tokens, token)
}
func (stack *StackToken) pop() (Token, error) {
	if len(stack.tokens) == 0 {
		return Token{}, errors.New("Stack Empty")
	}

	tokens := stack.tokens
	var token Token
	token, stack.tokens = tokens[len(tokens)-1], tokens[:len(tokens)-1]
	return token, nil
}

func (stack *StackToken) peak() (Token, error) {
	if len(stack.tokens) == 0 {
		return Token{}, errors.New("Stack Empty")
	}

	tokens := stack.tokens
	token := tokens[len(tokens)-1]
	return token, nil
}

func myPrint(tokens []Token) {

	for _, token := range tokens {
		if token.tokenType == NumberToken || token.tokenType == VarToken {
			fmt.Printf("%s ", token.value)
		} else {
			fmt.Printf("%s ", token.tokenType)
		}
	}
}

func rpnToAst(tokens []Token) Stmt {
	var expStack []Exp
	var stmtStack []Stmt

	for _, token := range tokens {
		switch token.tokenType {
		case NumberToken:
			num, _ := strconv.ParseInt(token.value, 10, 64)
			push[Exp](&expStack, number(int(num)))

		case BooleanToken:
			if token.value == "true" {
				push[Exp](&expStack, boolean(true))
			} else {
				push[Exp](&expStack, boolean(false))
			}

		case VarToken:
			push[Exp](&expStack, variable(token.value))

		case MultToken, PlusToken, AndToken, OrToken, EqualToken, LesserToken:
			right, _ := pop(&expStack)
			left, _ := pop(&expStack)
			push[Exp](&expStack, makeOperatorExp(token, *left, *right))

		case PrintToken:
			exp, _ := pop(&expStack)
			push[Stmt](&stmtStack, printStmt(*exp))

		case WhileToken:
			exp, _ := pop(&expStack)
			stmt, _ := pop(&stmtStack)
			push[Stmt](&stmtStack, whileStmt(*exp, block(*stmt)))

		case IfToken:

		case ElseToken:
			exp, _ := pop(&expStack)
			elsestmt, _ := pop(&stmtStack)
			ifstmt, _ := pop(&stmtStack)
			push[Stmt](&stmtStack, ifThenElse(*exp, block(*ifstmt), block(*elsestmt)))

		case DeclarationToken:
			right, _ := pop(&expStack)
			variable, _ := pop(&expStack)
			push[Stmt](&stmtStack, decl((*variable).pretty(), *right))

		case AssignToken:
			right, _ := pop(&expStack)
			variable, _ := pop(&expStack)
			push[Stmt](&stmtStack, assign((*variable).pretty(), *right))

		case SemikononToken:
			secondStmt, _ := pop(&stmtStack)
			firstStmt, _ := pop(&stmtStack)
			push[Stmt](&stmtStack, seq(*firstStmt, *secondStmt))

		case NotToken:
			exp, _ := pop(&expStack)
			push[Exp](&expStack, not(*exp))

		default:
			fmt.Printf("unexpected Token %s", token.tokenType)
		}

	}

	return stmtStack[0]
}

func makeOperatorExp(token Token, left Exp, right Exp) Exp {
	switch token.tokenType {
	case MultToken:
		return mult(left, right)
	case PlusToken:
		return plus(left, right)
	case AndToken:
		return and(left, right)
	case OrToken:
		return or(left, right)
	case EqualToken:
		return equal(left, right)
	case LesserToken:
		return lesser(left, right)
	default:
		return nil
	}
}

func pop[T any](stack *[]T) (*T, error) {
	if len(*stack) == 0 {
		return nil, errors.New("Stack Empty")
	}

	var t T
	t, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
	return &t, nil
}

func push[T any](stack *[]T, t T) {
	*stack = append(*stack, t)
}
