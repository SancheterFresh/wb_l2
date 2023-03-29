package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Visitor interface {
	VisitGPU(*GPU)
	VisitCPU(*CPU)
	VisitComputer(*Computer)
}

type ComputerPart interface {
	Accept(Visitor)
}

type Computer struct {
	Parts []ComputerPart
}

func (c *Computer) Accept(v Visitor) {
	for _, part := range c.Parts {
		part.Accept(v)
	}
	v.VisitComputer(c)
}

type GPU struct {
	temperature int
}

func (g *GPU) Accept(v Visitor) {
	v.VisitGPU(g)
}

type CPU struct {
	temperature int
}

func (c *CPU) Accept(v Visitor) {
	v.VisitCPU(c)
}

type temperatureChecker struct {
	cTemp int
	gTemp int
}

func (t *temperatureChecker) VisitCPU(c *CPU) {
	print(c.temperature)
	t.cTemp = c.temperature
}

func (t *temperatureChecker) VisitGPU(g *GPU) {
	print(g.temperature)
	t.gTemp = g.temperature
}

func (t *temperatureChecker) VisitComputer(c *Computer) {
	print((t.cTemp + t.gTemp) / 2)
}

/*
func main() {
	computer := &Computer{
		Parts: []ComputerPart{
			&GPU{temperature: 36},
			&CPU{temperature: 54},
		},
	}

	tChecker := &temperatureChecker{}
	computer.Accept(tChecker)
}
*/