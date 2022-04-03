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

