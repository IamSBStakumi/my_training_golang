package animal

import "fmt"

func (a *AStruct) Eat(food string) error {
	_, err := fmt.Printf("%s eats %s\n", a.Name, food)

	return err
}
