package routes

import (
	"context"
	"net/http"

	"github.com/asadlive84/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type GetProductRequestBody struct {
	ProductID string `json:"product_id"`
}

func GetProduct(ctx *gin.Context, d pb.ProductServiceClient) {

	body := GetProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := d.GetProduct(context.Background(), &pb.GetProductRequest{
		ProductID: body.ProductID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
