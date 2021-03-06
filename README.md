# Introduction to Go - from a C in systems perspective

Web page version of this tutorial:
[https://deculler.github.io/gotut/](https://deculler.github.io/gotut/)

Reference documention for Go: [https://golang.org/doc/](https://golang.org/doc/)

This tutorial is intended to introduce you to Go from the perspective of
systems programming in C.  Our goal is not to transliterate C into Go,
but to utilize your understanding of C, its flexibility and its
shortcomings, to help understand and appreciate how to approach Go.
To this end, our progressive example os the "word counting" exercise
[used in CS162](https://cs162.eecs.berkeley.edu/static/hw/hw1.pdf).  We build
variants of this up step by step to illustrate essential comncepts in Go
while putting them into action.

## Getting Started - `cmdln/main.go`

Our first example deals with basic command line arguments to illustrate
getting started in Go.  
It can be found in [src/cmdln/](https://github.com/deculler/gotut/tree/master/src/cmdln).

Although Go is a compiled language, rather than `cc` and `Makefiles`
the process of building Go applications uses a set of file system conventions
and the `go` utility.  To run this example, `cd` to the `cmdln` directory
and `go run main.go`.  Or you can build the executable with `go build`.
Notice that it builds `./cmdln` in the current directory - taking its name from
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

Being effective in a language is about the libraries as much as the language concepts,
and these are two important ones.
*[fmt package](https://golang.org/pkg/fmt/) provides formated printing and scanning.
*[os package](https://golang.org/pkg/os/) provides platform independent operating system
functionality, like files, directories, and processes.

### Comments

Like C, Go allows multiline comments bracketed by `/* */` and in-line comments with `//`.
A Descriptive `/* */` should procede the package statement.

### Declaration, Assignment and Types

The general form of a declaration optionally declares the type
and optionally the initial value.
```
var <name> [<type>] [= <initial value>]
```
Actually, it is more general than what is shown here, as 
the `var` statement can declare a list of variables, give them all
types and initialize them.

Within a function, a variable declaration and initialization
can be in a short form using `:=`.  For example,
```
	argc := len(argv)
```
We could have been explicit as follows, or left
off the `int` type.
```
	var argc int = len(argv)
```

Where a variable is declared with an initial value, its type is
inferred from that of the initializer.  Go utilizes type inference
extensively.  It also allows types to be inspected.  In our example,
we print the type of certain values with the `%T` format directives.

Assignment of declared variables is denoted `=` and tuple assignment is permitted, e.g.,
`x, y = 2, 4`.

The [Basic Types](https://tour.golang.org/basics/11) in Go are similar to
those you would find in `stdtypes.h`.  But there is the ambiguity of `char`
is eliminated.  `byte` is an alias for `int8` whereas `rune` represents
a Unicode code point and is an alias for `int32`.

In Go, the type of a variable follows the variable name and is read left to right,
the reverse of C. We see this with the `var` statement above.  But notice that
`argv` is of type `[]string` - which could be read "array of strings".  Recall
that in C, `main` woul declare
```
char *argv[];
or
char **argv;
```
read "argv is an array of pointers to strings" - the type is read right-to-left.

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
allocation.  We are provided a slice. Variable `argv` is a slice of strings.

As in Python, a slice is formed by specifying two indices, a low and high bound,
separated by a `:`, omitted bounds are treated as the origin and last.

In addition to length, a slice has a *capacity*, obtained by `cap(s)`, that is
the number of elements in the underlying arrauy starting from the first element of the
slice.  A slice can be truncarted or  extended by *reslicing*, for example `s = s[:5]`, up to
its capacity.

### Iteration

The first loop in our example shows the Go idiomatic form of iteration
over slices, arrays and maps,
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
or
```
for i := 0; i < argc; i++ { ... }
```
While loops have the condition but no init or post.  This can be expressed as
```
for ; <cond> ; { ... }
```
or more idiomatically as the concise
```
for <cond>  { ... }
```
Declaring the iteration variable in the `for` initializer keeps its
scope local to the for block.  This is one reason that loops which are
otherwise natural as a "while" may utilize the `for` more fully.

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

## High Level I/O - `basicio/main.go`

Our second example brings in basic IO operational concepts, along with some additional syntax
and command line support.
It can be found in [src/basicio/](https://github.com/deculler/gotut/tree/master/src/basicio).

The file defining the package `main`, i.e., the command, is one exception
to the name is last element of the path rule.  Notice that the `words` command
is implemented in `words.go`


### More on [`os`](https://golang.org/pkg/os)

In this example we are opening each file specified on the command
line, reading it, and closing the file.  The `os` pagkage, like
`stdio.h` exports these operating systems abstractions.  The standard file
handles are exported as variables, `os.Stdin`, 'os.Stdout', and 'os.Stderr`.
Note caps.

That package also introduces the `type File` and a set of functions that
support that type, including `os.Open` and `os.Close` which behave
largely as you'd expect.  The type is opaque, it does not export members
and the package provides a set of functions on objects of that type.  But
note that these are not methods of a class, as in Java or Python.  Note also
that the full package.Function is used to name the function.  We import the
package into our namespace, but we still refer to each of its exported
variables and functions with the full `<package>.<symbol>`.

`os` provides the highest level IO interface - [`os.Read`](https://golang.org/pkg/os/#File.Read)
a file to yield a `[]byte` and [`os.Write`](https://golang.org/pkg/os/#File.Write) to
write an `[]byte` in its entirety.  These and other methods on values of `File'
type are accessible as methods on the type as indicated by their declaration.
```
func (f *File) Read(b []byte) (n int, err error)
```

### Error handling

Error handling is cleaner than in C because Go provides multiple return
values.  For example, `os.Open` returns both the result and an error. If
there were an error, the result is of type `*PathError`.

### Pointers

Yes, that right Go has *pointers* - the address of a value in memory
(well, in the virtual address space).  They are much safer than in C.
But, given a variable `x`, `&x` yields a pointer to `x`.  And, given
a pointer `p`, `*p` dereferences it - i.e., access the value that the
pointer points to.  In lots of places where you might make errors in
working with pointers Go will help you out and do the right thing -
but it is good practice to get it right.

You need to understand pointers - a reference to an object is not the same
as the object itself.  (You send snail mail to an address, but it the
address is not the house.  It would shelter you from the elements.)  In
constructing data structures you  will have both objects and references
to objects.  The big difference with respect to C is that pointers in
Go are much more benign.  They are strongly type, they point to something,
you can't do arithmetic on them.  You can store them and you can dereference them
to get to the object they point to.

The potentially confusing nuance in Go is that in refering to a field in a struct
via a pointer to the struct you dereference it *implicitly*.  In C the
messy syntax `(*p).field` is syntactically improved by using `p->field`, whereas
generally `p.field` would be a type error - `p` is a pointer to a struct, not the
struct itself.  In Go one does exactly that.  Instead of the messy `(*p).field`
you are permitted to use `p.field`.  It is generally unambiguous, but can
be at odds with systematic use of `p`, `*p` and `&p`.

### Useful packages: flag and log

You will find that in Go the things you do over and over are well
represented in the standard pacakges.  A good example is parsing the command
line to support flags and such - provided by [`flag`](https://golang.org/pkg/flag/).
It not only separates the flags from the args but produces a help output.
Try `./basiocio -h`.

In this example we are using the simplest functionality - setting the value
of a variable (`wordflag`) if the `-words` flag is present.  The
declaration
```
var wordflag bool
```
illustrates scoping
in Go, as we have variables in the package. in addition to functions.
But note, outside of a function the short assignment declaration cannot be
used.  Every statement begins with a keyword.

Go has lexical scoping that is somewhat more general than C.  It is
routine practice to declare variables in the scope of blocks, such as
`for`, in a function, thereby minimizing the scope.  But, whether
`wordflag` was global or declared within `main`, the `&` allows
`flag.BoolVar` to set its value, since we have passed a reference.

Another example is the p[`log`](https://golang.org/pkg/log/) package.
Here we only use its printing functionality, while it exports a collection
of other methods relevant to captuting information that may be valuable
in diagnosing a problem.

### Functions

Functions in Go are declared with the keyword `func`, followed by the name,
argument list, and return type, before the block of statements forming
the body of the function.  Consistent with the left-to-right reading
of types everywhere in Go, each of the arguments are named followed by
the type and the return type of the function follows the argument list.
(Go also has anonymous functions, but we aren't showing that yet.)

Here we have the `rd_fun` to reading and print each of the tokens in the
input file.  The variables `tokens` and `input` are local to the block,
but `wordflag` is global.  Note the type of `infile`.

### Types and Methods

Go does not have classes, as in Java or Python, but it does provide rich
capabilities for creating abstractions of the form that are frequently
used in systems.  Where abstraction is possible in C, it takes discipline
and adherance to convention.  In Go, the languages significantly helps in
enforcing proper use of abstraction.

A *Method* is a functions defined on a *type* that specifies a special
*receiver* argument - the value of that type that is passed to the
function.  When declaring a method, the receiver argument and its
type appear between the keyword `func` and the name of the function,
mimicing the syntax of method invocation.  Thus, a type and a set of
methods defined on a type are used very much like a class and methods
on objects of the class.  In `rd_fun` we see this with `input` upon which
we may invoke the `Split` method to define the tokenizer for scanning
the input file (words, rather than lines, which the default) and
the `Text` method to obtain the next token as a string.

### Interfaces

The [`bufio`](https://golang.org/pkg/bufio)
package implements buffered IO that is particularly
well suited for processing sequences of text, such as we are doing here.
It provides a handful of types, `Reader`, `Scanner`, `Writer`, and `ReadWriter`.
We are using the `Scanner`, which provides a convenient abstraction
for reading data such as newline-delimited lines of text.  The
`bufio.NewScanner` method returns an object of `Scanner` type,
which we use for extracting a sequence of tokens.

In addition to hiding representation, proper abstraction design permits
reuse, which involves identifying common ensembles of functionality.
These are natural to associate with collections of types, for example numerical operators
on all the kinds of numeric types.  Supporting this commonality gives rise to
polymorphism - the same operations on variuous types.  We see this in manuy places
in the design of systems, for example all the different drivers supporting a
common interface.  Go make interfaces an explicit feature of the language.

An *interface type* is defined as a set of method signatures.  A value of an
interface type can hold any value that implements these methods, i.e., a
disciplined form of polymorphism.  A look at the `bufio` doc reveals that
`NewScanner` takes as argument a value that provides the
[`io.Reader`](https://golang.org/pkg/io/#Reader) interface.
The value returned by `os.Open` is such a thing.  But notice, that `io.Reader` is not
its type.  Its type is `*File` - this is one of many types that implement
the `io.Reader` interface.  In fact, any type that implements the method
```
Read(p []byte) (n int, err error)
```
implements this interface.  A `Scanner` can be manufactured wrapping any of these.
Its declaration defines the functionality it relies upon in the underlying type.
Types in Go do not explicitly "implement" an interface. They do so implicitly, simply
by implementing all the methods in the interface type with the appropriate signatures.

Having something of an interface type does not tell you what it is, it tells
you what it can do.

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

## Storage Management and Buffered IO - `words/words.go`

Our third example,
[`words/words.go`](https://github.com/deculler/gotut/blob/master/src/words/words.go) illustrates
other the richness of the Go storage model and medium level IO interfaces offered
by [`bufio`](https://golang.org/pkg/bufio).  In this we have chosen to read the
input file in a manner similar to C `getc` or `getchar`.  Partly, we have
modernization.  In the days of C, characters fit in 8-bit bytes, period.  The only
question was which character coding (ascii, or perhaps EBCDIC?).  But the world
has moved on to Unicode, there are simply more than 256 charaters in the world.
Go disinguishes between `byte` and `rune` - the first being an
[alias](https://tour.golang.org/basics/11) for `uint8` and
the latter for `uint32` - so you know the sizes of things.
(Even `getc` is of type `int getc( FILE * stream);`)
[A string value is a (possibly empty) sequence of bytes.](https://golang.org/ref/spec#String_types).

### Dynamic storage allocation

In our example, the function `words`, which returns a slice of strings, one for each
"word" parsed from the file creates a `Reader` to access the file, rather than a `Scanner`.
`Reader` allows us to do I/O like `fread` and `fwrite` in C, rather than `scanf` and `printf`
which are akin to `Scanner` above.

In C we would need to either pass in a buffer to
hold the data for the contents of the file, or explictly `malloc` each of the words
and the object pointing to those strings, either a list or an array, which we would
also need to `malloc` or `realloc`.  This is all simpler in Go.  We declare and
initialize `uwords` to be a slice backed by a dynamically allocated array and we
give it a type, length, and initial capacity.  Here the empty string "", which
is of length 0.  (String are well defined objects with length, not the dangerous
null terminated business of C.)  The capacity here is an optimization.  We can go ahead and
allocate enough space for the slice to grow.  If it doesn't exceed that, it wont need
to reallocate itself.  But, either way, we grow the slice with the idiomatic
```
uwords = append(uwords, str)
```
This is really binding `uwords` to a new slice, whether it allocates more storage or not.  Go
has automatic storage reclamation (Garbage collection), like Python and other modern languages
so if `uwords` was the only reference to the old slice, it can all be efficiently modified
in place.  If not, the right thing still happens and you don't have to worry about it.  We
return this dynamically allocated slice of dynamically allocated strings, providing the
most natural expression of what we are up to.

Not that the return type of `getword` is `string`, not a pointer to something.  That's not
too surprising if you think of a string as a pointer to a sequence of characters, rather than
the object itself.  But in general, Go is fine with returning a value that is allocated locally
to a function, what would be "on the stack" in C.  If the lifetime of the value exceed that of the
scope of its declaration, Go allocates it on the heap.  This means we can have closures and
all those other powerful properties of modern languages, with the kind of direct mapping to the
machine that make C so efficient.

### IO Interfaces

The `getword` function declares its argument to be an `interface` - this could be any type
that implements the `Reader` interface, i.e., provides all the methods associated with this
interface.  The one we use here is `ReadByte`, which yields both a value and an err.  Here
we have a very simple parser that skips over all non-alphabetic characters and collects the
following sequence alphabetic characters (what we have chosen to call a "word").  Note
that we simply form that with the append operator, `+` on strings.  But a `byte` is not
a string of length 1.  We form a string out of it by using the type as the operator,
`string(ch)`.

Common IO patterns are wrapped up as interfaces in the
[`io`](https://golang.org/pkg/io/) package.  So the IO picture in Go is
awfully nuanced, with `os`, `fmt`, `bufio`, and 'io` packages.

## Data Structures and abstractions - `wordct`

The next stage in our expedition illustrates the formation of abstractions that
can be used in various applications based on their external interface, independent
of their internal representation.  [`wordct.go`](https://github.com/deculler/gotut/blob/master/src/wordct/wordct.go)
builds a structure containing
the number of occurences of each word in a collection of files using the abstraction
provided in [`wc`](https://github.com/deculler/gotut/blob/master/src/wc_s/wc_s.go).
In C, the wordcount interface would be represented in an include file `wc.h`,
which was included in the main program allowing separate compilation.  That
interface would be implemented in a `wc.c` file, based on a concrete representation.

In Go we don't have the explicit separation of interface and implementation, partly
because separate compilation is no longer and important goal.  The Go tool can
look at the entire application in the build.  Each directory contains the set of
files comprising a package, and applications using the package refer to it
specifically and can only access the variables, functions and types that are
explicitly exported (by using a capital first letter in the name).

We want to explore multiple distinct concrete representations of the same
abstraction, without changing the code that uses the abstraction.  In C,
we might do this with changes to the Makefile.  In Go the build is driven
by the layout of the workspace and the `import` statements.  So we have chosen
to have two distinct packages, `wc_s` which uses a `slice` in the concrete
representation of `WordCounts` and `wc_l` which uses the
[`list`](https://golang.org/pkg/list/) that is analogous to the `list.h`
used throughout Linux (and Pintos) to provide a polymorphic lists in C.
The implementation used by the `main` package (in `wordct.go`) is
specified by which of these packages is imported *and* we rename it to `wc`
so none of the code changes when we change the implementation.  (Another
way of achieving this would have been to create a symbolic link in `src`
to one of the two implementations.)

`wc_s.go` defines two types, `WordCount` and `WordCounts`.  The type and
the struct go together; there is no need for the `typedef`.  The `WordCount`
type and its fields, `Word` and `Count` are exported.  (Note, type in Go follows
name - the reverse of C.)

Three methods are defined on the `WordCount` type, `AddCount`, `Inc`, and `String`.
Notice that all of them declare the same *special receiver* argument type,
`*WordCount`.  Thus, as pointer to a `WordCount` acts like an object reference,
upon which these three methods can be invoked.

`WordCounts` is a more involved data structure that could have a variety of
concrete representations - a collection of `Word`:`Count` bindings where we
can introduce new words dynamically and increase their counts.

In C, the most natural representation would be to introduce a field of
type `*WordCount` in the struct for each entry with which to form a
list.  The `WordCounts` analog would point to the head of the list.
Entries can be easily added onto the front.  This is straigthforward,
but any operations we might want to perform on the list, like sort,
need to be reimplemented for this particular type of list.

Alternatively, we could represent it as an array of pointers to `WordCount`
with `WordCounts` being a pointer to the array.  When we added elements to it,
we would need to `realloc` a larger array.  Since the array is just pointers and
pointers are the same size, regardless of the type of the object they point
to, we could use `(void *)` as the element type in polymorphic functions that we want
to have operate on the array.  We would need to record the lenth of the array in
the `WordCounts` struct, along with the pointer.  And if we wanted to expand it
in chunks as new elements were added, could have another field in the struct,
the size, or capicity, of the currently allocate block of storage under the array.

In Go, the latter approach is naturally supported by the language, without all the
rigamarole.  The
representation is just a slice of `WordCount`, i.e., `[]WordCount`.  It doesn't have
be `[]*WordCount`, although that would be another fine choice because the language
understands the types of the objects.  The slice inherently has a `len` and a
`cap`.  The operation in `AddWord`,
```
wcts.wcs = append(wcts.wcs, wc)
```
either extends the length within the current capacity or realloc's as needed. We
can iterate over the slice with `range` (cf, `Fprint`) or with the index
(cf `Find`) as desired.

### Implementing an interface

Given that our concrete representation of the collection of `WordCount` is
a slice, we could easily pass the slice to the function in the
[`sort` package](https://golang.org/pkg/sort/) for
[sorting slices](https://golang.org/pkg/sort/#SliceStable), which
achieves its polymorphism be receiving a `less` function to
compare the particular elements.

Alternatively, we can take care that `WordCounts` provides the methods that
implement the "data interface" at `sort` relies upon:

```
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```

These are trivial to implement on our concrete representation, since
slice supports indexing.  And notice how the simultaneous assigment
of multiple LHS is elegant.  
By providing these, we can simply invoke `sort.Sort` on a object of
type `WordCounts`.  The sort will invoke our methods as it performs
its work.

### Care in references

Having extolled the virtues of Go idioms, it's time for a little
cautionary note revealing how the kind of understanding what's going
on that you develop in C help deal with obscure confusing things
that might arise.  Note how the `Find` method takes care to iterate
through the indeces of the `Wordcounts.wcs` slice and explictly
creates the reference to the object in that slice,

```
var wc *WordCount = &(wcts.wcs[i])
```
rather than the `range` idiom

```
 for i, wc := range wcts.wcs {
      	...
        }
```
One might argue that this makes absolutely clear that `Find` returns
a pointer to the `WordCount` object in the slice that matches in
the `Word`.  And it does.  If you replace the loop with the
one above, and `return &wc`
you will discover that `AddWord` fails.  It obtains a pointer to an
object that is a copy of the object in the slice.  `AddCount` will
increment its `Count`, but the `WordCounts` is unchanged.

The lesson here is even though higher level languages, like Go (and Python and Java)
take the pain out of pointers, you still need to understand the nuances
of objects, references to objects, copies of objects that are "equal"
but not the same, and references to any of these.  Keep the concept of
pointer in the back of your mind.

## List abstraction in Go - `wc_l`

The `wc_l` package illustrates the discipline of ADTs with a concrete
implementation of `WordCounts` that uses the [`list`](https://golang.org/pkg/list/)
package.  `WordCount` type is unchanged as is the `WordCounts` "interface".
(We could have formally defined this as a `type interface`, but we'll
leave that as an exercise.)

The Go `list` approach is different from the Linux (and Pintos) *list* C macros.  Those
embed the list element (containing the pointers to next and prev)
within the structs that are the values in the list.
Alternatively, a list of structs could be formed with a (void *)
pointer to the objects in the list.  This is effectively what the Go
list does, but it is better supported within the language.  Operating
systems tend not to follow this approach because adding an element to
a list involves allocating the pointer structure.  In the embedded
approach the objects that could potentially be placed on a list
preallocate storage (within themselves) for the list structure.  And
those crazy C preprocessor macros provide a way to go from the list
element to the list entry that surrounds it.

In Go, storage management is automatic, so the external approach is
natural.  The methods for traversing lists are quite similar, as
illustrated in `Fprint`, as are the insertion methods, illustrated in
`AddWord`.  Care must be exercised in creating the list, illustrated
 in `NewWordCount`.  In Go, the builtin `new` is used to allocate
 an object of the specified type, initialized to zero.  Thus,
 `new(WordCounts)` allocates storage for the list (whereas
 `list.New()` would allocate a new list, initialize it and return
 a pointer to it) but we need to invoke the `Init()` method on the
 field of `list` type.

We haven't entirely escaped the tyrany of `(void *)`.  Notice that in
accessing the `Value` of the list element in `Fprint` and `Find` we need
to explicitly declare its type to match that of the object that was passed to
`PushFront` in `AddWord`.  This is a bit of a precarious vulnerability in
Go's otherwise elegant type system.

Many of the sophisticated methods available on the C list abstraction
of Linux, such as ordered insertion, sort, and remove max, are missing
in the Go `list`.  However, we can potentially gain those functions
by implementing some basic interfaces.  To 
use `sort` on our list representation of
`WordCounts` we need to implement the data interface.  The
key missing ingredient is the ability to index directly
into a list.  We address this with the private `locate`
method, which iterates down the list to locate the i-th element.
Not that we return the `*list.Element`, not the `.Value`.  This is
so that `swap` can exchange the values at two points in the list
without changing the structure of the list.  Again, pay attention to
pointers!

With this building blocks, implenting the data interface for sort is
straight forward.  The linear cost of `Less` and `Swap` may
seriously compromise the complexity of `sort` (O(nlogn) comparisons,
but they are not constant time), so it might be preferable to implement
sort directly on the `list` type.

### Go test

Since this indexing business is pretty weird, we use the `go test` tools
that are provided with the language to check it out.  Files with
names that end in `_test` are not used in the build.  They are built and
run by `go test` and can use the [`testing`](https://golang.org/pkg/testing/)
package to perform unit tests.

## Concurrency: Go Routines and Channels - `cwordct`

A forte of Go is the user-level threads that are implemented fully within the language,
called Go routines, rather than grafted on like `pthreads` in C.  This has many
subtle implications for optimizating compilers that are beyond the scope of this
tutorial, but importantly we don't have to sprinkle `volatile` decorators on
variable types, as in C, to make very sublte bugs disappear.

[`cwordct/wordct.go`](https://github.com/deculler/gotut/blob/master/src/cwordct/wordct.go)
uses a separate Go routine to read and parse each of the input files.

```
go readwords(f, name, addr, rdone)
```

They all can
be read in parallel.  The `go` operator preceeding a function call causes the invocation
to run in its own (user-level) thread, i.e., concurrently with the parent thread
and all others that have been spawned.  The thread exists upon completion of the
function call.  Arguments are easily conveyed to the Go routine through the
usual argument passing process, rather constructing a special structure and passing
a function pointer to `pthread_create` as in C.  Go routines operate in a common address
space, each with their own stack, but can access global variables and heap objects
accessible through pointers.

A second key innovation in Go is the use of *channels* to provide
typed message passing - like commonly used for interprocess
communication - among threads in a shared address space.  This
provides for highly structured communication and coordination among
threads, rather than relying on the inherently unstructured
interactions of shared memory.  Clear and comprehensible
communications structures via channels is a design problem, just as is
the creation of data structures.  Here we illustrate two common patterns.

To compute all the `WordCounts` we use a map-reduce approach.  Each `readwords`
go routine is independent, reading and parsing a file to obtain a sequence
of words.  The `adder` go routine accumulates all the word counts using
whatever implementation of `WordCounts`.  The `addr` channel connects the
maps to the reduce.  The `readwords` routines send word strings on the
channel that was passed to them as an aergument:
```
ch_adder <- str
```
the `adder` routine receives all these string on the channel passed to it
by iterating with range.
```
        for str := range ch_adder { // accumulate all the parsed words                    
                wcounts.AddWord(str, 1)
        }
```
The `range` operator terminates when the channel is closed.

Go provides an explicit channel receive as well using the `<-` on the RHS, for example
```
x := <- c
```
but the structured `for ... range` idiom is preferred for stream communication.  Note
the typed nature of channels is critical as it determines the "framing" of the
communication.  Each send is a complete string, as is each receive.  Any type can
be communicated over a channel.  Objects in their entirely are the basic units
of communication.  And since all the threads are in the same address space, the
communicated objects can freely include pointers to shared objects.

Note that with this pattern no additional synchronization is needed to
provide atomic updates to shared data structure.  Only `adder` modifies `WordCounts`.
All the `readwords` are independent.  The only synchronization is the producer-consumer
relationship that is naturally enforced by a channel - receive blocks until
a corresponding send is performed.

The second pattern deals with coordination, rather than data sharing.  It is a
variant of the fork-join pattern, commonly implemented in C with
pthread_exit/pthread_join.  Each of the Go routines are passed a `done` channel
which they send on when they have completed their operation.  More generally, this
could be the return value channel.  The parent thread recieves on this channel
for every go routine it creates.  Thus, the construct:
```
                for i, _ := range args[1:] {
                        n := <- rdone
                        fmt.Println("Done: ", i, n)
                }
                close(addr)     // close the channel, terminating adder
```
detects the completion of all the `readwords`.  This implies that no further
sends will occur on the `addr` channel.  The parent then closes the channel, causing
`range` in `adder` to terminate once all the buffered sends have been received.
Upon its completion, `adder` similarily conveys its completion to the parent with its
`done` channel.

Go provides considerable control over the buffering of a channel.  This can be used
to provide "synchronous" send/receive message passing of CSP (concurrent sequential
processes), asynchronous messaging passing, or various semaphore-like semantics.  But,
those sophisticated mechanations are not necessary for most uses.

## Shared variables and mutext - `mwordct`

A more conventional "shared memory" approach with Go routines is
illustrated by
[`mwordct/woordct.go`](https://github.com/deculler/gotut/blob/master/src/mwordct/wordct.go)
with a synchronized variable of our
slice-based `WordCounts` in
[`wc_sm`](https://github.com/deculler/gotut/blob/master/src/wc_sm/wc_sm.go).

Here the Go routine are a slightly modified variant of `countwords` that
closes its input file after reading, if it is not reading from `os.Stdin`.  This is
a bit of subtle change.  Previously our factoring was to have `main` deal
with command line arguments and file opening/closing.  `countwords` just
read an open file.  But we need to make sure the close happens after the file
has been completely read.  When calling `countwords` as a function, `main` can
close the file upon return.  Not so when each `countwords` is running concurrently
with `main`.

We have retained the use of channels for the fork/join pattern that allows `main`
to detect completion of its go routines, before sorting and printing the results.
Not that we could have done the `close` as part of the completion detection
loop.  But, we would need to match the go routine completion with the correct
`File`.  We could maintain a mapping for this purpose, or we could have defined
the completion channel to be
```
rdone := make(chan *os.File)
```
and send the `infile` value back on the `done` channel.

This takes care of the high level structure, but there remains the need for
synchronization to provide mutual exclusion in modifying the shared
`WordCounts`, since we have all the `countwords` invoking `wc.AddWord`.
We have created a sychronized `wc` package by adding a mutex as a
private field in the `WordCounts`.
```
type WordCounts struct {
	mu sync.Mutex
	wcs []WordCount
}
```
We then use `Lock`/`Unlock` to create a critical section in `AddWord`.  Note
that it might be wise to make `Find` a private method (i.e., `find`) since
it cannot create a critical section and be called within the critical section
of `AddWord`.

Also, we are assuming that the construction phase of `WordCounts`
does not overlap with the sorting and printing.  `Fprint` would not be a problem;
if the slice were changed while we are iterating along it, we would just print the
state of the `WordCounts` when we started.  But the data interface used in sorting
is not synchronized and it is not clear that sort would work correctly if there
were modifications during the sort.  We could obtain the lock in sort, but
`Swap`, `Len`, and `Less` need to be exported for `sort.Sort` to access them,
so they could not be synchronized.  The bulk synchronized view - cosntruction followed
by sorting and rendering is both natural, simple, and predictable.

## Maps - `wc_m`

One of the larger leaps coming to Go from C is the availability of
a key:value data type, common in scripting languages, e.g.,
`dict` in Python, associative arrays in
PhP, etc., but in a compiled, systems oriented language.  Obviously,
this relies on an dynamic storage model with automatic storage reclamation.

[`wc_m/wc_m.go](https://github.com/deculler/gotut/blob/master/src/wc_m/wc_m.go)
rebuilds our `wc` package using a map.

Not how significantly the use of map changes the approach = perhaps enough that
one would consider an very different ADT.  The `WordCounts` type no long even
references `WordCount`.  There isn't an object with the key (`Word`) and the
values (`Count`) in it.  They are simply in the map.

Not that in creating the object (`NewWordCounts`) we have to take some care to
initialize the map, but we also can preallocate storage for its growth.  We don't
need to append to it.  An index assignment does the allocation - within the map or
expanding as needed.

### Nils and zeros

Thus, `AddWord` becomes extremely simple.  The lookup is built in, so `Find` is
unneccessary.  This simple construct is also subtle and deserves clarification.
Note the difference with `Find` above.  If the key is not present in the map, indexing
by the key will return the zero of the value type.  In this case, `int` 0.  Here
that is just what we want.  If it exists and has a count, we add to it.  If not, we create
the entry and initialize the count.

In the `Find` just above, that is not the behavior we want.  It happens to be the case
here that there would be no `word` appearing 0 times.  But, we are treating non-occurence
differently.

Note also that the behavior of this `Find` is subtly different from that in the
others.  It does return a pointer to a `WordCount` with the `word` - but it does not
return a pointer to THE `WordCount` that is in the `WordCounts`.  That aspect of
behavior was not explicit in the interface.  The internal behavior was relied upon
within the implementation of the ADT, e.g., `wc_l`.  Here we don't rely upon it and
it also isn't present.

### Ordering

The concept of `sort` doesn't apply to maps.  They are unordered.  In our example
we treated sort as an inplace operation that defined order.  But the only way that is
visible is through `Fprint`.  So, here we record that the output shoudl be
printed in sorted order. It do the sort, we export the `map` to a `slice`, which
can then be sorted.

We have not implemented the rest of the data interface, allowing this representation
of `WordCounts` to be `sort.Sort`ed because swap would hardly make sense and
in general `less` would be inefficient.


### Anonymous functions, i.e., lambda

The sort requires a `less` function that takes two indexes as arguments.  It doesn't
take auxiliary inputs, such as the array that is being sorted, as is typical in C.
But, Go provides a full capability to build closures, unlike C.  Here we define
a `func` that carries the slice, `swc` in its environment.  Functions are first
class objects, and we've used that here assigning the anonymous function to
a variable, `wc_less`, but we could have just passed the `func(i,j)`
expression in to `sort.Slice` without the variable.

### Range return

Note also in `Fprint` the two distinct uses of range.  `range` of a match
provides the key and value of each element iteratively, whereas `range` of a slice
provides index, value pairs.  This appropriate as the key of the map and the
index of a slice function analogously.  A third use of `range` you saw previously
with a channel.























