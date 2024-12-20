package entities

type Character struct {
	Name     string
	Class    string
	Stats    Stats
	Position Position
}

type Stats struct {
	Health       int16 // PV
	Mana         int16 // PM
	Strength     int16 // Force
	Intelligence int16 // Intelligence
	Defense      int16 // Défense
	MagicResist  int16 // Résistance Magique
	Agility      int16 // Agilité
	Luck         int16 // Chance
	Endurance    int16 // Endurance
	Spirit       int16 // Esprit
}

type Position struct {
	X int
	Y int
}

var ClassStats = map[string]Stats{
	"Guerrier": {
		Health:       150,
		Mana:         50,
		Strength:     15,
		Intelligence: 5,
		Defense:      12,
		MagicResist:  6,
		Agility:      8,
		Luck:         5,
		Endurance:    10,
		Spirit:       4,
	},
	"Mage": {
		Health:       90,
		Mana:         150,
		Strength:     4,
		Intelligence: 15,
		Defense:      5,
		MagicResist:  12,
		Agility:      7,
		Luck:         6,
		Endurance:    5,
		Spirit:       10,
	},
	"Voleur": {
		Health:       110,
		Mana:         70,
		Strength:     10,
		Intelligence: 7,
		Defense:      8,
		MagicResist:  7,
		Agility:      15,
		Luck:         12,
		Endurance:    7,
		Spirit:       6,
	},
}
