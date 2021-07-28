# Go

- fast, statically typed, compiled language
  - compiled down to machine level code => doesn't need to be interpreted one line by line 
  - variable types are checked by compiler
  - cannot change variable after declaration
- general purpose language
- built-in testing support
- object-oriented language



## Go in VSCode

- add the "Go" extension
- make a new `.go` file and install the recommended packages



## main()

```go
// a collection of code. mostly called "main"
// "main" => compiled as an executable program
package main

// package from the "go standard library"
// in GOPATH
// "fmt" = format strings
import "fmt"

// entry point of the application
// go will look for the "main" function
// only one "main" function allowed in a package
func main() {

	fmt.Println("Hello, world!")
}
```



### Run your `.go` file

- open a terminal
- enter `go run file.go`



## Variables in Go

### Strings

- variables must be declared with their types

- an error occurs if a variable is **not used** after declaration

  ![image-20210728172457068](Go%20Tutorial.assets/image-20210728172457068.png)

- declaring an **empty** variable is possible => it will just return nothing but an empty space

  ```go
  func main() {
  
  	var apple string = "apple"
  	var orange string
      fmt.Println(apple, orange)		// returns: apple 
  
  }
  ```

- you can declare variables **without** stating the type using `:=`

  ```go
  func main() {
  
  	lemon := "lemon"
  	fmt.Println(lemon)		// works very well!
  
  }
  ```

- it's possible to reassign a new value to an existing variable

  ```go
  func main() {
  
  	lemon := "lemon"
  	fmt.Println(lemon)
  
  	lemon = "change the variable"
  	fmt.Println(lemon)
  
  }
  // ** return **
  // lemon
  // change the variable
  ```

  

### Integers

- specify the bit-size of the integers

  ![image-20210728174831074](Go%20Tutorial.assets/image-20210728174831074.png)

