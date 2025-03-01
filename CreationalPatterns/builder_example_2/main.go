package main

import "fmt"

func main() {
	car := NewCarBuilder().
		SetBrand("Toyota").
		SetModel("Corolla").
		SetYear(2024).
		SetEngine("2.0L Turbo").
		SetColor("Azul").
		Build()

	fmt.Printf("Carro Construído: %+v\n", car)
}
