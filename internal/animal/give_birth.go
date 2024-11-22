package animal

type AInterface interface {
	Eat(food string) error
	Walk() error
}

type AStruct struct {
	Name    string
	HasLegs bool
}

func GiveBirth(n string, l bool) AInterface {
	return &AStruct{Name: n, HasLegs: l}
}
