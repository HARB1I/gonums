# gonums
generator for golang enums
```yaml
# gonums -input=enums.yml -output=generated_enums

enum figure:
    square
    box
    cube

enum Size:
    Small
    Medium
    Large
```
OUTPUT
```golang
package figure

// Implements:
//
//	figure.Square
//	figure.Box
//	figure.Cube
type Enum interface {
	isFigure()
	Get() string
}

type square struct{}

func (square) isFigure() {}

func (square) Get() string { return "square" }

type box struct{}

func (box) isFigure() {}

func (box) Get() string { return "box" }

type cube struct{}

func (cube) isFigure() {}

func (cube) Get() string { return "cube" }

var (
	Square = square{}
	Box    = box{}
	Cube   = cube{}
)
```
```golang
package size

// Implements:
//
//	size.Small
//	size.Medium
//	size.Large
type Enum interface {
	isSize()
	Get() string
}

type small struct{}

func (small) isSize() {}

func (small) Get() string { return "small" }

type medium struct{}

func (medium) isSize() {}

func (medium) Get() string { return "medium" }

type large struct{}

func (large) isSize() {}

func (large) Get() string { return "large" }

var (
	Small  = small{}
	Medium = medium{}
	Large  = large{}
)
```
EXAMPLE
```golang
package main

import "test/generated_enums/figure"

func main() {
	foo(figure.Cube)

	bar(figure.Cube)
}

func foo(v figure.Enum) {
	switch v {
	case figure.Cube:
		println("cube")
	case figure.Square:
		println("square")
	case figure.Box:
		println("box")
	}
}

func bar(v figure.Enum) {
	if v == figure.Cube {
		println(v.Get())
	}
}
```
