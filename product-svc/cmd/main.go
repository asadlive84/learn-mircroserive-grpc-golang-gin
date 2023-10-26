package main

import (
	"fmt"
	"log"
	"net"

	cl "github.com/asadlive84/product-svc/pkg/client"
	"github.com/asadlive84/product-svc/pkg/config"
	"github.com/asadlive84/product-svc/pkg/db"
	"github.com/asadlive84/product-svc/pkg/pb"
	q "github.com/asadlive84/product-svc/pkg/query"
	"github.com/asadlive84/product-svc/pkg/services"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Default().Printf("error is %+v", err)
		log.Fatalln("Failed to getting config", err)
	}

	dbCOnfig, err := db.NewDatabaseConfig(db.DBConfig{
		POSTGRES_USER:     c.POSTGRES_USER,
		POSTGRES_PASSWORD: c.POSTGRES_PASSWORD,
		POSTGRES_DB:       c.POSTGRES_DB,
		POSTGRES_PORT:     c.POSTGRES_PORT,
		POSTGRES_HOST:     c.POSTGRES_HOST,
		PORT:              c.PORT,
		JWT_SECRET_KEY:    c.JWT_SECRET_KEY,
	})

	_, err = db.DbInit(dbCOnfig)

	if err != nil {
		log.Fatalln("Failed to database migration:", err)
	}

	hdb, err := sqlx.Connect("postgres", dbCOnfig)

	if err != nil {
		log.Fatalln("Failed to database sqlx open up:", err)
	}

	lis, err := net.Listen("tcp", c.PORT)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.PORT)

	queryDb := q.QueryInit{
		Db: hdb,
	}

	s := services.Server{
		Q:       &queryDb,
		AuthSvc: cl.NewAuthServiceClient(c.AUTH_PORT),
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	fmt.Println("=============>Product Server is running<===============")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
