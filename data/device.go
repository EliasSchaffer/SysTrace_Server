package data

type Device struct {
	ID       string   `json:"id"`
	OS       string   `json:"os"`
	Hostname string   `json:"hostname"`
	Hardware Hardware `json:"hardware"`
	GPS      GPS      `json:"gps"`
}

// Getter-Methoden
func (d Device) GetID() string {
	return d.ID
}

func (d Device) GetOS() string {
	return d.OS
}

func (d Device) GetHostname() string {
	return d.Hostname
}

func (d Device) GetHardware() Hardware {
	return d.Hardware
}

func (d *Device) GetGPS() *GPS {
	return &d.GPS
}

func (d *Device) SetID(id string) {
	d.ID = id
}

func (d *Device) SetOS(os string) {
	d.OS = os
}

func (d *Device) SetHostname(hostname string) {
	d.Hostname = hostname
}

func (d *Device) SetHardware(hardware Hardware) {
	d.Hardware = hardware
}

func (d *Device) SetGPS(gps GPS) {
	d.GPS = gps
}
