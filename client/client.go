package main

import (
	"fmt"
	"log"
	"context"
	calculatorPB "github.com/pjchender/besg-grpc/proto/calculator"
	"google.golang.org/grpc"
)

func doUnary(client calculatorPB.CalculatorServiceClient) {
	fmt.Println("Staring to do a Unary RPC")
	req := &calculatorPB.CalculatorRequest{
	 A: 6,
	 B: 15,
	}
   
	res, err := client.Sum(context.Background(), req)
	if err != nil {
	 log.Fatalf("error while calling CalculatorService: %v \n", err)
	}
   
	log.Printf("Response from CalculatorService: %v", res.Result)
   }

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
	 log.Fatalf("failed to dial: %v", err)
	}
   
	defer conn.Close()
   
	client := calculatorPB.NewCalculatorServiceClient(conn)
   
	doUnary(client)
}