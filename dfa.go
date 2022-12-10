package main

import "unicode"

const (
	NONE = iota
	Start
	Alpha
	Id
	Num
	CurlyOpen
	CurlyClose
	Open
	Close
	Stop
	WHITE
	WHILE
	AUTO
	CHAR
	DOUBLE
	FLOAT
	LONG
	SHORT
	STRING
	ELSE
	FOR
	IF
	INT

	START_INTERMED
	A
	AU
	AUT
	C
	CH
	CHA
	D
	DO
	DOU
	DOUB
	DOUBL
	F
	FL
	FLO
	FLOA
	I
	IN
	L
	LO
	LON
	S
	SH
	SHO
	SHOR
	ST
	STR
	STRI
	STRIN
	E
	EL
	ELS
	FO
	W
	WH
	WHI
	WHIL
	END_INTERMED
)

func IsAccepted(s int) bool {
	return s == AUTO ||
		s == CHAR ||
		s == DOUBLE ||
		s == FLOAT ||
		s == INT ||
		s == LONG ||
		s == SHORT ||
		s == STRING ||
		s == ELSE ||
		s == FOR ||
		s == IF ||
		s == WHILE ||
		s == CurlyClose ||
		s == CurlyOpen ||
		s == Open ||
		s == Close ||
		s == Stop ||
		s == Alpha ||
		s == Id ||
		s == Num ||
		s == WHITE ||
		(START_INTERMED < s && s < END_INTERMED)
}

