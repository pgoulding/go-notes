# Pointers


## Represenative in Ram

| Address | Value                       |
| ------- | --------------------------- |
| 001     | person{firstName, lastName} |

Turn `address` into `value ` with `*address`
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
