package main

// Multiple packages can be imported together.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadWords(filename string) []string {
	items := []string{}
	//open the text file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return items
	}
	//close the file later
	defer file.Close()

	//read in each file and split on the whitespace
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}
	return items
}

/*driver code*/
func main() {
	// chars := []string{}
	filename := "example.txt"
	//take in a text file with C++ code
	// chars = ReadWords(filename)
	// for i, s := range chars{
	// 	fmt.Println(i, s)
	// }
	//handle the \n to convert to the character for a stack
	//identify the \n
	//push our identifying character on the stack
	//continue the loop

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	tokens, err := Tokenizer(string(data))
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Println(token.ToString())
	}
}

//*/

/**/

// func main() {
// 	TestTokenizerRepl()

// }

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

//*/
