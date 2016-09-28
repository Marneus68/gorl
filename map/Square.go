package squareMap

type Square struct {
	North *Square
	South *Square
	East  *Square
	West  *Square
	Up    *Square
	Down  *Square

	Content interface{}
}

func New2DSquare(n, s, e, w *Square) (square *Square) {
	square.North = n
	square.South = s
	square.East = e
	square.West = w

	return square
}

func New3DSquare(n, s, e, w, u, d *Square) *Square {
	square := New2DSquare(n, s, e, w)

	square.Up = u
	square.Down = d

	return square
}

func (square Square) GetSquareAt(dir Direction) {

}
