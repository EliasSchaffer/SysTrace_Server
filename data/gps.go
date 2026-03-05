package data

type GPS struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Accuracy  float64 `json:"accuracy"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Region    string  `json:"region"`
}

func (g *GPS) GetLatitude() float64 {
	return g.Latitude
}

func (g *GPS) SetLatitude(latitude float64) {
	g.Latitude = latitude
}

func (g *GPS) GetLongitude() float64 {
	return g.Longitude
}

func (g *GPS) SetLongitude(longitude float64) {
	g.Longitude = longitude
}

func (g *GPS) GetAltitude() float64 {
	return g.Altitude
}

func (g *GPS) SetAltitude(altitude float64) {
	g.Altitude = altitude
}

func (g *GPS) GetAccuracy() float64 {
	return g.Accuracy
}

func (g *GPS) SetAccuracy(accuracy float64) {
	g.Accuracy = accuracy
}

func (g *GPS) GetCity() string {
	return g.City
}

func (g *GPS) SetCity(city string) {
	g.City = city
}

func (g *GPS) GetCountry() string {
	return g.Country
}

func (g *GPS) SetCountry(country string) {
	g.Country = country
}

func (g *GPS) GetRegion() string {
	return g.Region
}

func (g *GPS) SetRegion(region string) {
	g.Region = region
}
