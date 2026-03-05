package data

type Memory struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"usedPercent"`
	Model       string  `json:"model"`
	Speed       uint64  `json:"speed"`
}

// GetTotal returns the total memory value.
func (m Memory) GetTotal() uint64 {
	return m.Total
}

// GetUsed returns the amount of memory used.
func (m Memory) GetUsed() uint64 {
	return m.Used
}

func (m Memory) GetAvailable() uint64 {
	return m.Available
}

// GetUsedPercent returns the percentage of memory used.
func (m Memory) GetUsedPercent() float64 {
	return m.UsedPercent
}

// GetModel returns the model of the Memory.
func (m Memory) GetModel() string {
	return m.Model
}

// GetSpeed returns the speed of the memory.
func (m Memory) GetSpeed() uint64 {
	return m.Speed
}

func (m *Memory) SetTotal(total uint64) {
	m.Total = total
}

// SetUsed sets the used memory to the specified value.
func (m *Memory) SetUsed(used uint64) {
	m.Used = used
}

// SetAvailable sets the available memory to the specified value.
func (m *Memory) SetAvailable(available uint64) {
	m.Available = available
}

// SetUsedPercent sets the used percentage of memory.
func (m *Memory) SetUsedPercent(usedPercent float64) {
	m.UsedPercent = usedPercent
}

// SetModel sets the model of the Memory instance.
func (m *Memory) SetModel(model string) {
	m.Model = model
}

// SetSpeed sets the speed of the memory.
func (m *Memory) SetSpeed(speed uint64) {
	m.Speed = speed
}
