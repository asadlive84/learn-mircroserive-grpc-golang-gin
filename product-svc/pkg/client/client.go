package client

import (
	"context"
	"fmt"

	"github.com/asadlive84/product-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func NewAuthServiceClient(url string) AuthServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	p := AuthServiceClient{
		Client: pb.NewAuthServiceClient(cc),
	}

	return p
}

func (a *AuthServiceClient) CheckUser(userID string) (*pb.CheckUserResponse, error) {
	req := pb.CheckUserRequest{
		UserID: userID,
	}
	fmt.Println("######################################################")
	fmt.Printf("CheckUser CheckUser: %+v\n", userID)
	fmt.Println("######################################################")
	
	return a.Client.CheckUser(context.Background(), &req)
}
