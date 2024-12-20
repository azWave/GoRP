package entities

const (
	Land  CellType = "Land"
	Rock  CellType = "Rock"
	Water CellType = "Water"
)

type CellProperties struct {
	Wallkable bool
	Display   string
}

type CellType string

type Cell struct {
	X    int
	Y    int
	Type CellType
}
