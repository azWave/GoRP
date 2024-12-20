package entities

type TileType string

const (
	Land  TileType = "Land"
	Rock  TileType = "Rock"
	Water TileType = "Water"
)

type Tile struct {
	Type         TileType
	HasCollision bool
	DisplayChar  string
}

func (t Tile) IsWalkable() bool {
	return !t.HasCollision
}

var TileTypes = map[TileType]Tile{
	"Land":  {Type: Land, HasCollision: false, DisplayChar: ". "},
	"Rock":  {Type: Rock, HasCollision: true, DisplayChar: "¤ "},
	"Water": {Type: Water, HasCollision: true, DisplayChar: "~ "},
}
