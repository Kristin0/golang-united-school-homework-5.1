package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	var e Point
	e.x = s.start.x + int(s.a)
	e.y = s.start.y + int(s.a)
	return e
}

func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return 4 * s.a
}
