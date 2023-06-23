package addressable

import (
	"bytes"
	"os"
	"text/template"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
)

type Address struct {
	BuildingNumber string `json:"building_number"`
	StreetName     string `json:"street_name"`
	City           string `json:"city"`
	State          string `json:"state"`
	ZipCode        string `json:"zip_code"`
	Country        string `json:"country"`

	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

var AddressTemplate = "{{.BuildingNumber}} {{.StreetName}}, {{.City}}, {{.State}}, {{.ZipCode}}, {{.Country}}"
var Geocoder geo.Geocoder

func init() {
	if Geocoder == nil {
		Geocoder = google.Geocoder(os.Getenv("GOOGLE_API_KEY"))
	}
}

func (a Address) FullAddress() (address string, err error) {
	t, err := template.New("address").Parse(AddressTemplate)
	if err != nil {
		return "", err
	}

	b := new(bytes.Buffer)
	err = t.Execute(b, a)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (a Address) GenerateCoordinates() (Coordinates, error) {
	lat, lng, err := a.LatLng()
	if err != nil {
		return Coordinates{}, err
	}

	return Coordinates{lat, lng}, nil
}

func (a Address) LatLng() (lat, lng float64, err error) {
	addr, err := a.FullAddress()
	if err != nil {
		return 0, 0, err
	}

	location, err := Geocoder.Geocode(addr)
	if err != nil {
		return 0, 0, err
	}

	return location.Lat, location.Lng, nil
}
