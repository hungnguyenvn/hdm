package hdm

import (
	"context"
	"hdm/pb"
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
		return &pb.HotelsResponse{
			HotelId:           request.HotelId,
			DestinationId:     request.DestinationId,
			Name:              h.Name,
			Location:          makeLocation(&h),
			BookingConditions: h.BookingCondition,
		}, nil

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
		return &pb.HotelsResponse{
			HotelId:           request.HotelId,
			DestinationId:     request.DestinationId,
			Name:              h.Name,
			Location:          makeLocation(&h),
			BookingConditions: h.BookingCondition,
		}, nil
	}

	return &pb.HotelsResponse{}, nil

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
