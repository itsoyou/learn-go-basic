# Learn Go Basics

This is based on Udemy Course `Go - The Complete Guide`, lecture by Maximilian Schwarzmüller.

## Pointers

variables that store value's address instead of the actual value.

e.g.
```
age := 30

in computer memory: value -> 32, address -> 0xc00138475

agePointer := &age -> address -> 0xc00138475
```

Why do we need this?
- avoid unnecessary value copies
    - when you pass a variable to a function as parameter, go creates a copy of the variable.
    Later garbage collector comes and cleans. If values are too big/complicated, might be burden.
    But when you pass value as pointer, this won't happen
- directly mutate values
    - pass pointer to the function instead of variable. same reason as above.
    - function can EDIT the value directly, no return value needed. -> Less Code! (=Less understandable code, unexpected result)


## Structs

Group data together and also methods and functions
defined by "type"

```
type Note struct {

}
```

## Interfaces

Also use "type"
```
type saver interface {

}
```

## Concurrency

### Goroutines

use `go` keyword!

Remember that running a function as goroutine is to run it in a non-blocking way.

```
go greet("nice to meet you!")
```

`fmt.Println()` in greet function will not be printed.
`go` does not return a value, it just dispatches goroutine and doesn't wait for it to complete.

so the program ends before the goroutine finishes.

### Channels

To resolve this problem, we use channel. Channel transmit data, it's a communication device.

```
func slowGreet(phrase string, doneChan chan bool) {
    time.Sleep(3*time.Second)
    fmt.Println(phrase)
    doneChan <- true // sends data to channel
}

func main () {
    done := make(chan bool)
    go slowGreet("nice to meet you!", done)
    <- done // arrow describes the direction of data flow. This end the program after some data came out of the channel.
    // also can do this
    // fmt.Println(<- done)
}
```
This way the program will end after the execution of slowGreet is completed.


