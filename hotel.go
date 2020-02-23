package hdm

type Hotel_varian_1 struct {
	ID             string       `json:"Id"`
	DestinationsID uint         `json:"DestinationId"`
	Name           string       `json:"Name"`
	Latitude       *FloatString `json:"Latitude,omitempty"`
	Longtitude     *FloatString `json:"Longitude,omitempty"`
	Address        string       `json:"Address"`
	City           string       `json:"City"`
	Country        string       `json:"Country"`
	PostalCode     string       `json:"PostalCode"`
	Description    string       `json:"Description"`
	Facilities     []string     `json:"Facilities"`
}

type Location struct {
	Address string `json:"address"`
	Country string `json:"country"`
}
type Image struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

type HotelImages struct {
	Rooms []Image `json:"rooms"`
	Site  []Image `json:"site"`
}

type Hotel_varian_2 struct {
	ID            string      `json:"hotel_id"`
	DestinationID int64       `json:"destination_id"`
	Name          string      `json:"hotel_name"`
	Location      Location    `json:"location"`
	Description   string      `json:"details"`
	HotelImages   HotelImages `json:"images"`
	Conditions    []string    `json:"booking_conditions"`
}

type Hotel_varian_3 struct {
	ID            string       `json:"id"`
	DestinationID int64        `json:"destination"`
	Name          string       `json:"name"`
	Latitude      *FloatString `json:"lat"`
	Longtitude    *FloatString `json:"lng"`
	Address       *string      `json:"address"`
}
