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
	String() string
}

type square struct{}

func (square) isFigure() {}

func (square) String() string { return "square" }

type box struct{}

func (box) isFigure() {}

func (box) String() string { return "box" }

type cube struct{}

func (cube) isFigure() {}

func (cube) String() string { return "cube" }

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
	String() string
}

type small struct{}

func (small) isSize() {}

func (small) String() string { return "small" }

type medium struct{}

func (medium) isSize() {}

func (medium) String() string { return "medium" }

type large struct{}

func (large) isSize() {}

func (large) String() string { return "large" }

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
		println(v)
	}
}
```
