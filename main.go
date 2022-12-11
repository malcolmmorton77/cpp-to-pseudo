package main

// Multiple packages can be imported together.
import (
	"bufio"
	"fmt"
	"os"
)

func ReadWords(filename string) string {
	items := ""
	//open the text file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return items //return an empty string to buffer after printing error
	}

	//close the file at the end of the function
	defer file.Close()

	//read in each file and split on the bytes
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		items += scanner.Text() // append the scanner text in bytes to items string
	}
	return items // return the items string
}

// driver code
func main() {
	//declare variables
	chars := ""
	filename := "example.cpp"

	//take in a text file with C++ code
	chars = ReadWords(filename)

	//testing the printing of the C++ code
	fmt.Println(chars)
	var i = 0
	for i < len(chars) {
		fmt.Println(string(chars[i]))
		i += 1
	}
	//handle the \n to convert to the character for a stack
	//identify the \n
	//push our identifying character on the stack
	//continue the loop
}
