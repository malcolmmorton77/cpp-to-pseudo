# cpp-to-pseudo

## Problem:

> We propose our project convert C++ code into pseudocode. We used a subset of the C++ language such that any valid string in our grammar is valid C++ code, but not vice versa.

## Solution:

> The input is a text file containing a C++ code snippet containing only lexemes that obey the lexical rules. In addition, code blocks must be formatted according to typical C++ rules, and every statement must end with a semicolon, or the whitespace generator may not work correctly.
A deterministic finite automata first analyzes the maximum length of a particular token given a starting position. The tokenizer uses this DFA to, in linear time, tokenize the input into tokens that represent the series of lexemes present in the code sample. Please note that the code is not analyzed for or certified as being semantically or even grammatically correct. This tool assumes the code is correct, and all the lexemes match one of the rules listed below.
Then, the tokens are handed to a psuedocode generator, which converts the tokens into a "prettified" pseudocode representation that is language agnostic with the intention of being easier to read.

## Lexical Rules
### All lexemes must match one of the following comma-separated regular expressions:
See Go's unicode class implementation to see the exact definition of :class:
- Operations
+, -, *, /, %
++, --
<<, >>
<, >, ==, <=, >=
+=, -=, *=, /=, %=, =
- Keywords
auto, char, double, float, int, long
short, string, else, if, while, for
- Identifiers (This regex does not match keywords)
[_$:alpha:][_$:alphanum:]*
- Control Characters
{, }, (, ), ;
- Boolean Literals
true, false
- String Literals
"([^\\]|(\\.))*"
- Number Literals
((.\d+)|(\d+(.\d*)?))
- Whitespace (These tokens are discarded during pseudocode generation)
:space:+

## To run:

- `go run main.go tokenizer.go pseudo.go dfa.go`

## An example piece of code that we will translate:

```cpp
for(int i = 0; i < 10; i++){
  cout << i * i + i*i << endl;
}
int g = 0;
if(g < 4){
  cout << g is less than 4 << endl;
}
```

```pseudo
for int i = 0 ; i < 10 ; i ++ 
	print << i * i + i * i << endline 
int g = 0 
if g < 4 
	print << g is less than 4 << endline 
```

## Explanation of code

### main.go

- main.go is the driver code and reads in the text file full of example `cpp` code
- main.go contains the `ReadWords` function, and a call to Tokenizer

## Git Branch Notes

### Using branches

- Name your branches based on what you worked on and checkout using `git checkout -b <name_of_branch>`
- Once checked out verify what branch you are located on using `git branch`
- Delete a branch using `git branch -d <name_of_branch>`

### How to commit changes

- Switch to a new terminal
- Display your status using `git status`
- Record the current state of the working directory with `git stash`
- Checkout a new branch using `git checkout -b <name_of_branch>`
- Verify you checked out the correct new branch using `git branch`
  - you will see that an \* is next to the current branch you are working on
- Pop your current state of the working directory using `git stash pop`
- Stage the current changes to the working directory using `git add <path_to_directory>`
- Commmit your changes using `git commit`
  - the command will open a text editor and you can add your commit message above the #
- Finally, push your changes using `git push --set-upstream origin <name_of_branch>`
  - Delete the branch once you have merged
