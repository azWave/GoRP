package entities

var ClassCellProperties = map[CellType]CellProperties{
	Land: {
		Wallkable: true,
		Display:   ".",
	},
	Rock: {
		Wallkable: false,
		Display:   "¤",
	},
	Water: {
		Wallkable: false,
		Display:   "~",
	},
}
