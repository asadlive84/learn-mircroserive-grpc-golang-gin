package services

import (
	"context"
	"fmt"
	"os"
	"time"

	// "log"
	"net/http"

	"github.com/asadlive84/product-svc/pkg/pb"
	q "github.com/asadlive84/product-svc/pkg/query"

	log "github.com/sirupsen/logrus"
)

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req.Name}).Info("Service: CreateProduct method")
	logger.WithFields(log.Fields{"Description": req.Description}).Info("Service: CreateProduct method")

	fmt.Println("######################################################")
	fmt.Printf("UserID: %+v\n", req.UserID)
	fmt.Printf("Description: %+v\n", req.Description)
	fmt.Printf("Name: %+v\n", req.Name)
	fmt.Println("######################################################")
	

	userInfo, err := s.AuthSvc.CheckUser(req.UserID)
	if userInfo.Status == 400 || err != nil {
		logger.WithFields(log.Fields{"UserID": req.Name}).Errorf("auth service an error %+v", err)
		return &pb.CreateProductResponse{
			Status: http.StatusBadRequest,
		}, nil

	}

	fmt.Println("######################################################")
	fmt.Printf("userInfo: %+v\n", userInfo)
	fmt.Printf("err: %+v\n", err)
	fmt.Println("######################################################")

	err = s.Q.InsertProduct(q.Product{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
		CreatedBy:   req.UserID,
	})

	if err != nil {
		logger.WithFields(log.Fields{"Name": req.Name}).Errorf("an error in query %+v", err)
		return &pb.CreateProductResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	logger.WithFields(log.Fields{"Name": req.Name}).Info("Successfully created product")

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
	}, nil
}
