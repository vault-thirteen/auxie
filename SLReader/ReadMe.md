# S.L.Reader
### Speed-Limited Reader

A reader based on `io.Reader` and using speed limits.  

Main features of this reader are following.  
* Reader uses a common `io.Reader` interface. 
* Changes of speed limits on the fly are supported. 
* Speed limit changes are in their turn limited by the special ratio set in the 
constructor. 
* Zero speed limit is forbidden as it is nonsense.
* Reader implements `io.Closer` interface.
