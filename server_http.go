package hdm

import (
	"context"
	"encoding/json"
	"fmt"
	"hdm/pb"
	"net/http"

	"google.golang.org/grpc"
)

const httpAddr = "localhost:8080"

var client pb.HDMClient

func ConnectGRPC(ch chan error) *grpc.ClientConn {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		ch <- err
		return nil
	}

	client = pb.NewHDMClient(conn)
	return conn
}

func NewHTTPServer() *http.Server {
	http.HandleFunc("/hotel", func(rw http.ResponseWriter, rq *http.Request) {
		var hotelID string
		var desID string
		hotelIDs, ok := rq.URL.Query()["hotel_id"]
		if ok && len(hotelIDs) > 0 {
			hotelID = hotelIDs[0]
		}

		desIDs, ok := rq.URL.Query()["destination_id"]
		if ok && len(desIDs) > 0 {
			desID = desIDs[0]
		}

		resp, err := client.Hotels(context.Background(), &pb.HotelsRequest{
			HotelId:       hotelID,
			DestinationId: desID,
		})

		if err != nil {
			fmt.Fprint(rw, err.Error())
			return
		}
		b, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(rw, err.Error())
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(b))
	})
	server := http.Server{
		Addr: httpAddr,
	}
	return &server
}
