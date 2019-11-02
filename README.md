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

Every package that is used by the program is explicitly *imported*
using the `import` statement.  This replaces the `#include`s of your C
program.  There is no separation of `.h` files for signatures and types
from `.c` files for implementation - every package is defined by
the set of files that declare it.  Each source file declares THE
package which contributes to.

In Go, a name is exported from a package if it begins with a capital letter.

Command line argument strings are not defined in the signature of `main`;
instead the equivalent of `argv` is exported by the `os` package.
The analog of `stdio.h` is the `fmt` package, which exports the
`Printf` function used here.  Note capitalization in both, but not in the
package name.

### Comments

Like C, Go allows multiline comments bracketed by `/* */` and in-line comments with `//`.

### Assignment and Types

Assignment is denoted `:=` as in Python, and tuple assignment is permitted, e.g.,
`x, y := 2, 4`.

Assignment implicitly declares the LHS variables, if they are not already
declared.  The general form of a declaration optionally declares the type
and optionally the initial value.
```
var <name> [<type>] [= <initial value>]
```
Note here we are declaring its initial value with `=`, not performing
an assignment.  Thus, we could have been explicit as follows, or left
off the `int` type.
```
	var argc int = len(argv)
```
The `var` statement declares a list of variables.  

Where a variable is declared with an initial value, its type is
inferred from that of the initializer.  Go utilizes type inference
extensively.  It also allows types to be inspected.  In our example,
we print the type of certain values with the `%T` format directives.

More generally, assignment can include a [*type assertion*]
(https://tour.golang.org/methods/15):
```
x := v.(T)
```
Asserts that the underlying value of `v` is of type `T` and assigns that value
to `x`.  If not, it triggers a panic.  Or for programmatic control
```
x, ok := v.(T)
```
assigns a boolean to the second LHS of the type match.  The
[Basic Types](https://tour.golang.org/basics/11) in Go are similar to
those you would find in `stdtypes.h`.  But there is the ambiguity of `char`
is eliminated.  `byte` is an alias for `int8` whereas `rune` represents
a Unicode code point and is an alias for `int32`.


In Go, the type of a variable follows the variable name and is read left to right,
the reverse of C. We see this with the `var` statement above.  But notice that
`argv` is of type `[]string` - which could be read "array of strings".  Recall
that in C we would declare
```
char **argv[];
or
char ***argv;
```
read "argv is an array of pointers to strings" - the type is right to left.

### Arrays and Slices

Go takes arrays seriously and resolves many of the issues that have plagued
strongly typed languages of the past, as well as the need for explicit
memory management in C when attempting to extend an array.













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

