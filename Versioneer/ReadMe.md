# Versioneer

Product and dependency version helper.

### Description

This package does several useful things.

1. **It prints some nice looking text about the following things:**
* Name of the program
* Version of the program
* List of dependencies containing:
  * Dependency names
  * Dependency versions

2. **It checks for updates on GitHub:**
* If an updated version of the program is available, this library tell about it.

### Attention
Due to an old bug **#29228** issued on 13.12.2018 
( https://github.com/golang/go/issues/29228 ) which resulted in a chain of 
other opened issues, such as **#50603** 
( https://github.com/golang/go/issues/50603 ), we can see that current version 
of Go language incorrectly shows version text of the `main` package manually 
built from the repository's source code. It means that versioning mechanism in 
Go is still not working as it should `:/`

![Mister Bean](img/Mister_Been_640x360.png)

Well ... this is Go, my friends. This is not the C#.

The good news is that one of the comments in the issue **#50603** states that:
> The version is recorded as expected when doing **go install**

This means that if you see a `(devel)` version instead of a normal version
taken from the Git tag, you should use `go install` instead of `go build` to
build your projects.
