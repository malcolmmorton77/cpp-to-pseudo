package main

// Multiple packages can be imported together.
import (
	"fmt"
	"os"
	"bufio"
)


func ReadWords(filename string){
		//open the text file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
		//close the file later
	defer file.Close()

		//read in each file and split on the whitespace
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//driver code
func main() {
		filename := "example.txt"
		//take in a text file with C++ code
		ReadWords(filename)
		//handle the \n to convert to the character for a stack
			//identify the \n
			//push our identifying character on the stack
			//continue the loop
}