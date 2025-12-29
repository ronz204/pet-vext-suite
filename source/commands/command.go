package commands

type Command struct {
	Use     string
	Long    string
	Short   string
	Example string
}

func (c *Command) Info() {
	println("Use:", c.Use)
	println("Long:", c.Long)
	println("Short:", c.Short)
	println("Example:", c.Example)
}
