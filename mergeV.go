package hdm

func (h *Hotel) mergeV1(v1 HotelV1) {
	if v1.Latitude != nil {
		h.Location.Lat = float32(*v1.Latitude)
	}
	if v1.Longtitude != nil {
		h.Location.Long = float32(*v1.Latitude)
	}
	h.Location.Address = v1.Address
	h.Location.Country = v1.Country
	h.Location.City = v1.City

	h.BookingCondition = append(h.BookingCondition, v1.Description)
}

func (h *Hotel) mergeV2(v HotelV2) {
	h.BookingCondition = append(h.BookingCondition, v.Description)
}

func (h *Hotel) mergeV3(v HotelV3) {
	h.BookingCondition = append(h.BookingCondition, v.Infor)
}

func (h *Hotel) merge(varian interface{}) {
	switch varian.(type) {
	case HotelV1:
		h.mergeV1(varian.(HotelV1))
	case HotelV2:
		h.mergeV2(varian.(HotelV2))
	case HotelV3:
		h.mergeV3(varian.(HotelV3))
	}
}

func makeHotelFromVarian(h interface{}) Hotel {
	switch h.(type) {
	case HotelV1:
		return makeHotelFromVarian_1(h.(HotelV1))
	case HotelV2:
		return makeHotelFromVarian_2(h.(HotelV2))
	case HotelV3:
		return makeHotelFromVarian_3(h.(HotelV3))
	}
	return Hotel{}
}
func makeHotelFromVarian_1(h HotelV1) Hotel {
	return Hotel{
		ID:            h.ID,
		DestinationID: h.DestinationID,
		Name:          h.Name,
		Address:       h.Address,
		City:          h.City,
		Facilities:    h.Facilities,
		Location:      Location{},
	}
}

func makeHotelFromVarian_2(h HotelV2) Hotel {
	return Hotel{
		ID:            h.ID,
		DestinationID: h.DestinationID,
		Name:          h.Name,
		HotelImages:   h.HotelImages,
		Location:      Location{},
	}
}

func makeHotelFromVarian_3(h HotelV3) Hotel {
	return Hotel{
		ID:            h.ID,
		DestinationID: h.DestinationID,
		Name:          h.Name,
		Location:      Location{},
	}
}
