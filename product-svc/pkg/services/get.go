package services

import (
	"context"
	"os"
	"time"

	// "log"
	"net/http"

	"github.com/asadlive84/product-svc/pkg/pb"
	q "github.com/asadlive84/product-svc/pkg/query"

	log "github.com/sirupsen/logrus"
)

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"ProductID": req.ProductID}).Info("Service: ProductID")

	product, err := s.Q.GetProductByID(req.ProductID)
	if err != nil && err != q.NotFound {
		logger.WithFields(log.Fields{"ProductID": req.ProductID}).Errorf("an error in get ProductID query %+v", err)
		return &pb.GetProductResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	if product == nil {
		logger.Errorln("ProductID is not found")
		return &pb.GetProductResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	logger.WithFields(log.Fields{"ProductID": req.ProductID}).Info("login successfully")
	return &pb.GetProductResponse{
		Status: http.StatusCreated,
	}, nil
}
