package product

import (
	"fmt"

	"github.com/asadlive84/api-gateway/pkg/config"
	"github.com/asadlive84/api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {

	fmt.Println("######################################################")
	fmt.Printf("c.ProductSvcUrl: %+v", c.ProductSvcUrl)
	fmt.Println("######################################################")

	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewProductServiceClient(cc)
}
