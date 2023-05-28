package models

type Super struct {
	Name        string
	Description string
	Damage      string
	Scope       string
}
type Brawler struct {
	Id       int
	Name     string
	Image    string
	Type     string
	Category string
	Stellar  [2]string
	Gadget   [2]string
	Super
}
