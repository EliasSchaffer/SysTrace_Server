package data

type CPU struct {
	Usage   float64 `json:"usage"`
	Cores   int     `json:"cores"`
	Threads int     `json:"threads"`
	Temp    float64 `json:"temp"`
	Model   string  `json:"model"`
}

// Getter-Methoden
func (c CPU) GetUsage() float64 {
	return c.Usage
}

func (c CPU) GetCores() int {
	return c.Cores
}

func (c CPU) GetThreads() int {
	return c.Threads
}

func (c CPU) GetTemp() float64 {
	return c.Temp
}

func (c CPU) GetModel() string {
	return c.Model
}

// Setter-Methoden
func (c *CPU) SetUsage(usage float64) {
	c.Usage = usage
}

func (c *CPU) SetCores(cores int) {
	c.Cores = cores
}

func (c *CPU) SetThreads(threads int) {
	c.Threads = threads
}

func (c *CPU) SetTemp(temp float64) {
	c.Temp = temp
}

func (c *CPU) SetModel(model string) {
	c.Model = model
}
