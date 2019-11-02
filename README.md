# Little Go Tutorial

[https://deculler.github.io/gotut/](https://deculler.github.io/gotut/)

https://golang.org/doc/


This tutorial is intended to introduce you to Go from the perspects of systems programming in C.

## Getting Started - cmdln1

Our first example deals with command line arguments and illustrates some basic concepts in Go.  
It can be found in [src/cmdln1/](https://github.com/deculler/gotut/tree/master/src/cmdln1).

Although Go is a compiled language, rather than `cc` and `Makefiles`
the process of building Go applications uses a set of file system conventions
and the `go` utility.  To run this example, `cd` to the `cmdln1` directory
and `go run main.go`.  Or you can build the executable with `go build`.
Notice that it is `./cmdln1` in the current directory - taking its name from
the source directory, not the main file.

### Packages

A Go program is made up of *packages*.  They define the namespace of
program modules.  By convention, the package name is the same as the
last element of the import path.  The package `main` is important only
so far as it defines the function that will be called to start the program,
`main.main`.

Every package that is used by the program is explicitly imported
using the `import` statement.  This replaces the `#include`s of your C
program.  There is no separation of `.h` files for signatures and types
from `.c` files for implementation - every package is defined by
the set of files that declare it.  Each source file declares THE
package which contributes to.

In Go, a name is exported from a package if it begins with a capital letter.
Command line argument strings are not defined in the signature of `main`;
instead the equivalent of `argv` is exported by the `os` package.
The analog of `stdio.h` is the `fmt` package, which exports the
`Printf` function used here.






It treats a directory in your file system as *a workspace*.


The environment variable `$GOPATH` is set to yoiur current workspace. It
is assumed to contains subdirectories:

* `$GOPATH/src` contains Go source files.  Conventionally, each application is in
a directory there
* `$GOPATH/bin` contains executables










### Markdown


```markdown
Syntax highlighted code block

# Header 1
## Header 2
### Header 3

- Bulleted
- List

1. Numbered
2. List

**Bold** and _Italic_ and `Code` text
```

