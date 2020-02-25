package hdm

import "testing"

func TestMergeV1Location(t *testing.T) {
	h := mockHotel()
	lat := FloatString(1.0)
	lng := FloatString(2.0)
	h1 := HotelV1{
		ID:            "1",
		DestinationID: 1,
		Latitude:      &lat,
		Longtitude:    &lng,
		Address:       "A1",
		City:          "C1",
		Country:       "C1",
		PostalCode:    "P1",
		Description:   "D1",
		Facilities:    []string{"1", "2"},
	}
	h.mergeV1(h1)

	if h.Location.Address != "A1" {
		t.Errorf("Address should be A1, got: %s", h.Location.Address)
	}
	if h.Location.City != "C1" {
		t.Errorf("City should be city, got: %s", h.Location.City)
	}
	if h.Location.Country != "C1" {
		t.Errorf("Country should be city, got: %s", h.Location.Country)
	}
	if !almostEqual(float64(h.Location.Lat), 1.0) {
		t.Errorf("Location lat should be 1.0, got: %f", h.Location.Lat)
	}
	if !almostEqual(float64(h.Location.Long), 2.0) {
		t.Errorf("Location long should be 2.0, got: %f", h.Location.Long)
	}
}

func mockHotel() *Hotel {
	return &Hotel{
		ID:               "1",
		DestinationID:    1,
		Name:             "name",
		Location:         Location{},
		City:             "city",
		HotelImages:      HotelImages{},
		Facilities:       make([]string, 0),
		BookingCondition: make([]string, 0),
		Amenities:        Amenities{},
	}
}
