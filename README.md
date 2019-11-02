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

The [Basic Types](https://tour.golang.org/basics/11) in Go are similar to
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
memory management in C when attempting to extend an array.  Like
Python, and unlike C, an array carries its length, which can be
obtained with the `len` builtin function.  Can the language will ensure
that you do not index outside the bounds of the array.  In C we often allocate
an array like thing (e.g., with `malloc`) and play games with pointers and
`realloc` to get sections of the array or extensions of it.  Go tackles
these issues directly with *slices*.

Arrays are declared with an explicit type and length - and their length is part of their
type (like in Pascal of old...).  The following declares `a` to be an array of 10 integers (note
the order and the spacing).
```
var a [10]int
```
Generally, at the point of declaration the length of the array is given by a integer
expression and it cannot be resized.  The more flexible thing that we often use
pointers in C to express is a slice - which is a dynamically sized view into some array.
We could have declared
```
var argv []string
```
which you'll notice is its inferred type.  The size of the slice depends on how many
command line arguments were passed in.  That's what `argc` was for.  We need to be able
to write a general purpose `main` function that works for any number of args.  Thus,
something outside of it needs to worry about the array that provided the storage
allocation.  We are provided a slice.

As in Python, a slice is formed by specifying two indices, a low and high bound,
separated by a `:`, omitted bounds are treated as the origin and last.

In addition to length, a slice has a *capacity*, obtained by `cap(s)`, that is
the number of elements in the underlying arrauy starting from the first element of the
slice.  A slice can be truncarted or  extended by *reslicing*, for example `s = s[:5]`, up to
its capacity.

### Iteration

The first loop in our example shows the Go idiomatic form of iteration,
binding the index and value to each of the elements of a sequence.  For just
the values this would be `for _, v := range argv { ... }`

The 3-expression form common in C is retained, but note that there is no
enclosing parentheses - the braces enclosing the body of the loop are always required.
This should be interpreted as
```
for <init stmt>; <condition stmt>; <post stmt> { <body statements> }
```
and any of them are option.  The init statements run before the condition; the
comndition is run before each iteration of the body to determine of the loop has
terminated; the post statements are run after the body and before the next
condition is tested.  With this intepretation it is common to declare variables
in the init statement.  We have used the short form.  It could have been
```
for var i = 0; i < argc; i++ { ... }
```
While loops have the condition but no init or post.  This can be expressed as
```
for ; <cond> ; { ... }
```
or more idiomatically as the concise
```
for <cond>  { ... }
```

Conditionals and switch statements follow a similar reductionist approach as iteration.

### Go, make, and GOPATH

Whereas in typical C development we have a Makefile to define how we build
and deploy and executible, Go employs a set of conventions in the use of the
file system.

Go treats a directory in your file system as *a workspace*.
The environment variable `$GOPATH` is set to yoiur current workspace. It
is assumed to contains subdirectories:

* `$GOPATH/src` contains Go source files.  Conventionally, each application is in
a directory there
* `$GOPATH/bin` contains executables

`go install` will build the current command package and install its executable in `bin`.  You
can add that to your `PATH` if you want to run what you build. `go clean` does what you'd
expect.






## more stuff

### Type assertions

Need interfaces before this

An assignment can include a [*type assertion*]
(https://tour.golang.org/methods/15):
```
x := v.(T)
```
Asserts that the underlying value of `v` is of type `T` and assigns that value
to `x`.  If not, it triggers a panic.  Or for programmatic control
```
x, ok := v.(T)
```
assigns a boolean to the second LHS of the type match. 


















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

