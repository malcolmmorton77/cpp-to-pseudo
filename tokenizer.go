package main

import (
	"fmt"
)

const (
	TID = iota
	TKEY
	TCTRL
	TSLIT
	TNLIT
	TBLIT
	TSPACE
)

const debug bool = false

type Token struct {
	Raw      string
	Type     int
	Position int
	Len      int
}

// Tokenizer takes in a raw string input and creates an ordered list of Tokens representing
// the lexical meaning of the input.
func Tokenizer(input string) (Token, error) {

	runes := []rune(input)

	state := Start
	lastState := Start
	inputi := 0
	for state != NONE && inputi < len(runes) {
		lastState = state
		state = Dfa(state, runes[inputi])
		inputi++

		if debug {
			fmt.Printf("lastState %d currState %d\n", lastState, state)
		}
	}

	// If hit end of input need to convert state to lastState
	if !(inputi < len(runes)) && state != NONE {
		lastState = state
		inputi++
	}

	if !IsAccepted(lastState) {
		return Token{}, fmt.Errorf("syntax error at %d", inputi)
	}

	raw := runes[0 : inputi-1]
	pos := 0

	switch lastState {
	case Id, Alpha:
		return Token{
			Raw:      string(raw),
			Type:     TID,
			Position: pos,
			Len:      len(raw),
		}, nil
	case AUTO, CHAR, DOUBLE, FLOAT, INT, LONG,
		SHORT, STRING, ELSE, IF, WHILE, FOR:
		return Token{
			Raw:      string(raw),
			Type:     TKEY,
			Position: pos,
			Len:      len(raw),
		}, nil
	case CurlyClose, CurlyOpen, Open, Close, Stop:
		return Token{
			Raw:      string(raw),
			Type:     TCTRL,
			Position: pos,
			Len:      len(raw),
		}, nil
	case WHITE:
		return Token{
			Raw:      string(raw),
			Type:     TSPACE,
			Position: pos,
			Len:      len(raw),
		}, nil
	default:
		return Token{}, fmt.Errorf("Token %d not recognized", lastState)
	}
}
