# Encapsulation and Embedding

- Data validation: ensuring that the user data is valid before accepting it

## Setter methods

- Setter methods are methods used to set fields or other values within a defined type's underlying value
- Go's setter methods start named in the form `SetX`

```go
type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) {
	d.Year = year
}
```

- remember to use pointer receivers for setters

## Getter methods

- By convention, a getter method's name should be the same as the name of field or variable it accesses

```go
// Capitalized function so that it can be exporteds
func (d *Date) Year() int {
  return d.year
}
```



## Encapsulation

- Encapsulation: Hiding data in one part of a program from code in another part (=direct access is not allowed)
