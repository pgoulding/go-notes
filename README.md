# Go Notes

## Variables

You can establish a variable by using the `var` keyword and explicitly naming type (`var example = string "example"`), implicitly assuming type using **short variable declaration** (`example:="example"`), or by **type inference** using either `=` or `:=` and the type is *inferred* for the value on the rightr hand side.

In go you can also have **same line declarations** where you establish multiple variables on the same line separated by a comma.

```go
    fruit, quantity := "apple", 45
```

Go has 25 [keywords](https://go.dev/ref/spec#Keywords) that can't be used as names for variables:

|           |        |         |             |          |
| --------- | ------ | ------- | ----------- | -------- |
| break     | case   | chan    | const       | continue |
| default   | defer  | else    | fallthrough | for      |
| func      | go     | goto    | if          | import   |
| interface | map    | package | range       | return   |
| select    | struct | switch  | type        | var      |

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

### Constants

Constants in go can only be of type numbers, characters (runes), strings or booleans In go you can also declare a **constant** by using the `const` keyword with an identifier and an optional type specifier. The type specifier is optional because of implicit value from the compiler.

```go
    const identifier [type] = value
    // for example
    const pi float32 = 3.1415
```

## Operators

A symbo that is used to perform logiccal or mathemtaical tasks is called an operator. Go provides thte following built in operators:

- Arithmetic Operators
- Logical Operators
- Bitwise Opertaors

### Arithmetic Operators

The common binary operators `+`, `-`, `*`, and `/` thta exist for both integers and floats in Golang are:

- Addition Operator `+`
- Subtraction Operator `-`
- DIvision Operator `/`
- Modulus Operator `%`
- Multiplication Operator `*`

### Logical Operatos

### Bitwise Operators


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

Constants aree declared like variables, but use the `const` keyword. They must be know at compile time, and usually declared with a static value, however they may be computed as long as that computation *happens at compile time*. Constants must also be of a simple data type, and cannot be of a complex type.

| Allowed | Not Allowed |
| ------- | ----------- |
| int     | slice       |
| string  | map         |
| bool    | struct      |

## Conditionals

`if` statements don't use parenthesis around the condition in Go, and are for the most part very similar to what you would see in other lanaguages.

```go
    if CONDITION {
        return "do something"
    } else if ALTERNATE_CONDITION {
        return "do something else"
    } else {
        return "default condition"
    }
```

One feature of go conditionals is the **initial** statement that can be defined in the scope of the `if` body.

For example, normally you would write a conditional like this:

```go
    validEmail := checkEmailSyntax(email)
    if validEmail >1 {
        fmt.Println("Email is Valid")
    }
```

With Go you can use a bit of syntax sugar to shorten the code to something like this:

```go
    validEmail := checkEmailSyntax(email)
    if validEmail := checkEmailSyntax(email); validEmail > 1 {
        fmt.Println("Email is Valid")
    }
```

This can help remove unnecesary varibales from the parent scope, and only have it apply inside the conditional.

## Loops

```go
func printMap(c map[string]string) {
    // Traditional for loop
    // for initialization; condition; post;
    for i :=0; i < len(c); i++ {

    }

    // Traditional 'while' loops
    for CONDITION {

    }

    // infitinte loops
    for {

    }

    // using the 'range' keyword to go through the length of the map or slice 
    for index, value := range c {
        // in each iteration of the loop, 'range' produces a pair of values, the index and the value of the element at that index.
    }

    // if you don't need the index in the loop youy can use the a blank identifier _ 
    for _, value := range c {

    }
}
```

## Logging to the terminal

Go, like C and other langauges comes standard with a variety of tools to print to the terminal, especially useful for command lines tools and debugging. The most common being `fmt.Printf()`, which produces a formatted output from a list of expressions. It's firtst argument is format string that specifies how other arguments should be formatted. The format of each argument is determined by a consversion character and a letter following a perct sign.

`Printf` has over a dozen such conversions which Go programmers call *verbs*, [A full list can be found here.](https://pkg.go.dev/fmt#hdr-Printing)

| Verb             | Definition                                                                                                                   |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `%v`             | The value in default format                                                                                                  |
| `%d`             | Decimal integer                                                                                                              |
| `%x`, `%o`, `%b` | Integer in hexadecimal, octal or binary                                                                                      |
| `%f`             | Decimal point but no exponent, e.g. 123.456. Precision can be modified by addding a decimal point and number before the f  * |
| `%g`             | Float32, or complex64                       *                                                                                |
| `%e`             | Scientific notation, e.g. -1.234456e+78     *                                                                                |
| `%t`             | Boolean: true or false                                                                                                       |
| `%c`             | Rune (unicode code point) [details](https://exercism.org/tracks/go/concepts/runes)                                           |

*The default precision for `%e`, `%f` and `%#g` is 6, for `%g` it is the smallest number of digits necessary to identify the value correctly. See the docs for more info.

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

### Pointers

A variable that holds the memory address of another variable.

#### Representative in Ram

| Address | Value                       |
| ------- | --------------------------- |
| 001     |                             |
| 002     | person{firstName, lastName} |
| 003     |                             |

Turn `address` into `value` with `*address`
Turn `value` into `address` with `&value`

&variable = Give me the memory address of the value this variable is pointing at.

*pointer = give me the value the memory address is pointing at

```go
func (pointerToPerson *person) updateName(newFirstName string) {
    // *person = This is a type description, meaning we're working with a pointer to a person
    //*pointerToPerson = This is an operator, it mneans we want to manipulate actual the value the pointer is referencing.
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

Either way works, difference seems to be semantic, but speficity IS important as a genreal principle.

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

Here's a table comparing concrete types and interface types in Go:

| **Concrete Type**                                                                                                      | **Interface Type**                                                                                                                                                         |
| ---------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Definition:** A concrete type defines a specific data structure with fields and methods.                             | **Definition:** An interface type defines a set of method signatures without implementing them.                                                                            |
| **Implementation:** Concrete types implement methods that define the behavior for that type.                           | **Implementation:** Interfaces do not implement methods but specify what methods a type must implement to satisfy the interface.                                           |
| **Instantiation:** Can be directly instantiated using literals or constructors. Example: `var a MyStruct = MyStruct{}` | **Instantiation:** Cannot be directly instantiated; you assign a value that satisfies the interface. Example: `var b MyInterface = a` (where `a` implements `MyInterface`) |
| **Memory Allocation:** The size and layout are known at compile time.                                                  | **Memory Allocation:** The size is not known at compile time; it is determined at runtime when a concrete type is assigned to it.                                          |
| **Type Assertions:** Cannot use type assertions, as the type is known and fixed.                                       | **Type Assertions:** Type assertions can be used to extract the concrete type. Example: `if c, ok := b.(MyStruct); ok { ... }`                                             |
| **Usage:** Typically used to define the actual structure of data and provide the concrete behavior.                    | **Usage:** Typically used to define a common behavior across different concrete types without knowing their exact structure.                                               |
| **Method Sets:** Contains all the methods defined on the type.                                                         | **Method Sets:** Contains only the methods that match the interface's method signatures.                                                                                   |
| **Example:** `type MyStruct struct { field1 int }`                                                                     | **Example:** `type MyInterface interface { Method1() }`                                                                                                                    |
| **Common Use Cases:** Data models, structs, specific algorithms, and operations.                                       | **Common Use Cases:** Polymorphism, dependency injection, and abstraction over different types.                                                                            |

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
*A:* Used for communicatin between go routines.
