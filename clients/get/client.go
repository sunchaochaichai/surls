package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"surls/pb"
	"flag"
	"encoding/json"
)

func main() {

	var str string
	flag.StringVar(&str, "s", "", "string")

	var address string
	flag.StringVar(&address, "addr", "localhost:7070", "addr")

	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSUrlsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.GetReq{Url: str})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	resp, _ := json.Marshal(&r)
	log.Println(string(resp))
}
