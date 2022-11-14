# cpp-to-pseudo

## Problem:

We propose our project convert C++ code into pseudocode. We used a subset of the C++ language such that any valid string in our grammar is valid C++ code, but not vice versa.

## Solution:

This will be done using a push down automata. This PDA will consume C++ code character by character, left to right, top to bottom, pushing intermediate and final symbols on to the stack, with an accepted string generating a stack of defined, exclusive symbols using the final symbol set. These final symbols will represent the pseudocode of the C++ code, while intermediate symbols will be used by the PDA to help with ordering and creating semantic meaning in the stack. Since the final symbols won't be human readable, these symbols will then be converted to human readable pseudocode.

## An example piece of code that we will translate:

``` 
  for(int i = 0; i < 10; i++){
    cout << "i * i" << i*i << endl;
  }
  int g = 0
  if(g < 4){
    cout << "g is less than 4" << endl;
  }
```
