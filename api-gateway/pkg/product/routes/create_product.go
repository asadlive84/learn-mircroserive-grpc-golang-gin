package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asadlive84/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type ProductRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	UserID      string `json:"user_id"`
}

func CreateProduct(ctx *gin.Context, p pb.ProductServiceClient) {

	body := ProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println("######################################################")
	fmt.Printf("body: %+v\n", body)
	fmt.Println("######################################################")

	res, err := p.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:        body.Name,
		Description: body.Description,
		IsActive:    body.IsActive,
		UserID:      body.UserID,
	})

	fmt.Println("###################@@###################################")
	fmt.Printf("err: %+v\n", err)
	fmt.Printf("res: %+v\n", res)
	fmt.Println("###################@@@###################################")

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
