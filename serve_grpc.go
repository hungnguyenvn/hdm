package hdm

import (
	"context"
	"hdm/pb"
	"strconv"
)

func NewGRPCHandler() pb.HDMServer {
	handler := grpcHandler{}
	return handler
}

type grpcHandler struct {
}

func (grpcHandler) Hotels(ctx context.Context, request *pb.HotelsRequest) (*pb.HotelsResponse, error) {
	mt.RLock()
	defer mt.RUnlock()
	if len(request.HotelId) > 0 {
		h, ok := hotelsDataSource[request.HotelId]
		if !ok {
			return &pb.HotelsResponse{}, nil
		}
		return makeHotelResp(&h)
	}
	if len(request.DestinationId) > 0 {
		hID, ok := hotelsIDMaping[request.DestinationId]
		if !ok {
			return &pb.HotelsResponse{}, nil
		}
		h, ok := hotelsDataSource[hID]
		if !ok {
			return &pb.HotelsResponse{}, nil
		}
		return makeHotelResp(&h)
	}

	return &pb.HotelsResponse{}, nil

}

func makeHotelResp(h *Hotel) (*pb.HotelsResponse, error) {
	desID := strconv.Itoa(int(h.DestinationID))

	return &pb.HotelsResponse{
		HotelId:           h.ID,
		DestinationId:     desID,
		Name:              h.Name,
		Location:          makeLocation(h),
		Amenities:         makeAmenities(h),
		Images:            makeImage(h),
		BookingConditions: h.BookingCondition,
		Description:       h.Description,
	}, nil
}

func makeImage(h *Hotel) *pb.ImageResponse {
	ames := make([]*pb.ImageResponse_Image, 0)
	rooms := make([]*pb.ImageResponse_Image, 0)
	sites := make([]*pb.ImageResponse_Image, 0)

	for _, v := range h.HotelImages.Rooms {
		image := &pb.ImageResponse_Image{
			Link:        v.Link,
			Description: v.Description,
		}
		rooms = append(rooms, image)
	}

	for _, v := range h.HotelImages.Site {
		image := &pb.ImageResponse_Image{
			Link:        v.Link,
			Description: v.Description,
		}
		rooms = append(rooms, image)
	}
	for _, v := range h.HotelImages.Amenities {
		image := &pb.ImageResponse_Image{
			Link:        v.Link,
			Description: v.Description,
		}
		ames = append(ames, image)
	}
	return &pb.ImageResponse{
		Amenities: ames,
		Rooms:     rooms,
		Site:      sites,
	}
}

func makeLocation(h *Hotel) *pb.HotelsResponse_Location {
	return &pb.HotelsResponse_Location{
		Lat:     h.Location.Lat,
		Lng:     h.Location.Long,
		City:    h.Location.City,
		Country: h.Location.Country,
		Address: h.Location.Address,
	}
}

func makeAmenities(h *Hotel) *pb.HotelsResponse_Amenities {
	return &pb.HotelsResponse_Amenities{
		General: h.Amenities.General,
		Room:    h.Amenities.Room,
	}
}
