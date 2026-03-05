package data

type Hardware struct {
	CPU    CPU    `json:"cpu"`
	MEMORY Memory `json:"memory"`
}

func (h Hardware) GetCPU() CPU {
	return h.CPU
}

func (h Hardware) GetMemory() Memory {
	return h.MEMORY
}

func (h *Hardware) SetCPU(cpu CPU) {
	h.CPU = cpu
}

func (h *Hardware) SetMemory(memory Memory) {
	h.MEMORY = memory
}
