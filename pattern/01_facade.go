package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"strings"
)

func NewProfile() *Profile {
	return &Profile{
		name:  &Name{},
		login: &Login{},
		level: &Level{},
	}
}

type Profile struct {
	name  *Name
	login *Login
	level *Level
}

func (p *Profile) ShowInfo() string {
	res := []string{
		p.name.Say(),
		p.login.Enter(),
		p.level.Show(),
	}
	return strings.Join(res, " ")

}

type Name struct {
	name string
}

func (n *Name) Say() string {
	return n.name
}

type Login struct {
	login string
}

func (l *Login) Enter() string {
	return l.login
}

type Level struct {
	level string
}

func (l *Level) Show() string {
	return l.level
}

/*
func main() {
	profile := NewProfile()

	profile.name.name = "Nick"
	profile.login.login = "usernikolas"
	profile.level.level = "admin"

	fmt.Println(profile.ShowInfo())
}
*/
