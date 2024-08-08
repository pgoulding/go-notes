# Go Notes

## Variables

Established by either explicitly naming type `var example = string "example"` or implicitly assuming type `example:="example"`

There are two types of variables, __Value__ and __Reference__, Value types need pointers to change them in a function, while Reference types don't need pointers to change.

| Value Types | Reference Types |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

## For Loops

```go
func printMap(c map[string]string) {

    for key, value := range c {
        // do something
    }
}
```


## Reference Types

### Slices
*Reference Type*
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
*Reference Type*
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

### Pointers (*Reference Type*)

A variable that holds the memory address of another variable.

#### Represenative in Ram

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

### Channels


## Value Types

### Structs
*Value Type*
Similar to a JavaScript Object or a Python Dictionary. See [structs.go](./structs.go) for more info.

### Interfaces


### Go Routines


### Q & A

_Q:_ What is a Pointer in Go?
_A:_ A variable that holds the memory address of another variable.

_Q:_ How do yopu declare a pointer in Go?
_A:_ Using the `*` operator followed by the type of variable it will point to. e.g `var p *int`

_Q:_ How do you get the address of a pointer in Go?
_A:_ Using the `&` operator. For example:

```go
var x int = 10
var p *int = &x
```

_Q:_ Whenever you pass an integer, float, or string into a function what does Go do with that argument?
_A:_ It creates a copy of each arguement and these copies are used inside of the function.

_Q:_ What is the `&` operator used for?
_A:_ Turning a value into a pointer

_Q:_ WHen you see a `*` operator in front of a pointer , what will it turn the pointer into?
_A:_ A Value.