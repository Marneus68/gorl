package squareMap

type SquareMap2D struct {
	Content [][]Square
}

func NewSquareMap2D(width, height uint) (squareMap *SquareMap2D) {
	squareMap = new(SquareMap2D)
	squareMap.Content = make([][]Square, height)

	for i := range squareMap.Content {
		squareMap.Content[i] = make([]Square, width)
	}
	squareMap.UpdateSquareLinks()

	return squareMap
}

func (squareMap SquareMap2D) Grow(num uint, dir Direction) {

	squareMap.UpdateSquareLinks()
}

func (squareMap SquareMap2D) Shrink(num uint, dir Direction) {

	squareMap.UpdateSquareLinks()
}

func (squareMap SquareMap2D) UpdateSquareLinks() {
	for i, smyarr := range squareMap.Content {
		for j, smx := range smyarr {
			if i != 0 {
				smx.South = &squareMap.Content[i-1][j]
			}

			if i != (len(squareMap.Content) - 1) {
				smx.North = &squareMap.Content[i+1][j]
			}

			if j != 0 {
				smx.East = &squareMap.Content[i][j-1]
			}

			if j != (len(squareMap.Content[i]) - 1) {
				smx.West = &squareMap.Content[i][j+1]
			}
		}
	}
}

func (squareMap SquareMap2D) AStar(x, y uint) (dir Direction) {
	return NONE
}
