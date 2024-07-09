package main

import (
	"fmt"

	"vendingmachine/vm"
)

type vendingMachine interface {
	GetDrink(brand string) string
}

type Application struct {
	vm vendingMachine
}

func NewApplication(vm vendingMachine) *Application {
	return &Application{vm: vm}
}

func (a *Application) Run() {
	fmt.Println(a.vm.GetDrink("coke"))
}

func main() {
	vm := vm.NewVendingMachine()
	app := NewApplication(vm)
	app.Run()
}
