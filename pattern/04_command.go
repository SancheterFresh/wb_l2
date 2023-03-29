package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Execute()
}

type SayHiCommand struct {
}

func (c *SayHiCommand) Execute() {
	print("Hi!")
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(c Command) {
	i.command = c
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

/*
func main() {
	invoker := &Invoker{}
	sayHiCommand := &SayHiCommand{}

	invoker.SetCommand(sayHiCommand)
	invoker.ExecuteCommand()
}
*/