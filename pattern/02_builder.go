package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Builder interface {
	SetName(str string)
	SetLogin(str string)
	SetLevel(str string)
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Construct() {
	d.builder.SetName("Nicola")
	d.builder.SetLogin("nik.ola")
	d.builder.SetLevel("Creator")

}

type Page struct {
	Descriprion string
}

func (p *Page) Show() string {
	return p.Descriprion
}

type TwitterBuilder struct {
	profile *Page
}

func NewTweeterBuilder(p *Page) *TwitterBuilder {
	return &TwitterBuilder{
		profile: p,
	}
}

func (b *TwitterBuilder) SetName(str string) {
	b.profile.Descriprion += "TweetName: " + str + "\n"
}

func (b *TwitterBuilder) SetLogin(str string) {
	b.profile.Descriprion += "TweetLogin: " + str + "\n"
}

func (b *TwitterBuilder) SetLevel(str string) {
	b.profile.Descriprion += "TweetLevel " + str + "\n"
}

type VkBuilder struct {
	profile *Page
}

func NewVkBuilder(p *Page) *VkBuilder {
	return &VkBuilder{
		profile: p,
	}
}

func (b *VkBuilder) SetName(str string) {
	b.profile.Descriprion += "vkName: " + str + "\n"
}

func (b *VkBuilder) SetLogin(str string) {
	b.profile.Descriprion += "vkLogin: " + str + "\n"
}

func (b *VkBuilder) SetLevel(str string) {
	b.profile.Descriprion += "vkLevel " + str + "\n"
}

/*
func main() {
	tPage := new(Page)
	tBuilder := NewTweeterBuilder(tPage)
	tDir := NewDirector(tBuilder)

	tDir.Construct()
	print(tPage.Show())

	vkPage := new(Page)
	vkBuilder := NewVkBuilder(vkPage)
	vkDir := NewDirector(vkBuilder)

	vkDir.Construct()
	print(vkPage.Show())

} */