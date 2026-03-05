package data

type Hardware struct {
	CPU    CPU    `json:"cpu"`
	MEMORY Memory `json:"memory"`
}

// GetCPU returns the CPU of the hardware.
func (h Hardware) GetCPU() CPU {
	return h.CPU
}

// GetMemory returns the MEMORY field of the Hardware.
func (h Hardware) GetMemory() Memory {
	return h.MEMORY
}

// SetCPU sets the CPU for the Hardware.
func (h *Hardware) SetCPU(cpu CPU) {
	h.CPU = cpu
}

// SetMemory sets the MEMORY field of the Hardware struct.
func (h *Hardware) SetMemory(memory Memory) {
	h.MEMORY = memory
}
