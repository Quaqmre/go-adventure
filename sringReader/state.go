package token

import (
	"fmt"
	"text/scanner"
)

type runeEmitter interface {
	Next() rune
}
type tokenReceiver interface {
	Send(Token)
}

type StateFunc func(runeEmitter, tokenReceiver) (StateFunc, error)

func StateBegin(r runeEmitter, t tokenReceiver) (StateFunc, error) {
	next := r.Next()

	switch next {
	case ' ':
		return StateBegin, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return StateNumber([]rune{next}), nil
	default:
		return nil, fmt.Errorf("begin error")
	}
}
func StateNumber(read []rune) StateFunc {
	return func(r runeEmitter, t tokenReceiver) (StateFunc, error) {
		next := r.Next()
		switch next {
		case ' ':
			return StateNumber(read), nil
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return StateNumber(append(read, next)), nil
		case '+', '-', '*', '/':
			t.Send(Number{Value: string(read)})
			return StateOperator, nil
		case scanner.EOF:
			return nil, nil
		default:
			return nil, fmt.Errorf("thre is some error")
		}
	}
}

func StateOperator(r runeEmitter, t tokenReceiver) (StateFunc, error) {
	next := r.Next()
	switch next {
	case ' ':
		return StateOperator, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		t.Send(Operator{Value: string(next)})
		return StateNumber([]rune{next}), nil
	case scanner.EOF:
		return nil, nil
	default:
		return nil, fmt.Errorf("operator error")
	}
}
