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

## Git Branch Notes

### Using branches

- Name your branches based on what you worked on and checkout using git checkout -b <name_of_branch>
- Once checked out verify what branch you are located on using git branch
- Delete a branch using git branch -d <name_of_branch>

### How to commit changes

- Switch to a new terminal
- Display your status using git status
- Record the current state of the working directory with git stash
- Checkout a new branch using git checkout -b <name_of_branch>
- Verify you checked out the correct new branch using git branch
  - you will see that an \* is next to the current branch you are working on
- Pop your current state of the working directory using git stash pop
- Stage the current changes to the working directory using git add <path_to_directory>
- Commmit your changes using git commit
  - the command will open a text editor and you can add your commit message above the #
- Finally, push your changes using git push --set-upstream origin <name_of_branch>
  - Delete the branch once you have merged
