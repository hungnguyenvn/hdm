package hdm

import (
	"context"
	"encoding/json"
	"fmt"
	"hdm/pb"
	"io/ioutil"
	"net/http"
)

var dataSource = []string{
	"https://api.myjson.com/bins/gdmqa",
	"https://api.myjson.com/bins/1fva3m",
	"https://api.myjson.com/bins/j6kzm"}

func NewGRPCHandler() pb.HDMServer {
	handler := grpcHandler{}
	return handler
}

type grpcHandler struct {
}

func (grpcHandler) Hotels(ctx context.Context, request *pb.HotelsRequest) (*pb.HotelsResponse, error) {
	getData()
	return &pb.HotelsResponse{Msg: request.Msg + "____ response"}, nil
}

func getData() {
	response, err := http.Get("https://api.myjson.com/bins/gdmqa")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var hotels []Hotel_varian_1
	err = json.Unmarshal(responseData, &hotels)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i, hotel := range hotels {
		fmt.Printf("%d>>>>>> %s- %s \n", i, hotel.ID, hotel.Name)
		fmt.Println(hotel.Latitude)
	}
}
