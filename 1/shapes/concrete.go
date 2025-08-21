package shapes

// Concrete Shape

type shape struct {
	X     int
	Y     int
	Color string
}

func (s *shape) GetX() int {
	return s.X
}

func (s *shape) GetY() int {
	return s.Y
}

func (s *shape) GetColor() string {
	return s.Color
}

func (s *shape) SetX(v int) {
	s.X = v
}

func (s *shape) SetY(v int) {
	s.Y = v
}

func (s *shape) SetColor(v string) {
	s.Color = v
}
