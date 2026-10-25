# Start-Controller

A start-stop controlling mechanism which prevents incorrect usage of mechanisms.

Such constructs are not used in most programs written in `Go` programming 
language, but without such constructs it is impossible to write libraries which 
work safely.

Unfortunately, due to lack of classes, constructors and normal public-private 
visibility, programs written in `Go` language look like anarchistic hell.
