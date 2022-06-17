# Module 2 - Control flows and build-in data structures

Now that we know about variables and functions we can go on to more complex things like build-in data structures and control flows.

## For

Really common use-case is when we want to do something multiple times until a condition is satisfied. In that case, we are taught not to copy our code multiple times, but to use a loop. And of course, Go has loops. A basic for loop looks like this:

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

This will print all the numbers from 0 to 9.

Also, `for` is go’s while loop.

What in other languages you would write as:

```java
// ❗️❗️❗️ not valid Go code ❗️❗️❗️
while (isTrue) {
    // do something
}
```

in Go you will write as:

```go
for isTrue {
    // do something
}
```

If you want to write an infinite loop you can do:

```go
for true {
    // do something
}
```

OR you can use the short-hand syntax, provided by Go:

```go
for {
    // do something
}
```

There is no do-while loop in Go.

## If

If statements are also pretty standard in Go:

```go
if condition {
    // do something
} else {
    // do something else
}
```

```go
a := 5
b := 6
if a >= b {
    fmt.Println("a is bigger than or equal to b")
} else {
    fmt.Println("a is smaller than b")
}
```

Of course, we can chain multiple if conditions (if-else-if):

```go
a := 5
b := 6
if a > b {
    fmt.Println("a is bigger than b")
} else if a < b {
    fmt.Println("a is smaller than b")
} else {
    fmt.Println("a is equal to b")
}
```

## Switch

Another useful conditional expression is a switch statement. Switch is preferred over if when there are more than 2 scenarios.

Switch statements in Go are also pretty standard with a few caveats.

A switch is written like:

```go
os := "windows"
switch os {
case "windows":
    fmt.Println("you are on windows")
case "linux":
    fmt.Println("you are on linux")
case "macos":
    fmt.Println("you are on macos")
default:
    fmt.Println("you are on something else")
}
```

A `switch` can also be written like this:

```go
os := "windows"
switch os {
case "windows":
    fmt.Println("you are on windows")
case "linux", "macos", "bsd":
    fmt.Println("you are on unix based OS")
default:
    fmt.Println("you are on something else")
}
```

Here, one of the conditions has 3 possible values that will trigger it.

You can use switch exactly the same way you are using if:

```go
switch {
case a > b:
    fmt.Println("a > b")
case a < b:
    fmt.Println("a < b")
case a == b:
    fmt.Println("a == b")
}
```

In this example, we don’t "switch" over a variable, but all the logic is "inside the cases".

Let’s change this example a bit and see if we notice something weird:

```go
a, b := 5, 6
switch {
case a > b:
    fmt.Println("a > b")
case a <= b: // this is true
    fmt.Println("a <= b")
case a < b: // this is also true
    fmt.Println("a < b")
}
```

You might think that this will print both `"a <= b"` and `"a == b"`. This is not true. Go switches terminate upon the first match. If you want the switch to continue checking the conditions after the first match, you need to use the `fallthrough` keyword.

This is the reason why there is no break statement in go switches - because break is implicit.

```go
a, b := 5, 6
switch {
case a > b:
    fmt.Println("a > b")
    fallthrough
case a <= b: // this is true
    fmt.Println("a <= b")
    fallthrough
case a < b: // this is also true
    fmt.Println("a < b")
}
```

Written this way it will print both `"a <= b"` and `"a == b"`.

## Array and slice

One of the basic data structures in every language is the array. In Go an array is defined like this:

```go
var arr [3]string
```

OR

```go
arr := [3]string{}
```

To push data into the array we do it like this:

```go
arr[0] = "first element"
arr[1] = "second element"
arr[2] = "third element"
arr[3] = "fourth element" // compile-time error
```

As you can see, in this case `arr` holds only 3 elements and cannot hold a fourth one. In practice, this is sub-optimal.
Usually when we are working with long lists of data we don’t know how many elements exactly we are going to have.

Go solves this problem by introducing another data-type - slice.
Slice is a dynamically-sized array that can be of any size. It is defined like this:

```go
var s []string
```

OR

```go
s := []string{}
```

To push elements into the slice, we can use the build-in append function:

```go
s = append(s, "first element")
s = append(s, "second element")
s = append(s, "third element")
...
s = append(s, "n+1 element")
```

We can do this as many times as we like. As long as we have enough memory, Go will resize the slice and will save inside all the elements we want it to save.

In practice, slice are used more widely than arrays, because of the size flexibility.
However, the main use-case for arrays is memory-limit, so they are still used in situations where you need to ensure that your data won't go over a given limit.

## Map

Another built-in data structure that is used often is a map. Maps in Go are pretty standard and similar to the ones in other languages.

They are defined as:

```go
var scores map[string]int
```

OR

```go
scores := map[string]int{}
```

In this case, `scores` is a map of `string` keys to `int` values.

We can insert elements into it like this:

```go
scores["Ivan"] = 5
scores["Georgi"] = 3
```

And we lookup elements like this:

```go
ivanScore := scores["Ivan"]
georgiScore := scores["Georgi"]
```

But what happens if an element is not in the map?

```go
peterScore := scores["Peter"] // score = 0
```

Why `0`?
Because `0` is the default value for the `int` data type.
In this case we have to wonder - how do we differentiate between Peter not existing in the map and Peter actually having a score of 0?
Go has a solution for this problem as well.
We can actually write the following code as:

```go
ivanScore, ok := scores["Ivan"] // score = 5, ok = true
peterScore, ok := scores["Peter"] // score = 0, ok = false
```

In this case ok is a boolean value that tells us whether the key “Peter" exists in the map scores or not.

Go does not have a set datatype. If you need to use a set, just use `map[T]bool`.

## Range

So far so good - we are familiar with the build-in data structures.
But a common scenario is when we have a set of data (array, map, etc.) and we want to iterate over all elements of it.
Of course, Go has a way of doing that too. It is called the `range` operator.

To iterate over an array/slice using the `range` operator we do it like this:

```go
for i, el := range arr {
    fmt.Printf("el[%d] is %d", i, d)
}
```

`i` is the index of the element, and `el` is the actual element.

If `arr` has values `[1, 3, 5]` this loop will print out the following values:

```
el[0] is 1
el[1] is 3
el[2] is 5
```

The `range` operator also works on maps:

```go
for key, value := range m {
    fmt.Printf("m[%s] is %d", key, value)
}
```

As seen from this example, the first value is the key, while the second one is the value.

## Task

You are given a list of Star Wars character names.
You need to return a map that counts how many times each character is in the list.
But because Darth Vader and Darth Sidious are evil, we don’t want to count them.

Use this [scaffold](https://play.golang.org/p/X9U3Xk-SX9d).
Implement the `count` function.
After running the code, you should see `SUCCESS` on the screen.

### Additional exercise

Building upon the solution of the previous task exclude all character whose names start with "Darth".

**HINT:** Take a look at the `HasPrefix` function from the `strings` package.
