package main

import (
	"log"

	"github.com/asadlive84/api-gateway/pkg/auth"
	"github.com/asadlive84/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	// authSvc := *auth.RegisterRoutes(r, &c)
	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)

}
