package dijkstra

type Vertex uint8

const (
	İstanbul Vertex = iota + 1
	Samsun
	Ankara
	Bursa
	İzmir
	Kütahya
	Konya
	Antalya
	Adana
	Diyarbakır
)

func (v Vertex) String() string {
	return [...]string{"", "İstanbul", "Samsun", "Ankara", "Bursa", "İzmir", "Kütahya", "Konya", "Antalya", "Adana", "Diyarbakır"}[v]
}
