package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Handler interface {
	SendRequest(n int) (res string)
}

type ZeroHandler struct {
	next Handler
}

func (h *ZeroHandler) SendRequest(message int) (result string) {
	if message == 0 {
		result = "Its Zero"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type OddHandler struct {
	next Handler
}

func (h *OddHandler) SendRequest(message int) (result string) {
	if message%2 == 0 {
		result = "Its Odd"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type EvenHandler struct {
	next Handler
}

func (h *EvenHandler) SendRequest(message int) (result string) {
	if message%2 != 0 {
		result = "Its Even"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

/*
func main() {
	handlers := &ZeroHandler{
		next: &EvenHandler{
			next: &OddHandler{},
		},
	}

	print(handlers.SendRequest(7))
}
*/
