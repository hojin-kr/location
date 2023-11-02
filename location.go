package location

import (
	"github.com/hojin-kr/datastore"
)

type Location struct {
	UUID     string `json:"uuid"`
	Location string `json:"location"`
}

var ds datastore.GcpDatastore

func (location *Location) GetLocation() string {
	ds.Init()
	location.Location = ds.Get(location.UUID)
	return location.Location
}

func (location *Location) SetLocation() {
	ds.Init()
	ds.Put(location.UUID, location.Location)
}
