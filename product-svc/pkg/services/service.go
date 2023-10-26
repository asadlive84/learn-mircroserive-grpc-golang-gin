package services

import (
	c "github.com/asadlive84/product-svc/pkg/client"
	q "github.com/asadlive84/product-svc/pkg/query"
)

type Server struct {
	// H db.DbHandler
	Q       q.Query
	AuthSvc c.AuthServiceClient
}
