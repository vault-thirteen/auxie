# CSV

A writer for _CSV_ format.

As opposed to the poor _CSV_ writer of the standard library, this writer is 
using double quotes for strings and is able to distinguish different types of 
input variables. To do so, it uses the `interface` (`any`) type variables to 
emulate dynamic types.