func Dfa(state int, input rune) int {
	switch state {
	case NONE:
		return NONE
	case Start:
		if input == 'a' {
			return A
		} else if input == 'c' {
			return C
		} else if input == 'd' {
			return D
		} else if input == 'e' {
			return E
		} else if input == 'f' {
			return F
		} else if input == 'i' {
			return I
		} else if input == 'l' {
			return L
		} else if input == 's' {
			return S
		} else if input == 'w' {
			return W
		} else if input == '{' {
			return CurlyOpen
		} else if input == '}' {
			return CurlyClose
		} else if input == '(' {
			return Open
		} else if input == ')' {
			return Close
		} else if input == ';' {
			return Stop
		} else if unicode.IsDigit(input) {
			return Num
		} else if unicode.IsLetter(input) || input == '_' || input == '$' {
			return Alpha
		} else if unicode.IsSpace(input) {
			return WHITE
		} else {
			return NONE
		}
	case Alpha:
		return kwIdClosure(input)
	case Id:
		if unicode.IsLetter(input) || unicode.IsDigit(input) ||
			input == '_' || input == '$' {
			return Id
		} else {
			return NONE
		}
	case Num:

	case CurlyOpen:
		return NONE
	case CurlyClose:
		return NONE
	case Open:
		return NONE
	case Close:
		return NONE
	case Stop:
		return NONE
	case WHITE:
		if unicode.IsSpace(input) {
			return WHITE
		} else {
			return NONE
		}

	case AUTO:
		return kwIdClosure(input)
	case CHAR:
		return kwIdClosure(input)
	case DOUBLE: //
		return kwIdClosure(input)
	case FLOAT:
		return kwIdClosure(input)
	case INT:
		return kwIdClosure(input)
	case LONG:
		return kwIdClosure(input)
	case SHORT:
		return kwIdClosure(input)
	case STRING:
		return kwIdClosure(input)
	case ELSE:
		return kwIdClosure(input)
	case IF:
		return kwIdClosure(input)
	case WHILE:
		return kwIdClosure(input)
	case FOR:
		return kwIdClosure(input)

	case A:
		if input == 'u' {
			return AU
		} else {
			return kwIdClosure(input)
		}
	case AU:
		if input == 't' {
			return AUT
		} else {
			return kwIdClosure(input)
		}
	case AUT:
		if input == 'o' {
			return AUTO
		} else {
			return kwIdClosure(input)
		}
	case C:
		if input == 'h' {
			return CH
		} else {
			return kwIdClosure(input)
		}
	case CH:
		if input == 'a' {
			return CHA
		} else {
			return kwIdClosure(input)
		}
	case CHA:
		if input == 'r' {
			return CHAR
		} else {
			return kwIdClosure(input)
		}
	case D:
		if input == 'o' {
			return DO
		} else {
			return kwIdClosure(input)
		}
	case DO:
		if input == 'u' {
			return DOU
		} else {
			return kwIdClosure(input)
		}
	case DOU:
		if input == 'b' {
			return DOUB
		} else {
			return kwIdClosure(input)
		}
	case DOUB:
		if input == 'l' {
			return DOUBL
		} else {
			return kwIdClosure(input)
		}
	case DOUBL:
		if input == 'e' {
			return DOUBLE
		} else {
			return kwIdClosure(input)
		}
	case F:
		if input == 'o' {
			return FO
		} else if input == 'l' {
			return FL
		} else {
			return kwIdClosure(input)
		}
	case FL:
		if input == 'o' {
			return FLO
		} else {
			return kwIdClosure(input)
		}
	case FLO:
		if input == 'a' {
			return FLOA
		} else {
			return kwIdClosure(input)
		}
	case FLOA:
		if input == 't' {
			return FLOAT
		} else {
			return kwIdClosure(input)
		}
	case I:
		if input == 'n' {
			return IN
		} else if input == 'f' {
			return IF
		} else {
			return kwIdClosure(input)
		}
	case IN:
		if input == 't' {
			return INT
		} else {
			return kwIdClosure(input)
		}
	case L:
		if input == 'o' {
			return LO
		} else {
			return kwIdClosure(input)
		}
	case LO:
		if input == 'n' {
			return LON
		} else {
			return kwIdClosure(input)
		}
	case LON:
		if input == 'g' {
			return LONG
		} else {
			return kwIdClosure(input)
		}
	case S:
		if input == 'h' {
			return SH
		} else if input == 't' {
			return ST
		} else {
			return kwIdClosure(input)
		}
	case SH:
		if input == 'o' {
			return SHO
		} else {
			return kwIdClosure(input)
		}
	case SHO:
		if input == 'r' {
			return SHOR
		} else {
			return kwIdClosure(input)
		}
	case SHOR:
		if input == 't' {
			return SHORT
		} else {
			return kwIdClosure(input)
		}
	case ST:
		if input == 'r' {
			return STR
		} else {
			return kwIdClosure(input)
		}
	case STR:
		if input == 'i' {
			return STRI
		} else {
			return kwIdClosure(input)
		}
	case STRI:
		if input == 'n' {
			return STRIN
		} else {
			return kwIdClosure(input)
		}
	case STRIN:
		if input == 'g' {
			return STRING
		} else {
			return kwIdClosure(input)
		}
	case E:
		if input == 'l' {
			return EL
		} else {
			return kwIdClosure(input)
		}
	case EL:
		if input == 's' {
			return ELS
		} else {
			return kwIdClosure(input)
		}
	case ELS:
		if input == 'e' {
			return ELSE
		} else {
			return kwIdClosure(input)
		}
	case FO:
		if input == 'r' {
			return FOR
		} else {
			return kwIdClosure(input)
		}
	case W:
		if input == 'h' {
			return WH
		} else {
			return kwIdClosure(input)
		}
	case WH:
		if input == 'i' {
			return WHI
		} else {
			return kwIdClosure(input)
		}
	case WHI:
		if input == 'l' {
			return WHIL
		} else {
			return kwIdClosure(input)
		}
	case WHIL:
		if input == 'e' {
			return WHILE
		} else {
			return kwIdClosure(input)
		}
	}

	return NONE
}

func kwIdClosure(input rune) int {
	if unicode.IsLetter(input) {
		return Alpha
	} else if input == '_' || input == '$' {
		return Alpha
	} else if unicode.IsDigit(input) {
		return Id
	} else {
		return NONE
	}
}
