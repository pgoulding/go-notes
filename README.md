# Go Notes

## For Loops

```go
func printMap(c map[string]string) {

    for key, value := range c {
        // do something
    }
}
```

## Variables

Established by either explicitly naming type `var example = string "example"` or implicitly assuming type `example:="example"`

Two types of variables, Value and Reference. Value types need pointers to change them in a function, Reference types don't need pointers to change.

| Value Types | Reference Types |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

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

Differences between Maps and Structs
| Map                                                 | Struct                                                        |
| --------------------------------------------------- | ------------------------------------------------------------- |
| All Keys must be the same key type                  | Values can be of different types                              |
| All Values must be of same type                     |                                                               |
| Keys are indexed - we can iterate over them         | Keys don't support indexing                                   |
| Use to reperesnt a collection of related properties | Use to represent a "thing" with a lot of different properties |
| Don't need to know all the keys at compile time     | You need to know all the different fields at compile time     |
| *Reference Type*                                    | *Value Type*                                                  |

### Pointers

A variable that holds the memory address of another variable. See [pointers.md](./pointers.md) for more info.

### Channels


## Value Types

### Structs

Similar to a JavaScript Object or a Python Dictionary. See [structs.go](./structs.go) for more info.

### Interfaces

### Go Routines
