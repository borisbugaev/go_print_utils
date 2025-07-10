very simple wrappers for console output operations

functions:

**Print_Lines()** *-- passed string, returns int*

- prints string to console
- returns the number of lines in said string

**Clear_Lines()** *-- passed int, void*

- clears *int* lines from console

**Line_Select_MC()** *-- passed []string, returns string*

- assigns letter to each string in slice and prints all strings to console

- prompts user to select a letter

- returns string which corresponds with response letter

- multiple letters can be given as response, separated with spaces

- in this case, function returns all selected strings as a single string separated by commas

- if an invalid option is selected, returns "\\a" (bell)

these exist to simplify interactions between a user and command-line application

not all command line applications are complex enough to necessitate a full terminal UI implementation

these should bridge the gap
