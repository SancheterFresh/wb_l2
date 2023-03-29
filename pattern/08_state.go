package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	Act() string
}

type Actor struct {
	state State
}

func (a *Actor) Act() string {
	return a.state.Act()
}

func (a *Actor) SetState(s State) {
	a.state = s
}

func NewActor() *Actor {
	return &Actor{state: &Write{}}
}

type Write struct{}

func (s *Write) Act() string {
	return "Writing..."
}

type Read struct{}

func (s *Read) Act() string {
	return "Reading..."
}

/*
func main() {
	actor := NewActor()
	println(actor.Act())
	actor.SetState(&Read{})
	println(actor.Act())
	actor.SetState(&Write{})
	println(actor.Act())
}
*/
