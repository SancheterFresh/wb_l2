package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/


type Strategy interface {
	Operate(a, b int) int
}

type Add struct{}

func (add *Add) Operate(a, b int) int {
	return a + b
}

type Mult struct{}

func (mul *Mult) Operate(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

/*
func main() {
	context := &Context{}

	context.SetStrategy(&Add{})

	println(context.strategy.Operate(5, 4))

	context.SetStrategy(&Mult{})

	println(context.strategy.Operate(5, 4))

}
*/
