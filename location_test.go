package location_test

import (
	"testing"

	"github.com/hojin-kr/location"
)

func TestLocation(t *testing.T) {
	location := location.Location{}
	location.UUID = "test"
	location.Location = "test"
	location.SetLocation()
	if location.GetLocation() != "test" {
		t.Errorf("Expected %s, got %s", "test", location.GetLocation())
	}
}
