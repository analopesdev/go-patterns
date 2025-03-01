package main

type FordFactory struct{}

func (f FordFactory) CreateCar() Car {
	return Ford{}
}
