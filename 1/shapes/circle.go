package shapes

// Circle

type Circle struct {
	shape
	Radius int
}

func (c *Circle) Clone() Shape {

	return CloneCircle(c)
}

func CloneCircle(src *Circle) *Circle {

	dst := NewCircle()

	Clone(dst, src)

	dst.Radius = src.Radius

	return dst
}

func NewCircle() *Circle {

	return &Circle{}
}
