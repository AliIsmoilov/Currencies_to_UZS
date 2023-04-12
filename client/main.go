package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "app/currency"
)

func main() {

	conn, err := grpc.Dial("localhost:9001", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	
	if err != nil {
		log.Fatalf("error to connect: %+v", err)
	}

	c := pb.NewCurrencyServiceClient(conn)

	resp, _ := c.GetCurrency(context.Background(), &pb.Currency{Currency: "U"})

	fmt.Println(resp)

}