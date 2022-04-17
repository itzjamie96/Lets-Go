# Error Handling

## Defer

- Defer = 연기하다 (미루다
- If you have a function call that you want to ensure is run(no matter what), you can use a  `defer` statement

```go
func Socialize() {
  defer fmt.Println("Goodbye!")
  fmt.Println("Hello!")
  fmt.Println("Nice weather, eh?")
}

func main() {
  Socialize()
}
// Hello!
// Nice weather, eh?
// Goodbye!  --> Executed after Socialize() exited
```

- `defer` ensures a function call happens, even if the calling function exits early

```go
func Socialize() error {
  defer fmt.Println("Goodbye!")
  fmt.Println("Hello!")
  return fmt.Errorf("I don't want to talk.")
  
  fmt.Println("Nice weather, eh?")	//won't be run
  return nil	// won't be run
}

func main() {
  err := Socialize()
  if err != nil {
    log.Fatal(err)
  }
}
// Hello!
// Goodbye!   --> Executed after Socialize() exited
// 2022/04/17 12:51:12 I don't want to talk.
```

- `defer` can only delay function and method calls

## Panic

- When a program panics, the current function stops runnning and the program prints a log message and crashes

### Call stack

- Go keeps a call stack, a list of the function calls that are active at any given point
- When a program panics, a stack trace is included in the panic output
- Deferred calls will still be executed before crashing

## Recover

- Go offers a built in `recover` function that can stop a program from panicking
- you can place a call to recover in a separate function, then defer it

```go
func calmDown() {
  recover()
}
func freakOut() {
  defer calmDown
  panic("oh no")
}
func main() {
  freakOut()
  fmt.Prinln("exiting normally")
}
// exiting normally
```

- When there was a panic, `recover` will return thaever value was passed to `panic`
- `recover` and `panic` can be said to be discouraged in Go (they were made complicatedly so that they'd be used less often)