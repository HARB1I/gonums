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
}

type square struct{}

func (square) isFigure() {}

type box struct{}

func (box) isFigure() {}

type cube struct{}

func (cube) isFigure() {}

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
}

type small struct{}

func (small) isSize() {}

type medium struct{}

func (medium) isSize() {}

type large struct{}

func (large) isSize() {}

var (
	Small  = small{}
	Medium = medium{}
	Large  = large{}
)
```
EXAMPLE
```golang
package main

import "test/test/generated_enums/figure"

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
		println("cube")
	}
}
```
