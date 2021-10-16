# Module 1 - Packages, variables, functions and basic types

Go is a compiled, strongly typed-language. Compiled means that in order to run a Go program you need to first compile it for your operating system and then run the output binary. That binary is self-contained, it does not need anything else to run (unlike the JVM languages for example). Strongly-typed means that there are types in Go, these types are known at compile time, and one variable has only one type.

## Packages 
Go code is organized into packages. A package is a folder with go files inside.

Packages are logical grouping of your application. For example your code that deals with the database can live in a package named `database`, and your models can be in `models`. Packages can be nested, if you start having too many models, you can add additional sub-packages, for example `models/people`, `modele/animals`, `models/buildings`. In Go package structure == folder structure, so if we have the following folders structure:

```
models/
	people/
	animals/
	buildings/
	models.go
database/
main.go
```

It means that the packages we have are `models`, `models/people`, `models/animals`, `models/buildings`, `database`.

Each Go program, even the smallest one has at least one package - the `main` one. The `main` package has a `main` function. This is the entry point of our program.

Each Go file can import other packages via the `import` keyword. This is how we reuse code in Go, and how we use code not written by us.

## Functions 
Functions in Go are defined with the `func` keyword. Each function can accept zero or more arguments, and can return zero or more values.

```go
func printAnton() {
    fmt.Println("Anton")
}


func printName(name string) {
    fmt.Println(name)
}


func printNameAndReturnName(name string) string {
    fmt.Println(name)
    return name
}


func printNameAndReturnNameAndRandomInt(name string) (string, int) {
  fmt.Println(name)
  return name, 4 // random number - chosen by a fair dice roll
}
```


## Variables and constants

Variables in Go are defined with the `var` keyword. The type may be omitted, because it is inferred. If we give a type, but no value, the value is the default value for the type. We can omit the var keyword if we use `:=`. We can also define constants via the `const` keyword. However, not all types can be constants (more on that later).

There are few ways to define a variable:

```go
var i int
```

This defines `i` as variable of type `int`.
Since we do not assign value to it, it gets the default value for type `int` which is `0`.

```go
var i int = 5
```

This defines `i` as variable of type `int` and value `5`.

```go
var i = 10
```

This defines `i` as variable with value value `10`.
Even though we don't specify the type, Go is able to infer it as `int`.

```go
i := 15
```

This is the shorthand syntax for defining a value with the `:=` operator.
This way we do not need to use the `var` keyword.

NOTE: This syntax is valid only in function blocks.
You cannot define a global variable this way.

## Types

The basic types in Go are `bool`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, `byte` (alias for `uint8`), rune(alias for `int32`), `float32`, `float64`, `complex64`, `complex128`.

Go also has the `struct` type (more on that later) and `interface` (out of scope of this workshop).

## Task

Given this [scaffold](https://play.golang.org/p/HZUjNTzi2rA), implement the functions `sum`, `min`, `multiply` and `divide`.
