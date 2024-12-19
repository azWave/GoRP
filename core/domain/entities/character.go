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
