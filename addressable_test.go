package addressable

import (
	"testing"

	"github.com/codingsince1985/geo-golang/google"
)

type TestStruct struct {
	Name string
	Address
}

var addr = TestStruct{
	Address: Address{
		BuildingNumber: "123",
		StreetName:     "Main St.",
		City:           "Springfield",
		State:          "IL",
		ZipCode:        "12345",
		Country:        "USA",
	},
}

func InitTest() {
	Geocoder = google.Geocoder("APIKEY")
}

func TestAddress(t *testing.T) {
	expected := "123 Main St., Springfield, IL, 12345, USA"
	actual, _ := addr.FullAddress()

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGeocode(t *testing.T) {
	InitTest()

	expectedLat := 39.7498563
	expectedLng := -89.536712

	actualLat, actualLng, err := addr.LatLng()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if actualLat != expectedLat {
		t.Errorf("Expected %f, got %f", expectedLat, actualLat)
	}

	if actualLng != expectedLng {
		t.Errorf("Expected %f, got %f", expectedLng, actualLng)
	}
}

func TestGenerateCoordinates(t *testing.T) {
	InitTest()

	expectedLat := 39.7498563
	expectedLng := -89.536712

	actual, err := addr.GenerateCoordinates()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if actual.Lat != expectedLat {
		t.Errorf("Expected %f, got %f", expectedLat, actual.Lat)
	}

	if actual.Lng != expectedLng {
		t.Errorf("Expected %f, got %f", expectedLng, actual.Lng)
	}
}
