package client

import (
	"context"
	"fmt"
	contractApi "kratosEntContractService/api/contract"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	req := contractApi.CreateContractRequest{
		StudentCode:          "2180607001",
		FirstName:            "Cao",
		LastName:             "Quat",
		MiddleName:           "Ba",
		Email:                "quatcao@gmail.com",
		Phone:                "0123456789",
		Gender:               1,
		Dob:                  timestamppb.Now(),
		Address:              "475 A Dien Bien Phu",
		Avatar:               "SGVsbG8sIFdvcmxkIQ==",
		RoomId:               "B0510",
		IsActive:             false,
		Sign:                 "baquat@gmail.com",
		NotificationChannels: 1,
	}

	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := contractApi.NewContractServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rep, err := client.CreateContract(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rep)
}
