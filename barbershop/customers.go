package barbershop

type Customer struct {
	Name	string
	Style	string
}

func (c *Customer) SetName(name string) {
	c.Name = name
}

func (c *Customer) SetStyle(style string) {
	c.Style = style
}
