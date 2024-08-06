# Go Notes

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

### Pointers

A variable that holds the memory address of another variable. See [pointers.md](./pointers.md) for more info

### Channels

## Value Types

### Structs

Similar to a JavaScript Object or a Python Dictionary. See [structs.go](./structs.go) for more info.

### Interfaces

### Go Routines
