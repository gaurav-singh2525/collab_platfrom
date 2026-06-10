package websocket

type Room struct {

	Clients map[*Client]bool

	Code string

	Language string
}