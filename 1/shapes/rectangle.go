package shapes

// Rectangle

type Rectangle struct {
	shape
	Width  int
	Height int
}

func (r *Rectangle) Clone() Shape {

	return CloneRectangle(r)
}

func CloneRectangle(src *Rectangle) *Rectangle {

	dst := NewRectangle()

	Clone(dst, src)

	dst.Width = src.Width
	dst.Height = src.Height

	return dst
}

func NewRectangle() *Rectangle {

	return &Rectangle{}
}
