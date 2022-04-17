# Interface

- An interface is a set of methods that certain values are expected to have
  - a set of actions you need a type to be able to perform

```go
type myInterface interface {
  methodName(parameter)
  methodName() returnValue
}
```

- Any type that has all the methods listed in an interface definition **satisfies** that interface
- If a type has all the methods declared in an interface, it can be used anywhere without declaration

### Concrete type VS Interface type

**Concrete type**

- Specifies not only methods, but also type

**Interface type**

- Does not define type

### Assign any type that satisfies the interface

```go
type Whistle string
func (w Whistle) MakeSound() {
  fmt.Println("Tweet!")
}

type Horn string
func (h Horn) MakeSound() {
  fmt.Println("Honk!")
}

type NoiseMaker interface {
  MakeSound()
}

func main() {
  var toy NoiseMaker
  toy = Whistle("Toyco Canary")
  toy.MakeSound()	// Tweet!
  
  toy = Horn("Toyco Blaster")
  toy.MakeSound() // Honk!
}
```

- `Horn` and `Whistle` were called without declaration, as part of the interface `MakeSound()`

```go
func play(n NoiseMaker) {
  n.MakeSound()
}

func main() {
  play(Whistle("Toyco Canary")) //Tweet!
  play(Horn("Toyco Blaster")) //Honk!
}
```

- Function parameters can be declared to interface types as well
- You can only call methods of a type that are defined in the interface

## Type assertion

- Type assertion lets you get the concrete type

```go
var noiseMaker NoiseMaker = Robot("Botco Ambler")
var robot Robot = noiseMaker.(Robot) 
```

- *`robot` uses the `NoiseMaker` interface, but it's type is definitely a `Robot` (wink)*

## Error interface

- `error` is an interface!

