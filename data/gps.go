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

// GetLatitude returns the latitude of the GPS.
func (g *GPS) GetLatitude() float64 {
	return g.Latitude
}

// SetLatitude sets the latitude of the GPS.
func (g *GPS) SetLatitude(latitude float64) {
	g.Latitude = latitude
}

func (g *GPS) GetLongitude() float64 {
	return g.Longitude
}

// SetLongitude sets the longitude of the GPS.
func (g *GPS) SetLongitude(longitude float64) {
	g.Longitude = longitude
}

// GetAltitude returns the altitude from the GPS.
func (g *GPS) GetAltitude() float64 {
	return g.Altitude
}

func (g *GPS) SetAltitude(altitude float64) {
	g.Altitude = altitude
}

// GetAccuracy returns the accuracy of the GPS.
func (g *GPS) GetAccuracy() float64 {
	return g.Accuracy
}

// SetAccuracy sets the accuracy of the GPS.
func (g *GPS) SetAccuracy(accuracy float64) {
	g.Accuracy = accuracy
}

func (g *GPS) GetCity() string {
	return g.City
}

// SetCity sets the city for the GPS instance.
func (g *GPS) SetCity(city string) {
	g.City = city
}

// GetCountry returns the country associated with the GPS instance.
func (g *GPS) GetCountry() string {
	return g.Country
}

// SetCountry sets the country for the GPS instance.
func (g *GPS) SetCountry(country string) {
	g.Country = country
}

// GetRegion returns the region of the GPS.
func (g *GPS) GetRegion() string {
	return g.Region
}

func (g *GPS) SetRegion(region string) {
	g.Region = region
}
