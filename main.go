package main

// Multiple packages can be imported together.
import (
	"fmt"
	"os"
)

func ReadWords(filename string) []byte {
	//open the text file
	items, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	return items
}

// driver code
func main() {
	//declare variables
	filename := "text.txt"

	//take in a text file with C++ code
	chars := ReadWords(filename)

	// grab the tokens using tokenizer function
	tokens, err := Tokenizer(string(chars))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	// call the converter to pseudocode
	err = Pseudo(tokens)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
}
