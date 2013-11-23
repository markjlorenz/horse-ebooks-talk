
## Function signatures
Go adds methods to types (classes) by first declaring the type and then the composing function
```go
func (lname TypeAddedTo) FunctionName(arg1 arg1-type, arg2 arg2-type) (retName RetType){}
```

## Goroutines
Runs the supplied function in a separate thread.  Does not block.
```go
go myFuntionName()
```

## Channels
Similar to filesystem pipes, or Ruby's thread-safe queue
```go
from_splitter = make(chan string)
from_splitter <- "hello world"
```

## Arrays
```go
splits []chan string  // splits is an array of channels of strings
```

## Objects
Go is not an OO languages the way Ruby is.  There is no class based, only composition and interfaces.
```go
// Define a `Chum` type with some fields
type Chum struct {
  split_id      int
  splitter      *Splitter
  from_splitter chan string
}

Chum{} // initialize a new `Chum`
```

## Blocks
Go doesn't have them.  Of course.
Use functions like you would in javascript.  Useful for scoping varaibles without explicity passing them.
```go
//  `Times` takes a function, `block` that receives one string argument
func Times(num_times int, block func(msg string)) {
  iter := make([]int, num_times)
  for _ = range iter {
    block("hi")
  }
}

// Prints "hihihihihi"
Times(5, func(msg string){
  fmt.Fprint(msg)
})
```

## When to reach for Go
- If you're working with concurrency.
- If you want to distribute to non-ruby developers

## When to reach for Ruby
- Communicating with other programmers
- Collections of mix-typed objects
- Meta-programming
- Confident code
