package hdm

type QueryHotel interface {
	getHotelID() string
	getDestinationID() uint
}

func (h Hotel) getHotelID() string {
	return h.ID
}
func (h Hotel) getDestinationID() uint {
	return h.DestinationID
}

func (h HotelV1) getHotelID() string {
	return h.ID
}
func (h HotelV1) getDestinationID() uint {
	return h.DestinationID
}

func (h HotelV2) getHotelID() string {
	return h.ID
}

func (h HotelV2) getDestinationID() uint {
	return h.DestinationID
}

func (h HotelV3) getHotelID() string {
	return h.ID
}
func (h HotelV3) getDestinationID() uint {
	return h.DestinationID
}
