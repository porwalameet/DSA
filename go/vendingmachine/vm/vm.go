package vm

import "fmt"

type vendingMachine struct{}

func NewVendingMachine() *vendingMachine {
	return &vendingMachine{}
}

func (v *vendingMachine) GetDrink(brand string) string {
	return fmt.Sprintf("Enjoy your %s", brand)
}
