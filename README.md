# Go Notes


This Repo is a place for my notes on Go(lang) and is based an accumulation of books, courses, projects, articles and threads.

Go Source code is stored in *.go* files. Their filenames consist of lower-case letters like `main.go` if the name consists of multiple parts, they are seperated by underscores (`_`) like `tcp_transport.go`. Filenames cannot contain spaces or any other speical characters.

## Identifiers

An *identifier* is a name assigned by the user to a program element like a *variable*, *function*, *template*, *struct*, etc... Nearly all things in Go have a name or *identifier*. Valid identifiers begin with a letter or underscore and are followed by 0 or more letters or unicode digits, can not start with a digit, have an [operator](#operators), or be a keyword. Go has 25 [keywords](https://go.dev/ref/spec#Keywords) and 36 [predeclared identifiers](https://go.dev/ref/spec#Predeclared_identifiers) that can't be used for names. You can also use an underscore (`_`) or as it's called in go a *blank identifier*, however its value is discarded and cannot be used in the code that follows, on use case for that being [loops](#loops).

| Valid Identifiers | Invalid Identifiers |
| ----------------- | ------------------- |
| x56               | else                |
| group1            | 1group              |
| _group2           | group1+group2       |

Go Keywords:

|           |        |         |             |          |
| --------- | ------ | ------- | ----------- | -------- |
| break     | case   | chan    | const       | continue |
| default   | defer  | else    | fallthrough | for      |
| func      | go     | goto    | if          | import   |
| interface | map    | package | range       | return   |
| select    | struct | switch  | type        | var      |

An identifier can be exported to permit access to it from another package by having the first character of the identifer be reporesented by an uppercase letter, and if the identifier si declared in the either the package block, field name, or method name. All other identifiers are not exported.

## Variables

You can establish a variable by using the `var` keyword and explicitly naming type (`var example = string "example"`), implicitly assuming type using **short variable declaration** (`example:="example"`), or by **type inference** using either `=` or `:=` and the type is *inferred* for the value on the rightr hand side.

In go you can also have **same line declarations** where you establish multiple variables on the same line separated by a comma.

```go
    fruit, quantity := "apple", 45
```

There are many types in the golang ecosystem and it can get confusing to know which ones to use. Unless you have a very good reason to, stick to the following types:

- bool
- string
- int
- uint
- byte
- rune
- float64
- complex128

There are two types of variables, **Value** and **Reference**, Value types need pointers to change them in a function, while Reference types don't need pointers to change. Some Examples being:

| Value Types | Reference Types |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

### Type Sizes

Ints, uints, floats and complex numbers all have type sizes.

The standard sizes that should be used unless you have a specific need are:

int
uint
float64
complex128

| Whole Numbers | Positive Whole Numbers | Decimal Numbers | imaginary numbers |
| ------------- | ---------------------- | --------------- | ----------------- |
| int           | uint                   | float32         | complex64         |
| int8          | uint8                  | float64         | complex128        |
| int16         | uint16                 |                 |                   |
| int32         | uint32                 |                 |                   |
| int64         | uint64                 |                 |                   |
|               | uintptr                |                 |                   |

The standard sizes that should be used unless you have a specific need are:

- int
- uint
- float64
- complex128

### Constants

Constants in go can only be of type numbers, characters (runes), strings or booleans. Constants are declared like variables, but use the `const` keyword. In go you declare a **constant** by using the `const` keyword with an identifier and an optional type specifier. The type specifier is optional because of implicit value from the compiler. They must be know at compile time, and usually declared with a static value, however they may be computed as long as that computation *happens at compile time*. Constants must also be of a simple data type (string, boolean, numeric), and cannot be of a complex type.

```go
    const identifier [type] = value
    // for example
    const pi float32 = 3.1415
```

| Allowed | Not Allowed |
| ------- | ----------- |
| int     | slice       |
| string  | map         |
| bool    | struct      |

## Operators

A symbol that is used to perform logical or mathematical tasks is called an operator. Go provides thte following built in operators:

- Arithmetic Operators
- Logical Operators
- Bitwise Opertaors

### Arithmetic Operators

The common binary operators `+`, `-`, `*`, and `/` that exist for both integers and floats in Golang are:

- Addition Operator `+`
- Subtraction Operator `-`
- Division Operator `/`
- Modulus Operator `%`
- Multiplication Operator `*`

### Logical Operators

The following are Logical operators in Go:

- Equality Operator `==`
- Not-Equal Operator `!=`
- Greater-than Operator `>`
- Less-than Operator `<`
- Greater-than or Equal to `>=`
- Less-than or Equal to `<=`

In addition to these Logical operators, Go has three boolean logical operators:

- AND operator `&&`
- OR operator `|`
- NOT operator `!`

### Bitwise Operators

Following are some bitwise operators:

- Bitwise AND operator `&`
- Bitwise OR operator `|`
- Bitwise XOR operator `^`
- Bit CLEAR operator `&^`
- Bitwise COMPLEMENT operator `^`

There are other two major bitwise operators used for shifting:

- Left shift operator `<<`
- Right shift operator `>>`

## Conditionals

- `if-else` construct
- `switch-case` construct
- `select` construct

### If-Else

The `if` tests a conditional statement, that statement can be locgical or boolean. If the statement evaluates to **true**, the body of the statements between `{}` after the `if` is executed, and if it is false, these statements are ignored, and the statement after the `}` is executed.

The `{}` braces are mandatory even when there is only one statement in the body. The `{` after the `if`, `else`, `else if` must be on the **same line**, and the `else`, `else if` keywords must be on **same line** as the closing `}` of the previous part of the structure.

The parentheses `()` around the conditions are not needed. However, for complex conditions, they may be used to make the code clearer, and are for the most part very similar to what you would see in other lanaguages.

```go
    if CONDITION {
        return "do something"
    } else if ALTERNATE_CONDITION {
        return "do something else"
    } else {
        return "default condition"
    }
```

It is almost always better to test for true or positive conditions, but it is possible to test for the reverse with `!` (not).

```go
    if !(condition) {

    }
```

One feature of go conditionals is the **initialization statement** that can be defined in the scope of the `if` body.

For example, normally you would write a conditional like this:

```go
    val:= 10
    if val > max {
        fmt.Println("Value is greater and maximum allowed")
    }
```

With Go you can use a bit of syntax sugar to shorten the code to something like this:

```go

    if val:=10; val > max {
        fmt.Println("Value is greater and maximum allowed")
    }
```

This can help remove unnecesary varibales from the parent scope, and only have it apply inside the conditional. Just remember that in this case `val` is only accessible in the scope of this conditional. If you try and acces it anywhere else in this program, the compiler will give you an error `undefined: val`

### Switch-Case

The keyword `switch` can be used instead of long `if` statements that compare a variable to different values. The `switch` statement is a multiway brnach statement that provides a an easy way to transfer flow of execution to diferent parts of code based on the value. If you want to explicitly allow the code from **case 2** after **case 1** you'll need to use the `fallthrough` keyword at the end of the **case 1** branch.

**Fallthrough** can be used in a hierarchy of cases where at each level something has to be done iin addition to the code already executed at higher classes, and a default action also has to be executed.

There are a few rules you have to consider when writing a switch statement.

- Each `case` branch is exclusive
- They are tried from first to last, so you should place the most probable cases first.
- More than one value can be tried in a case statement
- When `case` statements end with a `return` keyword, there also has to be a `return` statement after the end of `}` of the `switch`.
- The optional `default` case is executed when no value is found to match the other cases. It acts like al `else` clause in `if-else` statements.
- Any type that supports the equality comparison operator, such as ints, strings or pointers, can be used like an if statement.
- A `switch` can also contain an initialization statement, like the `if` construct, by adding a semicolon at the end.

```go
    // example switch structure
    switch exampleVar {
        case val1: // these must be of the same type, or expressions evaluating to that type
            // do something
        case val2:
            // do something else
            // can be multiline
        case val3, val4: // more than one value can be tested in a comma-separated list
            // do something
        case val5 > 10: fallthrough
            // do something *if* 
            // and then do the next thing
        case val6:
            // this is now also evaluated, even after val5 is executed as true.  
        default: // optional
            // Default catch all if no other statements evaluates true
    }
```

## Iterative Loops

- `for` keyword
- `break` keyword
- `continue` keyword

In Go the `for` statement exists to repeat a body of statements a number of times, one pass through the body is called an **iteration**. In every iteration a condition has to be checked to see whehter a loop should stop. If the exit condition becomes true, the loop is terminated. If we want to specifically chande the flow of execution we have two statements `break` and `continue`. A `break` statement always breaks out of the innermost structure in which it occurs. It can be used in any for-loop, and also a `switch` or `select` statement.

There are two methods to control iteration:

- Counter-Controlled
- Condition-Controlled

The simplest form is the **Counter-Controlled**: `for initialization; condition; modification { }`

```go
    for i :=0; i < 10; i++ {

    }
```

The **Condition-Controlled** iteration contains no headed and is generally considered a `while-loop` in other langauges.

```go
    for condition {

    }
```

You can also have an **infinite loop** by not giving an intialization/condition/modification. Always check that you have an exit condition to avoid an endless loops. Break only exits from the loop while return exits from the function in which the loop is coded.

```go
    for {
        if(exitCondition) {
            break
        }
    }
```

The `for range` iterator construct in Go can be useful in many contexts, for example looping over every item in a collection.

```go
    // using the 'range' keyword to go through the length of the map or slice 
    c := ["yellow","green","blue","red"]
    for idx, value := range c {
        // in each iteration of the loop, 'range' produces a pair of values, the index and the value of the element at that index.
        fmt.Printf("The color %s is at position %d", value, idx)
    }
    
    // however if you don't need the index in the loop youy can use the a blank identifier _ 
    for _, value := range c {
        fmt.Printf("The color %s", value)
    }
```

## Labels & goto

A label is a sequence of characters that identifies a location within your code. A code line which starts with a `for`, `switch` or `select` statement can be preceded with a label `IDENTIFIER:`

In go there is a keyword `goto` whcih has to be followed by a label name. `goto IDENTIFIER`

## Functions in Go

Every program consists of several functions of the basic core functionality, the main prupose of which is to break a large problem into a smaller one, or the same task several times. When writing code you should always honor the [Don't Repeat Youyrself (DRY) principles][https://en.wikipedia.org/wiki/Don%27t_repeat_yourself], meaning the code which performs a certain task should only appear once in a program.

A function ends when it has executed its last statement before the ending "`}`" or when it executes the `return` statement.

There are three types of functions in Go:

- Normal Functions with an Identifier
- Anonymous AKA Lambda Functions
- Methods

### Parameters and Return Values

Parameters can be **actual** paramaters or **formal** parameters, the difference between them being actual values are the values passed to the function when it is invoked, while the formal values are the varuiables that are defined by the function when it is called.

```go
func main(){
    name := "Patrick"
    printGreeting(name) // here name is being passed as the ACTUAL parameter.
}


func printGreeting(name string) { // Here name is the FORMAL parameter.
    fmt.Println("Hello, ", name)
} 
```

Returing values from a function is done by invoking the keyword `return`, the code following it will be executed. The default way to call a function in Go is to pass a variable as an argument to a function by value. A copy is made of tha variable and the data in it, and the function it is passed to works with that data while leaving the original data unmodified. This is called **pass by value** If you want the function to be able to change the value of the argument itself you have to pass the memory address of the variable the the `&` operator, this is called **pass by reference**; Effectively a [pointer](#pointers) is then passsed to the function. If the variable that is passed is a pointer then the pointer is copied, not the data that it is pointing to. However through the pointer, the function can change the original value. Passing a pointer is in almost all cases cheaper than making a copy of the object.

Some functions just perform a task and do not return values. They perform what is called a **side-effect**, like printing to the console, sending an email, logging an error and so on...

When returning a variable from a function you have to list is type in the same line after the parameters.

```go
func multiplyByTwo(paramVal int) int {
    return paramVal *2
}
```

Go also supports **Named Return Variables** where you can name the variables you want to return at the top of the function, and they are initialized as nil, or the default value for it's type.

```go
func main() {
    fmt.Println(multiplyBy(5)) // returns 10
}

func multiplyBy(paramVal int) (retVal int) { // named variable retVal initialized in top level    
    retVal = paramVal * 2
    return
}
```

If you have an unkown number of parameters you can pass them using [variadic functions](https://gobyexample.com/variadic-functions). In this case the last parameter of a function is followed by **...type**, this indicates that the function can deal with a variable number of paramaters of that type.

```go
func main(){
    myVariadicFunction("Bob", "Bill", "Anne")
}

func myVariadicFunction(name ...string) {
    for _, n := range name {
        fmt.Println("Hello,", n)
    }
}

// Prints:
// Hello, Bob
// Hello, Bill
// Hello, Anne
```

Use of the `defer` keyword can allow us to postpone exeuction of a statement or a function until the end of the calling function. Defer then executes something when the closing function returns. When multiple `defer`'s are used in a function they execute in LIFO order. The defer can be used to guarantee that certain tasks are performed before we return from a function.

Tasks that may need a `defer`:

- Closing a file stream
- Unlocking a locked resource (a mutex)
- Closing a database connection

## Testing in Go

Some functions in Go are defined so that they return *two* results, on is the falue, and the other is the status of the execution. For example the function would return both the value, and `true` in case of a succesful execution, or return the value (probably `nil`) and `false` in case of an unsuccessful execution. Other functions return an error-variable, where in case of a successful execution it will likely return `nil`, otherwise it contains the error informatiion. Commonly this is called the **comma, ok** pattern.

```go
    var notAnInt string = "Blue"
    
    val, err := strconv.Atoi(notAnInt)

    if err != nil { // if an error was found
        fmt.Printf("Program stopping with error %v", err)
        os.Exit(1) // exit the program with a error code other than 0
    }
```

To implement your own error checking in Go, we can use the[`errors` package](https://pkg.go.dev/errors).

[Example](./example_errors.go)


## Reference Types

### Slices

 | Arrays                   | Slices                                                     |
 | ------------------------ | ---------------------------------------------------------- |
 | Primitive Data Structure | Dynamic Data Structure consistitng of a slice and an Array |
 | Can't be resized         | Can Grow and Shrink                                        |
 | Rarely used directly     | Used 99% of the time for lists of elements                 |

A Slice initializes an Array, with three elements inside of it:

- the Pointer to the head
- the capacity
- and the length

When you create a slice it initalizes an Array, and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array (pointer to the head).

### Maps

A **map** holds a set of key/value pairs and provides constant-time (or O(1) in Big-O notation) operations to store,retrieve, or test for an item in the set. The key may be of any type whose values can be compared with *==*, strings being the most common example; the value may be of any type at all.

All keys and values must be same same type. Keys and their values can be different types, but all the Keys must be the same type as each other, same for the values.

Can access values by square brackets

```go
    colors := map[string]string{
        "red":   "#fff0000",
        "green": "#00ff00",
        "white": "#ffffff",
    }
    greenHex := colors["green"]
    fmt.Println(greenHex)
    // should print #00ff00
```

Why would you use a Map over a struct?

#### Differences between Maps and Structs

| Map (*Reference Type*)                              | Struct (*Value Type* )                                        |
| --------------------------------------------------- | ------------------------------------------------------------- |
| All Keys must be the same key type                  |                                                               |
| All Values must be of same type                     | Values can be of different types                              |
| Keys are indexed - we can iterate over them         | Keys don't support indexing                                   |
| Use to reperesnt a collection of related properties | Use to represent a "thing" with a lot of different properties |
| Don't need to know all the keys at compile time     | You need to know all the different fields at compile time     |

## Pointers

Pointers store values in memory, and each memory block (or word) has an *address*, which is suually represnted as a hexadecimal number. Go has the *address-of operator* `&`, which when placed before a variable, gives us the memory address of that variable, that address is a special data type called a pointer. You cannot take the address of a literal or constant. The size of a pointer is 4 bytes on a 32 bit machine and 8 bytes on a 64 bit machine, regardless of the size of the value they point to. Using a pointer to refer to a value is called indirection. A newly declared pointer which has not been assigned to a variable has the *nil* value. A pointer variable is often abbreviated as *ptr*.

One advantage of pointers and where you might want to use them is that you can pass a refernce to a variable instead of passing a copy of that varibale itself. Pointers are cheap to pass (only 4 to 8 bytes), so when the program has to work with variables that occupy a lot of memory, many variables, or both, working with pointers can reduce meory usage and increase effiecncy. Pointer variables persist in memory as long as at least 1 pointer is pointing to them, so their lifetime is independent of the scope in which they are created.

On the other hand, because a pointer causes what it calls an indirection (a shift in the processing to another address), using them unnecesarily could cause a performance decrease. Pointers can point to other pointers and this nesting can go arbitrarily deep so that you can have multiple levels of indirection.

While Go has the concept of pointers, it doesn't allow calculations with them (pointer arithmetic); as that often causes faulty memory access and fatal crashes in C and other low level languages, this feature makes Go *memory safe*.

### Representative in Ram

| Variable | Value                       | Memory Address |
| -------- | --------------------------- | -------------- |
| i1       | 5                           | 0x1000         |
| intPtr   | 0x1000                      | 0x2000         |
|          |                             | 0x0010         |
|          |                             | 0x0020         |
| alex     | person{firstName, lastName} | 0x0030         |

Turn `address` into `value` with `*address`
Turn `value` into `address` with `&value`

&variable = Give me the memory address of the value this variable is pointing at.

*pointer = give me the value the memory address is pointing at

```go
func (pointerToPerson *person) updateName(newFirstName string) {
    // *person = This is a type description, meaning we're working with a pointer to a person
    // *pointerToPerson = This is an operator, it mneans we want to manipulate actual the value the pointer is referencing.
    (*pointerToPerson).firstName = newFirstName
}
```

You can either use `alexPointer := &alex` to specify the value, or Go allows you to shortcut the pointer if your receiver function designates it as a type of pointer (using `*` to specify pointer).
eg:

Getting the memory address from the variable

```go
alexPointer := &alex
alexPointer.updateName("Alexander")
```

or just use the value itself:

```go
alex.updateName("Alexander")
```

As long as the receiver function you're passing it to has the `*` pointer:

```go
func (pointerToPerson *person) updateName(newFirstName string) 
```

Either way works, difference seems to be semantic, but speficity IS important as a general principle.

## Value Types

### Structs (*Value Type*)

Similar to a JavaScript Object or a Python Dictionary.

### Interfaces

Interfaces are tough to understand, learning to read the standard library documentation is good place to start. Writing your own interface is hard and require experience. Interfaces are used as a contract to help us manage types.

#### Syntax to define interface

```go
// Interface name = bot
type bot interface {
    //         Argument Types (optional)
    getGreeting(string, int) (string, error)
    //    ^                         ^
    // Function Name           Return Types
}
```

- Interfaces are **not** generic types, Go doesn't support these.
- Interfaces are 'implicit', no need to declare that
- Interfaces are a contract to help us manage types
  - Not a built in test, Garbage in -> Garbage Out

Comparing concrete and interface types in Go:

|                        | **Concrete Type**                                                                                   | **Interface Type**                                                                                                                                      |
| ---------------------- | --------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Definition:**        | A concrete type defines a specific data structure with fields and methods.                          | An interface type defines a set of method signatures without implementing them.                                                                         |
| **Implementation:**    | Concrete types implement methods that define the behavior for that type.                            | Interfaces do not implement methods but specify what methods a type must implement to satisfy the interface.                                            |
| **Instantiation:**     | Can be directly instantiated using literals or constructors. Example: `var a MyStruct = MyStruct{}` | Cannot be directly instantiated; you assign a value that satisfies the interface. Example: `var b MyInterface = a` (where `a` implements `MyInterface`) |
| **Memory Allocation:** | The size and layout are known at compile time.                                                      | The size is not known at compile time; it is determined at runtime when a concrete type is assigned to it.                                              |
| **Type Assertions:**   | Cannot use type assertions, as the type is known and fixed.                                         | Type assertions can be used to extract the concrete type. Example: `if c, ok := b.(MyStruct); ok { ... }`                                               |
| **Usage:**             | Typically used to define the actual structure of data and provide the concrete behavior.            | Typically used to define a common behavior across different concrete types without knowing their exact structure.                                       |
| **Method Sets:**       | Contains all the methods defined on the type.                                                       | Contains only the methods that match the interface's method signatures.                                                                                 |
| **Example:**           | `type MyStruct struct { field1 int }`                                                               | `type MyInterface interface { Method1() }`                                                                                                              |
| **Common Use Cases:**  | Data models, structs, specific algorithms, and operations.                                          | Polymorphism, dependency injection, and abstraction over different types.                                                                               |

### Go Channels & Routines

Go rountines are initialized by using the `go` keyword before a function call.

By default Go runs on only one CPU core.

#### Concurrency is NOT Parallelism

Concurrency: Dealing with a lot of things at once.
Parallelism: Doing a lot of things at once.

Concurrency is using single core efficiently (sequentially, no downtime as much as possible), while Parallelism is using multiple cores at once.

Using Go Routines is typically applying concurrency, e.g. using a single core but using it efficiently.

However Parallelism is also available in Go. If you have 4 cores, you could set `runtime.GOMAXPROCS(4)` ([docs](https://pkg.go.dev/runtime#GOMAXPROCS)) to make use of all 4 cores, by Default it is already the number of cores as reported by your system. So Go does use excplicit concurrency via Go Routines and implicit Parallelism, however its parallelism is abstracted away and is managed by the runtime itself.

To learn more you can watch this [talk](https://go.dev/blog/waza-talk) from Rob Pike, the creator of Go.

## Logging to the terminal

Go, like C and other langauges comes standard with a variety of tools to print to the terminal, especially useful for command lines tools and debugging. The most common being `fmt.Printf()`, which produces a formatted output from a list of expressions. It's firtst argument is format string that specifies how other arguments should be formatted. The format of each argument is determined by a consversion character and a letter following a perct sign.

`Printf` has over a dozen such conversions which Go programmers call *verbs*, [A full list can be found here.](https://pkg.go.dev/fmt#hdr-Printing)

| Verb             | Definition                                                                                                                   |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `%v`             | The value in default format                                                                                                  |
| `%d`             | Decimal integer                                                                                                              |
| `%s`             | String value, or the uninterpreted bytes of the string or slice                                                              |
| `%t`             | Boolean: true or false                                                                                                       |
| `%f`             | Decimal point but no exponent, e.g. 123.456. Precision can be modified by addding a decimal point and number before the f  * |
| `%x`, `%o`, `%b` | Integer in hexadecimal, octal or binary                                                                                      |
| `%g`             | Float32, or complex64                       *                                                                                |
| `%e`             | Scientific notation, e.g. -1.234456e+78     *                                                                                |
| `%c`             | Rune (unicode code point) [details](https://exercism.org/tracks/go/concepts/runes)                                           |

*The default precision for `%e`, `%f` and `%#g` is 6, for `%g` it is the smallest number of digits necessary to identify the value correctly. See the docs for more info.

| Escape Characters | Description        |
| ----------------- | ------------------ |
| `\n`              | newline            |
| `\r`              | carriage return    |
| `\t`              | tab                |
| `\u`, `\U`        | unicode characters |

For compound objects, the elements are printed using these rules, recursively, laid out like this:

| Compound Object   | Representation in terminal       |
| ----------------- | -------------------------------- |
| struct:           | {field0 field1 ...}              |
| array, slice:     | [elem0 elem1 ...]                |
| maps:             | map[key1:value1 key2:value2 ...] |
| pointer to above: | &{}, &[], &map[]                 |


### Q & A

*Q:* What is a Pointer in Go?
*A:* A variable that holds the memory address of another variable.

*Q:* How do you declare a pointer in Go?
*A:* Using the `*` operator followed by the type of variable it will point to. e.g `var p *int`

*Q:* How do you get the address of a pointer in Go?
*A:* Using the `&` operator. For example:

```go
var x int = 10
var p *int = &x
```

*Q:* Can you create a pointer with no data in it?
*A:* No, this is called a *nil* pointer and will cause the program to crash with the error: `panic: runtime error: invalid memory address or nil pointer dereference`

*Q:* Whenever you pass an integer, float, or string into a function what does Go do with that argument?
*A:* It creates a copy of each arguement and these copies are used inside of the function.

*Q:* What is the `&` operator used for?
*A:* Turning a value into a pointer

*Q:* When you see a `*` operator in front of a pointer , what will it turn the pointer into?
*A:* A Value.

*Q:* To say that a type satisfies an interface means that...
*A:* The type implements all of the functions contains in the interface definition.

*Q:* Types that implement the Reader interface are generally used to...
*A:* Read information from an outside source into our application.

*Q:* What is a go routine (short answer)?
*A:* A separate line of code execution that can be used to handle blocking code.

*Q:* Whats the purpose of a channel (short answer)?
*A:* Used for communication between go routines.
