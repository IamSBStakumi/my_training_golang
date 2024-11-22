package animal

import "fmt"

func (a *AStruct) Walk() error {
	if !a.HasLegs {
		return fmt.Errorf("%s doesn't have legs", a.Name)
	}

	_, err := fmt.Printf("%s is walking\n", a.Name)

	return err
}
