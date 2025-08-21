package shapes

// Base Prototype

type Shape interface {
	GetX() int
	SetX(int)
	GetY() int
	SetY(int)
	GetColor() string
	SetColor(string)
	Clone() Shape
}

func Clone(dst, src Shape) {
	dst.SetX(src.GetX())
	dst.SetY(src.GetY())
	dst.SetColor(src.GetColor())
}
