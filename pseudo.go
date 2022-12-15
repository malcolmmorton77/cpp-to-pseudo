package main

import (
	"errors"
	"fmt"
	"os"
)

// takes in the list of tokens, parses them
// returns nothing
// outputs a file full of the pseudocode
// taking out control blocks ({}) and semicolons, replaces cout with print and endl with endline
func Pseudo(tokens []Token) error {
	flag := 0
	var tab []rune
	block := -2
	code := ""
	// read from tokens and output the translated piece of code
	for _, token := range tokens {
		switch token.Type {
		case TID: // IDs
			if token.Raw == "cout" {
				code += "print"
				code += " "
				continue
			}
			if token.Raw == "endl" {
				code += "endline"
				code += " "
				continue
			}
			code += token.Raw
			code += " "
		case TKEY: // Keys
			if token.Raw == "for" {
				code += token.Raw
				flag = 1
				block = -1
				continue
			}
			if token.Raw == "if" {
				code += token.Raw
				block = -1
				continue
			}
			if token.Raw == "while" {
				code += token.Raw
				block = -1
				continue
			}
			if token.Raw == "int" || token.Raw == "auto" || token.Raw == "string" || token.Raw == "char" || token.Raw == "double" ||
				token.Raw == "float" || token.Raw == "long" || token.Raw == "short" {
				code += "let"
				code += " "
				continue
			}
			code += token.Raw
			code += " "
		case TCTRL: //Ctrl structures
			if token.Raw == "{" {
				tab = append(tab, '\t')
				code += "\n"
				code += string(tab)
				continue
			}
			if token.Raw == "}" { // end of control block
				if len(tab) > 0 {
					tab = tab[:len(tab)-1]
				}
				continue
			}
			if token.Raw == ")" {
				if flag == 1 {
					flag = 0
				}

				//code += fmt.Sprintf("~%d", block)

				switch block {
				case -2:
					code += ")"
				case -1:
					fmt.Println("Block paren err")
				case 0:
					block = -2
				default:
					block--
					code += ")"
					if block == 0 {
						block = -2
					}
				}
				continue
			}
			if token.Raw == "(" {
				// code += fmt.Sprintf("~%d", block)
				switch block {
				case -2:
					code += "("
				case -1:
					code += " "
					block = 0
				default:
					block++
					code += "("
				}

				continue
			}
			if token.Raw == ";" && flag == 1 {
				code += token.Raw
				code += " "
				continue
			}
			if token.Raw == ";" {
				code += "\n"
				code += string(tab)
			}
		case TSLIT, TNLIT, TBLIT, TOP: // print literals
			code += token.Raw
			code += " "
		case TSPACE: // ignore space
			continue
		default:
			return errors.New("Token is not recognized")
		}
	}

	err := os.WriteFile("pseudocode.txt", []byte(code), 0644)
	return err
}
