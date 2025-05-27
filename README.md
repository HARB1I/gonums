# gonums
generator for golang enums
```yaml
# gonums -input=../enums.yml -output=../generated_enums

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

type Figure interface {
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

type Size interface {
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
