package data

type Memory struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"usedPercent"`
	Model       string  `json:"model"`
	Speed       uint64  `json:"speed"`
}

func (m Memory) GetTotal() uint64 {
	return m.Total
}

func (m Memory) GetUsed() uint64 {
	return m.Used
}

func (m Memory) GetAvailable() uint64 {
	return m.Available
}

func (m Memory) GetUsedPercent() float64 {
	return m.UsedPercent
}

func (m Memory) GetModel() string {
	return m.Model
}

func (m Memory) GetSpeed() uint64 {
	return m.Speed
}

func (m *Memory) SetTotal(total uint64) {
	m.Total = total
}

func (m *Memory) SetUsed(used uint64) {
	m.Used = used
}

func (m *Memory) SetAvailable(available uint64) {
	m.Available = available
}

func (m *Memory) SetUsedPercent(usedPercent float64) {
	m.UsedPercent = usedPercent
}

func (m *Memory) SetModel(model string) {
	m.Model = model
}

func (m *Memory) SetSpeed(speed uint64) {
	m.Speed = speed
}
