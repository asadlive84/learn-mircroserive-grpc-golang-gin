package product


import (
	"github.com/asadlive84/api-gateway/pkg/product/routes"
	"github.com/asadlive84/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.POST("/create", svc.CreateProduct)
	routes.POST("/get", svc.GetProduct)

	return svc

}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) GetProduct(ctx *gin.Context) {
	routes.GetProduct(ctx, svc.Client)
}
