package main

// Multiple packages can be imported together.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadWords(filename string) []byte {
	//open the text file
	items, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return items
}

// driver code
func main() {
	//declare variables
	filename := "example.txt"

	//take in a text file with C++ code
	chars := ReadWords(filename)

	// grab the tokens using tokenizer function
	tokens, err := Tokenizer(string(chars))
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Println(token.ToString())
	}

	// call the converter to pseudocode
	Pseudo(tokens)
}

func TestTokenizerRepl() {
	fmt.Println(Tokenizer("for(){}"))

	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(Tokenizer(input))

	}
}
