package token

import "text/scanner"

type TokenReceiver struct {
	tokens []Token
}

func Lex(input string) ([]Token, error) {
	t := TokenReceiver{}
	r := NewRuneEmitter(input)
	for next, err := StateBegin(r, &t); next != nil || err != nil; {
		if err != nil {
			return t.tokens, err
		}
		next, err = next(r, &t)
	}
	return t.tokens, nil
}

func (t *TokenReceiver) Send(item Token) {
	t.tokens = append(t.tokens, item)
}

type RuneEmitter struct {
	ch   chan rune
	done chan struct{}
}

func NewRuneEmitter(input string) *RuneEmitter {
	ch := make(chan rune)
	done := make(chan struct{})
	re := &RuneEmitter{ch: ch, done: done}
	go func(re *RuneEmitter) {
		for _, r := range input {
			re.ch <- r
		}
		re.ch <- scanner.EOF
		close(done)
	}(re)
	return re
}

func (r *RuneEmitter) Next() rune {
	select {
	case next := <-r.ch:
		return next
	case <-r.done:
		return scanner.EOF
	}
}
