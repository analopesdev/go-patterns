package main

type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (c *CarBuilder) SetBrand(brand string) *CarBuilder {
	c.car.Brand = brand
	return c
}

func (c *CarBuilder) SetModel(model string) *CarBuilder {
	c.car.Model = model
	return c
}

func (c *CarBuilder) SetYear(year int) *CarBuilder {
	c.car.Year = year
	return c
}

func (c *CarBuilder) SetEngine(engine string) *CarBuilder {
	c.car.Engine = engine
	return c
}

func (c *CarBuilder) SetColor(color string) *CarBuilder {
	c.car.Color = color
	return c
}

func (c *CarBuilder) Build() Car {
	return c.car
}
