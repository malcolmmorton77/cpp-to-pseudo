package main

// Multiple packages can be imported together.
import (
	"fmt"
	"os"
	"bufio"
)


func ReadWords(filename string) []string{
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

//driver code
func main() {
		chars := []string{}
		filename := "example.cpp"
		//take in a text file with C++ code
		chars = ReadWords(filename)
		for i, s := range chars{
			fmt.Println(i, s)
		}
		//handle the \n to convert to the character for a stack
			//identify the \n
			//push our identifying character on the stack
			//continue the loop
}