package main

import (
	"fmt"
)

const (
	TID    = iota //0
	TKEY          //1
	TCTRL         //2
	TSLIT         //3
	TNLIT         //4
	TBLIT         //5
	TSPACE        //6
	TOP
)

const debug bool = false

type Token struct {
	Raw      string
	Type     int
	Position int
	Len      int
}

func (t Token) ToString() string {
	s := "INV_TOK"
	switch t.Type {
	case TID:
		s = "ID"
	case TKEY:
		s = "KEY"
	case TCTRL:
		s = "CTRL"
	case TSLIT:
		s = "STR_LIT"
	case TNLIT:
		s = "NUM_LIT"
	case TBLIT:
		s = "BOOL_LIT"
	case TSPACE:
		s = "SPACE"
	case TOP:
		s = "OP"
	}

	return fmt.Sprintf("%s\t%s\t(%d,%d)", s, t.Raw, t.Position, t.Len)
}

// Tokenizer takes in a raw string input and creates an ordered list of Tokens representing
// the lexical meaning of the input.
func Tokenizer(input string) ([]Token, error) {

	var ret []Token
	lastTokenPos := 0
	for i := 0; i < len(input); i++ {
		fmt.Println(lastTokenPos)
		runes := []rune(input)

		state := START
		lastState := START
		for state != NONE && i < len(input) {
			lastState = state
			state = Dfa(state, runes[i])
			i++

			if debug {
				fmt.Printf("lastState %d currState %d\n", lastState, state)
			}
		}

		// If hit end of input need to convert state to lastState
		if lastState == NONE || lastState == START {
			lastState = state
			i++
		}

		if !IsAccepted(lastState) {
			return nil, fmt.Errorf("syntax error at %d, final state %d curr state %d. State may not be accepted", i, lastState, state)
		}

		pos := lastTokenPos
		raw := runes[lastTokenPos : i-1]
		lastTokenPos += len(raw)

		switch lastState {
		case Stringlit:
			ret = append(ret, Token{
				Raw:      string(raw),
				Type:     TSLIT,
				Position: pos,
				Len:      len(raw),
			})
			i = lastTokenPos - 1
			continue
		case Id, Alpha:
			ret = append(ret, Token{
				Raw:      string(raw),
				Type:     TID,
				Position: pos,
				Len:      len(raw),
			})
			i = lastTokenPos - 1
			continue
		case Auto, Char, Double, Float, Int, Long,
			Short, String, Else, If, While, For:
			ret = append(ret, Token{
				Raw:      string(raw),
				Type:     TKEY,
				Position: pos,
				Len:      len(raw),
			})
			i = lastTokenPos - 1
			continue
		case CompOp, AssignOp, ArithOp, Plus, Minus, Asterisk, Slash,
			Percent, Equal, Less, More, Incr, Decr, Lshift, Rshift:
			ret = append(ret, Token{
				Raw:      string(raw),
				Type:     TOP,
				Position: pos,
				Len:      len(raw),
			})
			i = lastTokenPos - 1
			continue
		case CurlyClose, CurlyOpen, Open, Close, Stop:
			ret = append(ret, Token{
				Raw:      string(raw),
				Type:     TCTRL,
				Position: pos,
				Len:      len(raw),
			})
			i = lastTokenPos - 1
			continue
		case White:
			ret = append(ret, Token{
				Raw:      string(raw),
				Type:     TSPACE,
				Position: pos,
				Len:      len(raw),
			})
			i = lastTokenPos - 1
			continue
		case Num, Digits:
			ret = append(ret, Token{
				Raw:      string(raw),
				Type:     TNLIT,
				Position: pos,
				Len:      len(raw),
			})
			i = lastTokenPos - 1
			continue
		default:
			if START_INTERMED < lastState && lastState < END_INTERMED {
				ret = append(ret, Token{
					Raw:      string(raw),
					Type:     TID,
					Position: pos,
					Len:      len(raw),
				})
				i = lastTokenPos - 1
				continue
			}
			return nil, fmt.Errorf("Token %d not recognized", lastState)
		}
	}

	return ret, nil
}
