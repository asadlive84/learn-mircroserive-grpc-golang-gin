package main

import (
	"fmt"
	"log"
	"net"

	"github.com/asadlive84/auth-svc/pkg/config"
	"github.com/asadlive84/auth-svc/pkg/db"
	"github.com/asadlive84/auth-svc/pkg/pb"
	q "github.com/asadlive84/auth-svc/pkg/query"
	"github.com/asadlive84/auth-svc/pkg/services"
	"github.com/jmoiron/sqlx"

	// "github.com/asadlive84/auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h, err := db.DbInit(db.DBConfig{
		DbUserName: c.DatabaseUserName,
		DbName:     c.DatabaseName,
		DbPassword: c.DbPassword,
		DbPort:     c.DbPort,
	})

	if err != nil {
		log.Fatalln("Failed to database open up:", err)
	}

	dbCOnfig, err := db.NewDatabaseConfig(db.DBConfig{
		DbUserName: c.DatabaseUserName,
		DbName:     c.DatabaseName,
		DbPassword: c.DbPassword,
		DbPort:     c.DbPort,
		DbHost:     c.DbHost,
	})

	hdb, err := sqlx.Connect("postgres", dbCOnfig)

	if err != nil {
		log.Fatalln("Failed to database sqlx open up:", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	queryDb := q.QueryInit{
		Db: hdb,
	}

	s := services.Server{
		H: h,
		Q: &queryDb,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
