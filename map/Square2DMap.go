package squareMap

type SquareMap2D struct {
	Content [][]Square
}

func NewSquareMap2D(width, height uint) (squareMap *SquareMap2D) {
	return new(SquareMap2D)
}

func (squareMap SquareMap2D) Grow(num uint, dir Direction) {

	squareMap.UpdateSquareLinks()
}

func (squareMap SquareMap2D) Shrink(num uint, dir Direction) {

	squareMap.UpdateSquareLinks()
}

func (squareMap SquareMap2D) UpdateSquareLinks() {

}
