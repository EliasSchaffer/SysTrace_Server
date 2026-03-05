package data

type Device struct {
	ID       string   `json:"id"`
	OS       string   `json:"os"`
	Hostname string   `json:"hostname"`
	Hardware Hardware `json:"hardware"`
	GPS      GPS      `json:"gps"`
}

// GetID returns the ID of the device.
func (d Device) GetID() string {
	return d.ID
}

func (d Device) GetOS() string {
	return d.OS
}

// GetHostname returns the hostname of the device.
func (d Device) GetHostname() string {
	return d.Hostname
}

// GetHardware returns the Hardware associated with the Device.
func (d Device) GetHardware() Hardware {
	return d.Hardware
}

// GetGPS returns a pointer to the GPS data of the device.
func (d *Device) GetGPS() *GPS {
	return &d.GPS
}

// SetID sets the ID of the Device.
func (d *Device) SetID(id string) {
	d.ID = id
}

func (d *Device) SetOS(os string) {
	d.OS = os
}

// SetHostname sets the hostname of the device.
func (d *Device) SetHostname(hostname string) {
	d.Hostname = hostname
}

func (d *Device) SetHardware(hardware Hardware) {
	d.Hardware = hardware
}

func (d *Device) SetGPS(gps GPS) {
	d.GPS = gps
}
