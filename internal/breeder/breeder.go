package breeder

import "train/internal/animal"

type BreederInterface interface {
	Breeder() error
}

type Pet struct {
	do animal.AInterface
}

func NewBreeder(do animal.AInterface) BreederInterface {
	return &Pet{do: do}
}

func (pet *Pet) Breeder() error {
	err := pet.do.Eat("food")
	if err != nil {
		return err
	}

	err = pet.do.Walk()
	if err != nil {
		return err
	}

	return nil
}
