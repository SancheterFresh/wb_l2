package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Product interface {
	GetName() string
}

type Factory interface {
	CreateProduct() Product
}

type playstation struct{}

func (p *playstation) GetName() string {
	return "ps5"
}

type xbox struct{}

func (p *xbox) GetName() string {
	return "xboxS"
}

type SonyFactory struct{}

func (f *SonyFactory) CreateProduct() Product {
	return &playstation{}
}

type MicrosoftFactory struct{}

func (f *MicrosoftFactory) CreateProduct() Product {
	return &xbox{}
}

/*
func main() {
	sonyFactory := &SonyFactory{}
	microsoftFactory := &MicrosoftFactory{}

	ps := sonyFactory.CreateProduct()
	println(ps.GetName())

	xbox := microsoftFactory.CreateProduct()
	println(xbox.GetName())

}
*/
