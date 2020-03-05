package hdm

type HotelV1 struct {
	ID            string       `json:"Id"`
	DestinationID uint         `json:"DestinationId"`
	Name          string       `json:"Name"`
	Latitude      *FloatString `json:"Latitude"`
	Longtitude    *FloatString `json:"Longitude"`
	Address       string       `json:"Address"`
	City          string       `json:"City"`
	Country       string       `json:"Country"`
	PostalCode    string       `json:"PostalCode"`
	Description   string       `json:"Description"`
	Facilities    []string     `json:"Facilities"`
}

type Amenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type Hotel struct {
	ID               string
	DestinationID    uint
	Name             string
	Location         Location
	Address          string
	City             string
	HotelImages      HotelImages
	Facilities       []string
	BookingCondition []string
	Description      string
	Amenities        Amenities
}

type Location struct {
	Lat     float32 `json:"-"`
	Long    float32 `json:"-"`
	Address string  `json:"address"`
	Country string  `json:"country"`
	City    string  `json:"-"`
}
type Image struct {
	Link        string `json:"link"`
	Description string `json:"caption"`
}

type ImageV1 struct {
	Link        string `json:"url"`
	Description string `json:"description"`
}

type HotelImages struct {
	Rooms     []Image   `json:"rooms"`
	Site      []Image   `json:"site"`
	Amenities []ImageV1 `json:"amenities"`
}

type HotelV2 struct {
	ID            string      `json:"hotel_id"`
	DestinationID uint        `json:"destination_id"`
	Name          string      `json:"hotel_name"`
	Location      Location    `json:"location"`
	Description   string      `json:"details"`
	HotelImages   HotelImages `json:"images"`
	Conditions    []string    `json:"booking_conditions"`
	Amenities     Amenities   `json:"amenities"`
}

type HotelV3 struct {
	ID            string       `json:"id"`
	DestinationID uint         `json:"destination"`
	Name          string       `json:"name"`
	Latitude      *FloatString `json:"lat"`
	Longtitude    *FloatString `json:"lng"`
	Address       *string      `json:"address"`
	Descriptions  string       `json:"info"`
	HotelImages   HotelImages  `json:"images"`
}
