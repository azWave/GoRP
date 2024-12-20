package interfaces

type MapRepository interface {
	SaveMap(mapName string, mapData string) error
	LoadMap(mapName string) (string, error)
}
