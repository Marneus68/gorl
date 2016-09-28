package squareMap

type Direction int

const (
	NONE Direction = iota
	NORTH
	SOUTH
	EAST
	WEST
	UP
	DOWN
)

type SquareMap interface {
	Grow(num uint, dir Direction)
	Shrink(num uint, dir Direction)
}
