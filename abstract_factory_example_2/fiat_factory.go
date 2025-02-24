package main

type FiatFactory struct{}

func (f FiatFactory) CreateCar() Car {
	return Fiat{}
}
