package main

import "unicode"

const (
	NONE = iota
	START
	Alpha
	Id
	Num
	White

	CurlyOpen
	CurlyClose
	Open
	Close
	Stop
	CompOp
	AssignOp
	ArithOp
	Incr
	Decr

	Plus
	Minus
	Asterisk
	Slash
	Percent
	Equal
	Less
	More

	While
	Auto
	Char
	Double
	Float
	Long
	Short
	String
	Else
	For
	If
	Int

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
	return s == Auto ||
		s == Char ||
		s == Double ||
		s == Float ||
		s == Int ||
		s == Long ||
		s == Short ||
		s == String ||
		s == Else ||
		s == For ||
		s == If ||
		s == While ||

		s == CurlyClose ||
		s == CurlyOpen ||
		s == Open ||
		s == Close ||
		s == Stop ||
		s == CompOp ||
		s == AssignOp ||
		s == ArithOp ||
		s == Plus ||
		s == Minus ||
		s == Asterisk ||
		s == Slash ||
		s == Percent ||
		s == Equal ||
		s == Less ||
		s == More ||
		s == Incr ||
		s == Decr ||

		s == Alpha ||
		s == Id ||
		s == Num ||
		s == White ||
		(START_INTERMED < s && s < END_INTERMED)
}

func Dfa(state int, input rune) int {
	switch state {
	case NONE:
		return NONE
	case START:
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
		} else if input == '<' {
			return Less
		} else if input == '>' {
			return More
		} else if input == '+' {
			return Plus
		} else if input == '-' {
			return Minus
		} else if input == '*' {
			return Asterisk
		} else if input == '/' {
			return Slash
		} else if input == '%' {
			return Percent
		} else if input == '=' {
			return Equal
		} else if input == ';' {
			return Stop

		} else if unicode.IsDigit(input) {
			return Num
		} else if unicode.IsLetter(input) || input == '_' || input == '$' {
			return Alpha
		} else if unicode.IsSpace(input) {
			return White
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
	case CompOp:
		return NONE
	case Plus:
		if input == '=' {
			return AssignOp
		} else if input == '+' {
			return Incr
		} else {
			return NONE
		}
	case Minus:
		if input == '=' {
			return AssignOp
		} else if input == '-' {
			return Decr
		} else {
			return NONE
		}
	case Asterisk:
		if input == '=' {
			return AssignOp
		} else {
			return NONE
		}
	case Slash:
		if input == '=' {
			return AssignOp
		} else {
			return NONE
		}
	case Percent:
		if input == '=' {
			return AssignOp
		} else {
			return NONE
		}
	case Equal:
		if input == '=' {
			return CompOp
		} else {
			return NONE
		}
	case Less:
		if input == '=' {
			return CompOp
		} else {
			return NONE
		}
	case More:
		if input == '=' {
			return CompOp
		} else {
			return NONE
		}
	case Incr, Decr:
		return NONE
	case White:
		if unicode.IsSpace(input) {
			return White
		} else {
			return NONE
		}

	case Auto:
		return kwIdClosure(input)
	case Char:
		return kwIdClosure(input)
	case Double: //
		return kwIdClosure(input)
	case Float:
		return kwIdClosure(input)
	case Int:
		return kwIdClosure(input)
	case Long:
		return kwIdClosure(input)
	case Short:
		return kwIdClosure(input)
	case String:
		return kwIdClosure(input)
	case Else:
		return kwIdClosure(input)
	case If:
		return kwIdClosure(input)
	case While:
		return kwIdClosure(input)
	case For:
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
			return Auto
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
			return Char
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
			return Double
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
			return Float
		} else {
			return kwIdClosure(input)
		}
	case I:
		if input == 'n' {
			return IN
		} else if input == 'f' {
			return If
		} else {
			return kwIdClosure(input)
		}
	case IN:
		if input == 't' {
			return Int
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
			return Long
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
			return Short
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
			return String
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
			return Else
		} else {
			return kwIdClosure(input)
		}
	case FO:
		if input == 'r' {
			return For
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
			return While
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
