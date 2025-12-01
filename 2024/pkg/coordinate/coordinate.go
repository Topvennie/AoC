package coordinate

type Coord struct {
	X int
	Y int
}

type Dir int

const (
	Up Dir = iota + 1
	Right
	Down
	Left
)

var Dirs = []Dir{Up, Right, Down, Left}

var DirToCoord = map[Dir]Coord{
	Up:    {X: 0, Y: -1},
	Right: {X: 1, Y: 0},
	Left:  {X: -1, Y: 0},
	Down:  {X: 0, Y: 1},
}

func New(x, y int) *Coord {
	return &Coord{X: x, Y: y}
}

func (c *Coord) Equal(c2 Coord) bool {
	return c.X == c2.X && c.Y == c2.Y
}

func (c *Coord) Add(c2 Coord) *Coord {
	return &Coord{X: c.X + c2.X, Y: c.Y + c2.Y}
}

func (c *Coord) AddDir(dir Dir) *Coord {
	return c.Add(DirToCoord[dir])
}

func (c *Coord) Sub(c2 Coord) *Coord {
	return &Coord{X: c.X - c2.X, Y: c.Y - c2.Y}
}
