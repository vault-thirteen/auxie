# Tester

This package provides various useful functions to perform tests.  

This package has methods for:
* Assertions:
  * Equality – `MustBeEqual`
  * Difference – `MustBeDifferent`


* Checking errors:
  * Error is set – `MustBeAnError`
  * No error is set – `MustBeNoError`

    
These methods can be used to compare values, to check errors and even to make 
more complex comparisons from simple ones.

When a comparison fails, it not only shows the values which failed, but also 
prints a 'Diff', a difference between these values. As opposed to many existing 
_Golang_ libraries in the Internet, this package uses the object-oriented 
programming style (O.O.P.) as in such mature programming languages as _C#_ 
and _Java_.
