# 06. Slice

- A slice is just like an array

  - same type of elements can be held
  - but lenght can be omitted

  ```go
  var mySlice []int
  ```

- For the most part, slices work just like arrays

  - accessing elements through index
  - initialization with zero values
  - for loops / for...range loops

- A slice literal looks like an array literal, but with the length omitted

  ```go
  mySlice := []int{1,7,10}
  ```

- A slice is a <u>view</u> of the elements of an underlying array

  - slices are built on top of arrays

  - you can make a slice from an existing array using index
  
    ```go
    underlyingArr := [5]string{"a", "b", "c", "d", "e"}
    
    slice1 := underlyingArr[0:3]
    fmt.Println(slice1)  //[a b c]
    ```
  
  - `i : j` = `i` : element in index `j-1`
  
- `append` allows to add onto a slice

  ```go
  slice := []string{"a", "b"}
  fmt.Println(slice)   //[a b]
  
  slice = append(slice, "c")
  fmt.Println(slice)	 //[a b c]
  ```



## Variadic Fuctions(가변 인자 함수)

- A variadic function is a function that can be called with a **varying** number of arguments

- To make a function variadic, write `...` before the type of the **last** function parameter

  ```go
  func myFunc(param1 int, param2 ...string) {
    // function code
  }
  ```

- the last parameter of a variadic function receives te variadic argument as a slice