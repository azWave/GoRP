package entities

type Character struct {
	Name  string
	Class string
	Stats Stats
}

type Stats struct {
	Health       int // PV
	Mana         int // PM
	Strength     int // Force
	Intelligence int // Intelligence
	Defense      int // Défense
	MagicResist  int // Résistance Magique
	Agility      int // Agilité
	Luck         int // Chance
	Endurance    int // Endurance
	Spirit       int // Esprit
}
