package models

type Jugador struct {
	Id       int `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Numero   int `json:"numero"`
	Posicion string `json:"posicion"`
}
type Database struct {
	Jugadores []Jugador `json:"jugadores"`
}
type Message struct {
	Text string `json:"message"`
}
