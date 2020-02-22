package hdm

import (
	"context"
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

func NewHTTPServer() http.Server {
	http.HandleFunc("/hotel", func(rw http.ResponseWriter, rq *http.Request) {
		resp, err := client.Hotels(context.Background(), &pb.HotelsRequest{
			Msg: "1",
		})

		if err != nil {
			fmt.Fprint(rw, err.Error())
			return
		}

		fmt.Fprint(rw, "abc"+resp.GetMsg())
	})
	server := http.Server{
		Addr: httpAddr,
	}
	return server
}
